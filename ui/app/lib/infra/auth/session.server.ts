import { UserSession } from "@manekani/core";
import { Session as OrySession } from "@ory/client";
import { ory } from "./auth.server";

export interface Session {
	ctx: OrySession;
	readonly logout_url: string;

	userSession: UserSession;
}

export async function getSession(cookies: string): Promise<Session | null> {
	try {
		const params = { cookie: cookies };
		const session = await ory.toSession(params).then(({ data }) => data);

		if (!session || !session.active) {
			return null;
		}

		const logout_url = await ory
			.createBrowserLogoutFlow(params)
			.then(({ data }) => data.logout_url);

		if (!session.identity) {
			console.error("identity is not present in session");
			return null;
		}

		return {
			ctx: session,
			logout_url,
			userSession: {
				userId: session.identity.id,
				email: session.identity.traits.email,
			},
		};
	} catch (e) {
		console.error("could not fetch the user");
		return null;
	}
}
