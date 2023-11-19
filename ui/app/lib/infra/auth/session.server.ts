import { UserSession } from 'manekani-core';
import { ory } from './auth.server';
import { Session as OrySession } from '@ory/client';

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

    return {
      ctx: session,
      logout_url,
      userSession: {
        userId: session.identity!.id,
        email: session.identity!.traits['email'],
      },
    };
  } catch (e) {
    console.error('could not fetch the user');
    return null;
  }
}
