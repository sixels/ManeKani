import type { LoaderFunctionArgs } from '@remix-run/node';
import type { Token } from 'manekani-core';
import { json, redirect } from '@remix-run/node';
import { TokensAdapter } from 'manekani-core';
import { getSession } from '~/lib/infra/auth/session.server';
import { tokensDb } from '~/lib/infra/db/tokens.server';
import { useLoaderData } from '@remix-run/react';

export async function loader({ request }: LoaderFunctionArgs) {
  const cookies = request.headers.get('cookie');

  if (!cookies) {
    // unauthenticated user
    return json({ tokens: [] });
    throw new Response('Unauthorized', { status: 401 });
  }
  const userSession = await getSession(cookies);
  const userId = userSession.session.identity.id;

  const a = tokensDb;
  // const tokensAdapter = new TokensAdapter(tokensDb);
  // return json(tokensAdapter.getTokens(userId));
  return json({ tokens: [] });
}

export default function Component() {
  const { tokens } = useLoaderData<typeof loader>();
  return <div>{tokens.length}</div>;
}
