import * as argon2 from "argon2";
import * as base58 from "bs58";

import { randomBytes, randomUUID } from "crypto";
import {
	InvalidRequestError,
	ResourceNotFoundError,
	UnknownError,
} from "../domain/error";
import {
	GenerateTokenDto,
	GenerateTokenSchema,
	Token,
	UpdateTokenDto,
	UpdateTokenSchema,
} from "../domain/token";

import { ITokenRespository } from "../ports/tokens";
import { Validator } from "../validator";
import { validateId } from "./common";

const PREFIX_LEN = 8;
const TOKEN_SEPARATOR = "-";

export const GenerateTokenValidator = new Validator(GenerateTokenSchema);
export const UpdateTokenValidator = new Validator(UpdateTokenSchema);

export class TokensAdapter<R extends ITokenRespository> {
	constructor(private tokensRepository: R) {}

	async getTokens(userId: string): Promise<Token[]> {
		return this.tokensRepository.getTokens(userId);
	}

	async getToken(userId: string, tokenId: string): Promise<Token> {
		const foundToken = await this.tokensRepository.getToken(userId, tokenId);
		if (!foundToken) {
			throw new ResourceNotFoundError({
				cause: new ResourceNotFoundError({
					cause: new Error("Token not found"),
					context: { tokenId },
					description: "The provided token does not exists.",
				}),
			});
		}
		return foundToken;
	}

	async generateToken(
		userId: string,
		data: GenerateTokenDto,
	): Promise<Token & { generatedToken: string }> {
		GenerateTokenValidator.validate(data);

		console.debug("creating token:", { ownerId: userId, data });
		const generated = await generateRandomToken();
		console.debug("generated random token");

		const createdToken = await this.tokensRepository.createToken(userId, {
			name: data.name,
			claims: data.claims,
			prefix: generated.prefix,
			token: generated.tokenHash,
		});

		return {
			...createdToken,
			generatedToken: `${generated.prefix}${TOKEN_SEPARATOR}${generated.token}`,
		};
	}

	async updateToken(
		userId: string,
		tokenId: string,
		changes: UpdateTokenDto,
	): Promise<Token> {
		validateId(tokenId);

		const updatedToken = await this.tokensRepository.updateToken(
			userId,
			tokenId,
			changes,
		);

		return updatedToken;
	}

	async deleteToken(userId: string, tokenId: string): Promise<void> {
		validateId(tokenId);
		await this.tokensRepository.deleteToken(userId, tokenId);
	}
}

const TokenHashOptions = {
	timeCost: 2,
	memoryCost: 4 * 1024,
	parallelism: 4,
	hashLength: 32,
	type: argon2.argon2id,
} satisfies Parameters<typeof argon2.hash>[1];

async function generateRandomToken(): Promise<{
	token: string;
	prefix: string;
	tokenHash: string;
}> {
	console.debug("calling randomBytes");
	const prefixBytes = randomBytes(PREFIX_LEN / 2);
	console.debug("calling Buffer.from randomUUID");
	const tokenBytes = Buffer.from(randomUUID({ disableEntropyCache: true }));

	console.debug("calling encodeToken");
	const [prefix, token] = encodeToken([prefixBytes, tokenBytes]);
	console.debug("calling hashToken");
	const tokenHash = await hashToken(tokenBytes, prefixBytes);

	console.debug("generated token");

	return { token, prefix, tokenHash };
}

export function hashToken(prefix: Buffer, token: Buffer): Promise<string> {
	console.debug("concatenating prefix and token");
	const prefixExt = Buffer.concat([prefix, Buffer.from("0".repeat(16), "hex")]);
	console.debug("salt length:", prefixExt.length);
	try {
		console.debug("calling argon2.hash");
		return argon2.hash(token, {
			...TokenHashOptions,
			salt: prefixExt,
		});
	} catch (error) {
		throw new UnknownError({
			cause: error,
			description: "An unknown error occurred while hashing the token.",
		});
	}
}

/**
 * Unmarshal a token into its prefix and token parts.
 *
 * @param encodedToken The encoded token to unmarshal.
 * @returns The prefix and token parts of the token.
 */
export function unmarshalToken(encodedToken: string): [string, string] {
	const tokenParts = encodedToken.split(TOKEN_SEPARATOR, 2);
	if (tokenParts.length !== 2) {
		throw new InvalidRequestError({
			cause: new Error("Invalid token format"),
			description: "The provided token is not valid.",
		});
	}
	return tokenParts as [string, string];
}

/**
 * Decode the token parts into binary buffers.
 *
 * @param tokenParts The token to decode.
 * @returns The decoded prefix and token parts.
 */
export function decodeToken([prefix, token]: [string, string]): [
	Buffer,
	Buffer,
] {
	try {
		const tokenBytes = Buffer.from(base58.decode(token));
		const prefixBytes = Buffer.from(prefix, "hex");

		return [prefixBytes, tokenBytes];
	} catch (error) {
		throw new InvalidRequestError({
			cause: error,
			description: "The provided token is not valid.",
		});
	}
}

export function encodeToken([prefixBytes, tokenBytes]: [Buffer, Buffer]): [
	string,
	string,
] {
	const prefix = prefixBytes.toString("hex");
	const token = base58.encode(tokenBytes);
	return [prefix, token];
}
