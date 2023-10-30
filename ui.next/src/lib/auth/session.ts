import { cache } from "react";
import { ory } from ".";
import { Session as OrySession } from "@ory/client";
import "server-only";
import { cookies } from "next/headers";

export interface Session {
  session: OrySession;
  readonly logout_url: string;

  username: string;
}

export const preloadSession = async (cookies: string) => {
  void (await getSession(cookies));
};

export const getSession = cache(async (cookies: string): Promise<Session> => {
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
    console.error("could not fetch the user:", e);
    return Promise.reject(e);
  }
});

export const getSessionSSR = () => {
  return getSession(cookies().toString());
};
