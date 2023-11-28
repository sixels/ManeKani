import {
	BaseError,
	ForbiddenError,
	ISsoAuthenticator,
	ITokenRespository,
	SsoAuthAdapter,
	TokenAuthAdapter,
	UnauthorizedError,
	UserSession,
} from "@manekani/core";
import { CanActivate, ExecutionContext, Injectable } from "@nestjs/common";

import { Reflector } from "@nestjs/core";
import { Request } from "express";
import { SsoAuthService } from "./sso.service";
import { TokenAuthService } from "./token.service";
import { getRequest } from "./util";

// const AUTH_COOKIE_KEY = 'ory_kratos_session';
export const META_AUTH_SCOPES_KEY = "MK_TOKEN_SCOPES";
export const META_AUTH_METHOD_KEY = "MK_TOKEN_METHOD";

export type AuthMethod = "any" | "login";

@Injectable()
export class AuthGuard implements CanActivate {
	constructor(
		private reflector: Reflector,
		// biome-ignore lint/suspicious/noExplicitAny: any is just a placeholder type
		private ssoProvider: SsoAuthService<any>,
		// biome-ignore lint/suspicious/noExplicitAny: any is just a placeholder type
		private tokenProvider?: TokenAuthService<any>,
	) {}

	async canActivate(context: ExecutionContext): Promise<boolean> {
		const request = getRequest(context);
		const authMethod =
			this.reflector.getAllAndOverride<AuthMethod | undefined>(
				META_AUTH_METHOD_KEY,
				[context.getHandler(), context.getClass()],
			) ?? "any";
		const requiredScopes = this.reflector.getAllAndOverride<string[]>(
			META_AUTH_SCOPES_KEY,
			[context.getHandler(), context.getClass()],
		);

		if (authMethod === "any" && this.tokenProvider) {
			try {
				request.user = await authorizeApiToken(
					this.tokenProvider,
					request,
					requiredScopes,
				);
				return true;
			} catch (error) {
				if (error instanceof ForbiddenError) {
					throw error;
				}
				if (error instanceof BaseError) {
					console.debug("Api token authentication failed:", {
						message: error.message,
						context: error.context,
					});
				}
			}
		}

		if (this.ssoProvider) {
			try {
				request.user = await authorizeLogin(this.ssoProvider, request);
				return true;
			} catch (error) {
				if (error instanceof BaseError) {
					console.debug("Login authentication failed:", {
						cause: error.cause,
						context: error.context,
					});
				}
			}
		}

		throw new UnauthorizedError({
			cause: new Error("No auth providers could authorize the request"),
			description:
				"User authentication failed. Make sure to login or provide an API key first.",
			context: { authMethod, requiredScopes },
		});
	}
}

async function authorizeApiToken(
	tokenProvider: TokenAuthAdapter<ITokenRespository>,
	request: Request,
	requiredScopes: string[] = [],
): Promise<UserSession> {
	const requestToken = extractApiTokenFromRequest(request);
	if (!requestToken) {
		throw new UnauthorizedError({
			cause: new Error("The request is missing an API token"),
			description:
				"User authentication failed. Make sure to provide an API key.",
		});
	}

	return await tokenProvider.getSessionFromToken(requestToken, requiredScopes);
}

async function authorizeLogin(
	ssoProvider: SsoAuthAdapter<ISsoAuthenticator>,
	request: Request,
): Promise<UserSession> {
	const loginCookies = extractLoginCookiesFromRequest(ssoProvider, request);
	if (!loginCookies) {
		throw new UnauthorizedError({
			cause: new Error("The request is missing required login cookies"),
			description: "User authentication failed. Make sure to login first.",
		});
	}
	return await ssoProvider.getCookieSession(loginCookies);
}

function extractApiTokenFromRequest(request: Request): string | null {
	const tokenPrefix = "Bearer ";
	const authHeader = request.headers.authorization ?? "";

	console.log(`"${authHeader.slice(tokenPrefix.length)}"`);
	return authHeader.startsWith(tokenPrefix)
		? authHeader.slice(tokenPrefix.length)
		: null;
}

function extractLoginCookiesFromRequest(
	ssoProvider: SsoAuthAdapter<ISsoAuthenticator>,
	request: Request,
): string | null {
	const cookies = request.headers.cookie || "";
	console.log("cookies", cookies);
	return ssoProvider.hasRequiredCookies(request.cookies) ? cookies : null;
}
