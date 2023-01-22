import React, { createContext, useContext, useEffect, useState } from "react";
import { useSessionContext } from "supertokens-auth-react/recipe/session";
import { API_URL } from "../api/fetchApi";
import { fetchJSON, isApiError } from "../api/utils";

interface UserInfo {
  loading: false;
  email: string;
  username: string;
}

type ProviderState =
  | { loading: false; user: UserInfo }
  | { loading: true }
  | { loading: false; user: null };

const Context = createContext<ProviderState>({
  loading: true,
});

export function UserDataProvider(props: React.PropsWithChildren) {
  const session = useSessionContext();

  const [userData, setData] = useState<ProviderState>({
    loading: true,
  });

  useEffect(() => {
    if (session.loading) {
      return;
    }

    if (session.doesSessionExist) {
      fetchJSON<UserInfo>(`${API_URL}/user`).then((data) => {
        if (data == null || isApiError(data)) {
          setData({ loading: false, user: null });
        } else {
          setData({ loading: false, user: data });
        }
      });
    } else {
      setData({ loading: false, user: null });
    }
  }, [session]);

  return <Context.Provider value={userData}>{props.children}</Context.Provider>;
}

export const useUserData = () => useContext(Context);
