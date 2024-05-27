import { Await, Outlet, useLoaderData } from "@remix-run/react";
import { LoaderFunctionArgs, json } from "@remix-run/router";
import Footer from "~/lib/components/template/Footer";
import Navbar from "~/lib/components/template/Navbar";
import Sidebar from "~/lib/components/template/Sidebar";
import { requireCompleteUserSession } from "~/lib/util/session";
import { decks } from "../lib/infra/db/db.server";

export async function loader({ request }: LoaderFunctionArgs) {
  const { user } = await requireCompleteUserSession(request);

  const userInfo = {
    username: user.username,
    email: user.email,
    displayName: user.displayName,
    isVerified: user.isVerified,
  };

  return json({ userInfo });
}

async function getDecks() {
  const featured = await decks.v1GetDecks({
    limit: 5,
    // featured: true,
  });

  const recent = await decks.v1GetDecks({
    limit: 5,
    // order: "update",
  });

  featured.map((deck) => {
    deck.id;
  });

  return { featured, recent };
}

export default function Explore() {
  const { userInfo } = useLoaderData<typeof loader>();

  return (
    <div className="min-h-screen bg-neutral-100 flex flex-col">
      <div className="flex h-full w-full">
        <Sidebar />
        <div className="page w-full relative">
          <Navbar user={userInfo} />
          {/* <div className="navbar w-full h-16 bg-neutral-400"></div> */}
          <main className="mt-4 max-w-screen-2xl w-full md:p-2.5 mb-28">
            <div className="flex flex-col gap-6 px-5">
              <div className="w-full">
                <h1 className="font-bold text-3xl text-neutral-900 py-2 ">
                  Explore
                </h1>
                <p className="max-w-3xl">Explore and find new decks</p>
              </div>
              <div className="space-y-2">
                <Await resolve={getDecks()}>
                  {({ featured, recent }) => (
                    <>
                      <div className="grid grid-cols-5">
                        {featured.map((deck) => (
                          <div
                            key={deck.id}
                            className="flex flex-col bg-white rounded-md p-4 mb-4"
                          >
                            <h2 className="text-2xl font-bold">{deck.name}</h2>
                            <p className="text-neutral-500">
                              {deck.description}
                            </p>
                            <p className="text-neutral-500">
                              {deck.subjectIds.length} cards
                            </p>
                          </div>
                        ))}
                      </div>
                      <div className="grid grid-cols-5">
                        {recent.map((deck) => (
                          <div
                            key={deck.id}
                            className="flex flex-col bg-white rounded-md p-4 mb-4"
                          >
                            <h2 className="text-2xl font-bold">{deck.name}</h2>
                            <p className="text-neutral-500">
                              {deck.description}
                            </p>
                            <p className="text-neutral-500">
                              {deck.subjectIds.length} cards
                            </p>
                          </div>
                        ))}
                      </div>
                    </>
                  )}
                </Await>
              </div>
            </div>
          </main>
        </div>
      </div>
      <Footer />
    </div>
  );
}
