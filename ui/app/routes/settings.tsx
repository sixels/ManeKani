import { LoaderFunctionArgs, json } from '@remix-run/node';
import { Outlet, useLoaderData } from '@remix-run/react';
import Footer from '~/lib/components/template/Footer';
import Navbar from '~/lib/components/template/Navbar';
import { requireCompletedUserSession } from '~/lib/util/session';

export async function loader({ request }: LoaderFunctionArgs) {
  const { user } = await requireCompletedUserSession(request);

  const userInfo = {
    username: user.username!,
    email: user.email,
    displayName: user.displayName,
    isVerified: user.isVerified,
  };

  return json({ userInfo });
}

export default function Settings() {
  const { userInfo } = useLoaderData<typeof loader>();

  return (
    <div className="min-h-screen bg-neutral-100 flex flex-col">
      <div className="flex h-full w-full">
        <div className="hidden md:block min-w-[300px] w-1/4 bg-white"></div>
        <div className="page w-full relative">
          <Navbar user={userInfo} />
          {/* <div className="navbar w-full h-16 bg-neutral-400"></div> */}
          <main className="mt-4 max-w-screen-2xl w-full md:p-2.5">
            <Outlet />
          </main>
        </div>
      </div>
      <Footer />
    </div>
  );
}
