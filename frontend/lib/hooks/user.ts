import useSWR from "swr";
import { API_URL } from "../api/fetchApi";

export interface UserData {
  email: string;
  username: string;
  level: number;
}

export function useUser() {
  const { data, isLoading, error } = useSWR<UserData>(`${API_URL}/user`, {
    shouldRetryOnError: false,
    revalidateOnFocus: false,
    revalidateIfStale: false,
  });

  return {
    user: data,
    isLoading,
    isError: error,
  };
}
