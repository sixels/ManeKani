import useSWR from "swr";
import { API_URL } from "../api/fetchApi";
import { SessionQueue, SRSUserData } from "../models/srs";

export const WANIKANI_DECK_ID = "9f9b7a42-1152-4fc5-87e1-8990e38871b0";

export function useSRS() {
  // TODO: fetch wanikani deck id and store in a client state
  const { data, isLoading, error, mutate } = useSWR<SRSUserData["cards"]>(
    `${API_URL}/api/v1/card?decks=${WANIKANI_DECK_ID}`,
    {
      errorRetryCount: 0,
    }
  );

  return {
    data,
    isLoading,
    isError: error,
    mutate,
  };
}

export function useSRSSession(session_type: "lesson" | "review") {
  const { data, isLoading, error } = useSWR<SessionQueue>(
    `${API_URL}/study-session?deck=${WANIKANI_DECK_ID}&session_type=${session_type}`,
    {
      errorRetryCount: 0,
    }
  );

  return {
    session: data,
    isLoading,
    isError: error,
  };
}
