import { ORY_BROWSER_URL } from '@/lib/auth';
import { preloadSession } from '@/lib/auth/session';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';
import { Suspense } from 'react';

type LayoutProps = React.PropsWithChildren;

export default async function UserLayout({ children }: LayoutProps) {
  const cookie = cookies().toString();
  await preloadSession(cookie).catch((e) => {
    redirect(
      `${ORY_BROWSER_URL}/self-service/login/browser?return_to=http://127.0.0.1:11011/settings/tokens`,
    );
  });

  return (
    <>
      {/* <Suspense fallback={<></>}> */}
      {/* <Navbar /> */}
      {/* </Suspense> */}
      <div className="max-w-6xl pb-6 px-2 mx-auto mt-24  gap-y-2.5 gap-x-3 lg:grid-rows-[auto_1fr]">
        {/* <nav className="lg:col-span-2">
          <ul className="inline-flex font-medium gap-2 items-center [&:not(:last-child)]:bg-white">
            <li>
              <a
                href="/profile"
                className="hover:text-wk-accent-500 transition-colors"
              >
                profile
              </a>
            </li>
            <li>token</li>
          </ul>
        </nav> */}
        {/* <aside className="bg-white rounded-lg p-2.5">a</aside> */}
        <div className="w-full">{children}</div>
      </div>
    </>
  );
}
