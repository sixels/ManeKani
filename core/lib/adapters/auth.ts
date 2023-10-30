import { ISsoAuthenticator } from '../ports/auth';
import { decodeToken, hashToken, unmarshalToken } from './tokens';

import { ForbiddenError } from '../domain/error';
import { Token } from '../domain/token';
import { UserSession } from '../domain/auth';
import { ITokenRespository } from '../ports';

export class SsoAuthAdapter<R extends ISsoAuthenticator> {
  constructor(private ssoProvider: R) {}

  hasRequiredCookies(cookies: Record<string, string>): boolean {
    for (const required of this.ssoProvider.requiredCookies()) {
      if (!cookies[required]) {
        return false;
      }
    }
    return true;
  }

  async getCookieSession(cookies: string): Promise<UserSession> {
    const session = await this.ssoProvider.getCookieSession(cookies);
    return session;
  }
}

export class TokenAuthAdapter<R extends ITokenRespository> {
  constructor(private tokenProvider: R) {}

  async getSessionFromToken(
    token: string,
    requiredScopes: string[],
  ): Promise<UserSession> {
    const tokenHash = await hashToken(...decodeToken(unmarshalToken(token)));

    const userToken = await this.tokenProvider.useToken(tokenHash);

    if (!hasRequiredScopes(userToken, requiredScopes)) {
      throw new ForbiddenError({
        cause: new Error("The token doesn't have the required scopes"),
        description: `The following permission are missing from the token: ${requiredScopes.join(
          ', ',
        )}`,
      });
    }

    return { userId: userToken.ownerId };
  }
}

function hasRequiredScopes(token: Token, requiredScopes: string[]): boolean {
  // NOTE: keep in-sync with TokenClaimsSchema
  const claimsMap = {
    'deck:create': token.claims.deckCreate,
    'deck:delete': token.claims.deckDelete,
    'deck:update': token.claims.deckUpdate,
    'subject:create': token.claims.subjectCreate,
    'subject:delete': token.claims.subjectDelete,
    'subject:update': token.claims.subjectUpdate,
    'review:create': token.claims.reviewCreate,
    'study:material:create': token.claims.studyMaterialCreate,
    'study:material:delete': token.claims.studyMaterialDelete,
    'study:material:update': token.claims.studyMaterialUpdate,
    'user:update': token.claims.userUpdate,
  };
  return requiredScopes.every((scope) =>
    Boolean(claimsMap[scope as keyof typeof claimsMap]),
  );
}
