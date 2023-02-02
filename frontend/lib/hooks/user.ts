import useSWR from "swr";
import { API_URL } from "../api/fetchApi";

export interface UserData {
  email: string;
  username: string;
  level: number;
}

export function useUser() {
  const { data, isLoading, error } = useSWR<UserData, any>(`${API_URL}/user`);

  return {
    user: data,
    isLoading,
    isError: error,
  };
}
