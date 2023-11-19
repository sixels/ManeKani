import { redirect } from '@remix-run/node';
import { getSession } from '../infra/auth/session.server';
import { getLoginURL } from '../infra/auth/auth.server';
import { User, UserSession } from 'manekani-core';
import { users } from '../infra/db/db.server';

type UserAndSession = {
  session: UserSession;
  user: User;
};

export async function requireSession(request: Request) {
  const cookies = request.headers.get('cookie');
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
    console.error('could not find the user in the database');
    throw redirect(loginUrl);
  }

  return { session, user };
}

export async function requireCompletedUserSession(request: Request) {
  const { session, user } = await requireUserSession(request);
  if (!user.isComplete) {
    throw redirect('/complete-profile');
  }
  return { session, user };
}
