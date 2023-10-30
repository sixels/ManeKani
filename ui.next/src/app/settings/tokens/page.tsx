import { CreateToken } from '@/components/Settings/CreateToken';
import { UserTokens } from '@/components/Settings/UserTokens';
import { ory } from '@/lib/auth';
import { headers } from 'next/headers';
import { PrismaClient } from '@prisma/client';
import { TokensAdapter } from 'manekani-core';
import { TokensDatabase, DbClient } from 'manekani-infra-db';

process.env.DATABASE_URL =
  'postgresql://manekani:secret@postgres-manekani:5432/manekani?schema=public';

export default async function TokensPage() {
  const header = headers();

  const db = new PrismaClient();
  // await db.connect();

  // const tokens = new TokensAdapter(new TokensDatabase(db));

  // const userSession = await ory.toSession({
  //   cookie: header.get('cookie')!,
  // });
  // const userId = userSession.data.identity.id;

  // const userTokens = await tokens.getTokens(userId);

  return (
    <div className="w-full">
      <section>
        <div className="inline-flex justify-between items-center w-full">
          <h1 className="text-2xl sm:text-3xl font-bold">
            Personal API Tokens
          </h1>
          <CreateToken />
        </div>
        <p className="mt-1 py-3">
          An API token can be used by third-party applications to improve your
          experience. Be careful with what each token can do as it may be used
          to access your data.
        </p>
      </section>
      <section className="mt-2">
        <UserTokens tokens={[]} />
      </section>
    </div>
  );
}
