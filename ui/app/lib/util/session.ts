import { User } from "@manekani/core";
import { redirect } from "@remix-run/node";
import { getLoginURL } from "../infra/auth/auth.server";
import { Session, getSession } from "../infra/auth/session.server";
import { users } from "../infra/db/db.server";

export async function requireSession(request: Request) {
	const cookies = request.headers.get("cookie");
	const loginUrl = getLoginURL(request.url);

	if (!cookies) {
		throw redirect(loginUrl);
	}

	const session = await getSession(cookies);
	if (!session) {
		throw redirect(loginUrl);
	}

	return session;
}

export async function requireUserSession(request: Request) {
	const loginUrl = getLoginURL(request.url);
	const session = await requireSession(request);

	const user = await users.getUser({
		userId: session.userSession.userId,
		email: session.userSession.email,
	});
	if (!user) {
		console.error("could not find the user in the database");
		throw redirect(loginUrl);
	}

	return { session, user };
}

export async function requireCompleteUserSession(request: Request): Promise<{
	session: Session;
	user: { username: NonNullable<User["username"]> } & Omit<User, "username">;
}> {
	const { session, user } = await requireUserSession(request);
	if (!user.isComplete || !user.username) {
		console.error("user is not complete")
		throw redirect("/complete-profile");
	}
	return { session, user: Object.assign(user, { username: user.username }) };
}
