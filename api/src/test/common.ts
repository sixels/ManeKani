import {
	ForbiddenError,
	ISsoAuthenticator,
	ITokenRespository,
	TokenClaims,
	UnauthorizedError,
	decodeToken,
	hashToken,
	unmarshalToken,
} from "@manekani/core";

import { UUID } from "crypto";
import { createMock } from "@golevelup/ts-jest";

export const EMPTY_UUID: UUID = "00000000-0000-0000-0000-000000000000",
	TEST_UUID: UUID = "00000000-0000-0000-0000-000000000001";

export const TEST_API_TOKEN =
		"FFFFFFFFFF-kh3vvyXiR9gTBRfQXkGohfxmr3fwYzAEuCXRsxeeyuDWhcoDS",
	UNPRIVILEGED_API_TOKEN =
		"AAAAAAAAAA-kgqFduTG8CpmTskLRHiz5rp1QjRg59eFcZ1JFDuRscumKHrb1";

const TEST_API_TOKEN_HASH = async () =>
	await hashToken(...decodeToken(unmarshalToken(TEST_API_TOKEN)));
const UNPRIVILEGED_API_TOKEN_HASH = async () =>
	await hashToken(...decodeToken(unmarshalToken(UNPRIVILEGED_API_TOKEN)));

export async function mockTokenService(): Promise<ITokenRespository> {
	const tokenService = createMock<ITokenRespository>();

	const testTokenHash = await TEST_API_TOKEN_HASH();
	const unprivilegedTokenHash = await UNPRIVILEGED_API_TOKEN_HASH();

	const testToken = {
		id: TEST_UUID,
		ownerId: EMPTY_UUID,
		createdAt: new Date(),
		isActive: true,
		name: "test api token",
		prefix: "FFFFFFFFFF",
		token: testTokenHash,
		usedAt: undefined,
		claims: {
			...createMock<TokenClaims>(),
			subjectCreate: true,
			subjectDelete: true,
			subjectUpdate: true,
		},
	};

	tokenService.useToken.mockImplementation(async (tokenHash) => {
		console.debug(tokenHash, testTokenHash);
		if (tokenHash === testTokenHash) {
			return testToken;
		}
		if (tokenHash === unprivilegedTokenHash) {
			throw new ForbiddenError({});
		}
		throw new UnauthorizedError({});
	});

	return tokenService;
}

export function mockSsoService(): ISsoAuthenticator {
	const ssoService = createMock<ISsoAuthenticator>();
	ssoService.requiredCookies.mockReturnValue(["foo"]);
	ssoService.getCookieSession.mockResolvedValue({
		userId: EMPTY_UUID,
		email: "foo@bar.baz",
	});

	return ssoService;
}
