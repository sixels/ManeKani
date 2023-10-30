import { ory } from './auth.server';
import { Session as OrySession } from '@ory/client';

export interface Session {
  session: OrySession;
  readonly logout_url: string;

  username: string;
}

// async function preloadSession(cookies: string) {
//   void (await getSession(cookies));
// }

async function getSession(cookies: string): Promise<Session> {
  try {
    const params = { cookie: cookies };
    const session = await ory.toSession(params).then(({ data }) => data);

    const logout_url = await ory
      .createBrowserLogoutFlow(params)
      .then(({ data }) => data.logout_url);

    return {
      session,
      logout_url,
      username:
        session.identity.traits.username || session.identity.traits.email,
    };
  } catch (e) {
    console.error('could not fetch the user:', e);
    return Promise.reject(e);
  }
}

export { getSession };
