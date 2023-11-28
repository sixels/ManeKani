import { BaseError, UserSession } from "@manekani/core";
import { UnauthorizedError } from "@manekani/core";
import { ISsoAuthenticator } from "@manekani/core";
import {
	Configuration,
	ConfigurationParameters,
	FrontendApi,
} from "@ory/client";

export type { ConfigurationParameters } from "@ory/client";

export class SsoAuthenticator implements ISsoAuthenticator {
	private ory: FrontendApi;

	constructor(options?: ConfigurationParameters) {
		this.ory = new FrontendApi(new Configuration(options));
	}

	requiredCookies(): string[] {
		return ["ory_kratos_session"];
	}

	async getCookieSession(cookies: string): Promise<UserSession> {
		try {
			const { data: orySession } = await this.ory.toSession({
				cookie: cookies,
			});

			if (!orySession.identity) {
				throw new UnauthorizedError({
					description: "User identity is not available",
				});
			}

			return {
				userId: orySession.identity.id,
				email: orySession.identity.traits.email,
			};
		} catch (error) {
			if (error instanceof BaseError) {
				throw error;
			}
			throw new UnauthorizedError({
				cause: error,
				description: "Could not fetch the user session",
			});
		}
	}

	async registerUsername(_userId: string, _username: string): Promise<void> {
		throw new Error("Method not implemented.");
	}

	async updateUsername(_userId: string, _username: string): Promise<void> {
		throw new Error("Method not implemented.");
	}
}
