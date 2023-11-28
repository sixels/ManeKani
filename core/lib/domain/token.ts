import { Static, Type } from "@sinclair/typebox";

import { UuidSchema } from "./common";

export type TokenClaims = Static<typeof TokenClaimsSchema>;
export const TokenClaimsSchema = Type.Object({
	deckCreate: Type.Boolean({ default: false }),
	deckDelete: Type.Boolean({ default: false }),
	deckUpdate: Type.Boolean({ default: false }),
	subjectCreate: Type.Boolean({ default: false }),
	subjectDelete: Type.Boolean({ default: false }),
	subjectUpdate: Type.Boolean({ default: false }),
	reviewCreate: Type.Boolean({ default: false }),
	studyMaterialCreate: Type.Boolean({ default: false }),
	studyMaterialDelete: Type.Boolean({ default: false }),
	studyMaterialUpdate: Type.Boolean({ default: false }),
	userUpdate: Type.Boolean({ default: false }),
});

export type TokenWithHash = Static<typeof TokenWithHashSchema>;
export const TokenWithHashSchema = Type.Object({
	id: UuidSchema,
	createdAt: Type.Date(),
	usedAt: Type.Optional(Type.Date()),

	/**
	 * Is this token active.
	 */
	isActive: Type.Boolean(),
	/**
	 * The token's name.
	 */
	name: Type.String({ maxLength: 25, minLength: 1 }),
	/**
	 * An already hashed token.
	 */
	token: Type.String(),
	/**
	 * The token's prefix.
	 */
	prefix: Type.String(),

	/**
	 * The token's claims/permissions.
	 */
	claims: TokenClaimsSchema,

	ownerId: Type.String(),
});
export type Token = Static<typeof TokenSchema>;
export const TokenSchema = Type.Omit(TokenWithHashSchema, ["token"]);

/** Stores the final token */
export type CreateTokenDto = Static<typeof CreateTokenSchema>;
export const CreateTokenSchema = Type.Pick(TokenWithHashSchema, [
	"name",
	"token",
	"prefix",
	"claims",
]);

/** Generates a new token */
export type GenerateTokenDto = Static<typeof GenerateTokenSchema>;
export const GenerateTokenSchema = Type.Pick(TokenSchema, ["name", "claims"]);

export type UpdateTokenDto = Static<typeof UpdateTokenSchema>;
export const UpdateTokenSchema = Type.Partial(
	Type.Pick(TokenSchema, ["name", "isActive"]),
);
