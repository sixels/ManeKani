// import React, { useEffect } from 'react';

// type State =
//   | {
//       tokens: TokenGetAllResponseInner[];
//     }
//   | {
//       error: unknown;
//     }
//   | null;

// export const TokensContext = React.createContext<{
//   state: State;
//   createToken: (
//     name: string,
//     permissions: TokenClaims,
//   ) => Promise<TokenCreateResponse>;
//   deleteToken: (id: string) => Promise<void>;
// }>({
//   state: null,
//   createToken: () => Promise.reject('tokens context not ready'),
//   deleteToken: () => Promise.reject('tokens context not ready'),
// });

// export const TokensProvider: React.FC<React.PropsWithChildren> = ({
//   children,
// }) => {
//   const [state, setState] = React.useState<State>(null);

//   const API = new TokenApi(new Configuration({ credentials: 'include' }));

//   const fetchTokens = async () => await API.getTokens();

//   const refetchTokens = () => {
//     console.log('refetching tokens');

//     fetchTokens()
//       .then((res) => {
//         console.info(res.data[0].claims);
//         if ('data' in res) {
//           return setState({ tokens: res.data });
//         }
//         setState({ tokens: [] });
//       })
//       .catch((e) => {
//         console.warn(e);
//         setState({ error: e });
//       });
//   };

//   const createToken = async (name: string, permissions: TokenClaims) => {
//     return API.createToken({ request: { name, permissions } }).then((r) => {
//       refetchTokens();
//       return r.data;
//     });
//   };

//   const deleteToken = async (id: string) => {
//     return API.deleteToken({ id }).then(() => {
//       refetchTokens();
//     });
//   };

//   useEffect(() => {
//     refetchTokens();
//   }, []);

//   return (
//     <TokensContext.Provider value={{ state, createToken, deleteToken }}>
//       {children}
//     </TokensContext.Provider>
//   );
// };

// export const useTokens = () => React.useContext(TokensContext);
