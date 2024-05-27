import { useMatch } from "@remix-run/react";
import classNames from "classnames";

export default function Sidebar() {
  return (
    <aside className="hidden md:block min-w-[300px] w-1/4 bg-white py-4 px-2">
      <div className="flex items-center justify-center">
        <a href="/">
          <img
            src="/assets/logop.png"
            alt="ManeKani Logo"
            className="h-[35px] hidden md:block"
          />
          <img
            src="/assets/logop.png"
            alt="ManeKani Logo"
            className="h-[35px] block md:hidden"
          />
        </a>
      </div>

      <div className="mt-8">
        <ul className="space-y-1">
          {SidebarItems.map(({ label, url, subItems }) => (
            <li key={label} className="w-full font-medium text-neutral-500">
              {subItems?.length ? (
                <div className="flex flex-col space-y-2">
                  <a
                    href={url}
                    className={classNames(
                      useMatch(url) ? "text-neutral-900" : "",
                      "block w-full py-2 px-2.5 hover:text-neutral-800 rounded-sm hover:bg-neutral-100 transition-colors",
                    )}
                  >
                    {label}
                  </a>
                  <div className="relative group">
                    <div className="absolute h-full w-[2px] left-4 top-0 bg-neutral-500 group-hover:bg-neutral-800" />
                    <ul className="space-y-1 pl-6">
                      {subItems.map(({ label, url }) => (
                        <li
                          key={label}
                          className="w-full font-medium text-neutral-500 hover:text-neutral-800 rounded-sm hover:bg-neutral-100 transition-colors"
                        >
                          <a
                            href={url}
                            className={classNames(
                              useMatch(url) ? "text-neutral-900" : "",
                              "block w-full py-2 px-2.5",
                            )}
                          >
                            {label}
                          </a>
                        </li>
                      ))}
                    </ul>
                  </div>
                </div>
              ) : (
                <a
                  href={url}
                  className={classNames(
                    useMatch(url) ? "text-neutral-900" : "",
                    "block w-full py-2 px-2.5 hover:text-neutral-800 rounded-sm hover:bg-neutral-100 transition-colors",
                  )}
                >
                  {label}
                </a>
              )}
            </li>
          ))}
        </ul>
      </div>
    </aside>
  );
}

type SidebarItem = {
  label: string;
  url: string;
  subItems?: Omit<SidebarItem, "subItems">[];
};

const SidebarItems: SidebarItem[] = [
  { label: "Dashboard", url: "/" },
  { label: "Decks", url: "/decks" },
  { label: "Explore", url: "/explore" },
  {
    label: "Settings",
    url: "/settings",
    subItems: [
      {
        label: "App",
        url: "/settings/app",
      },
      {
        label: "Account",
        url: "/settings/account",
      },
      {
        label: "API Tokens",
        url: "/settings/api-tokens",
      },
    ],
  },
];
