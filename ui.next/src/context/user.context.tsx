// "use client";

// import { ory } from "@/lib/auth";
// import { Session } from "@ory/client";
// import { edgeConfig } from "@ory/integrations/next";
// import { useRouter } from "next/router";
// import React, { useEffect, useState } from "react";

// interface State {
//   username: string;
//   session: Session;
//   logout_url?: string; // TODO: use a logout function instead
// }

// export const UserContext = React.createContext<State | null>(null);

// const UserProvider: React.FC<React.PropsWithChildren> = ({ children }) => {
//   const router = useRouter();

//   const [session, setSession] = useState<Session | undefined>();
//   const [logoutUrl, setLogoutUrl] = useState<string | undefined>();

//   useEffect(() => {
//     ory
//       .toSession()
//       .then(({ data }) => {
//         // User has a session!
//         setSession(data);
//         // Create a logout url
//         ory.createBrowserLogoutFlow().then(({ data }) => {
//           setLogoutUrl(data.logout_url);
//         });
//       })
//       .catch((e) => {
//         // Redirect to login page
//         console.error(e);
//         return router.push(edgeConfig.basePath + "/self-service/login/browser");
//       });
//   }, [router]);

//   if (!session) {
//     // Still loading
//     return null;
//   }

//   return (
//     <UserContext.Provider
//       value={{
//         username: getUsername(session.identity),
//         session: session,
//         logout_url: logoutUrl,
//       }}
//     >
//       {children}
//     </UserContext.Provider>
//   );
// };

// export default UserProvider;

// export const useUser = () => React.useContext(UserContext);
