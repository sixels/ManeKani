import { join } from "path";
import { fetchJSON, isApiError } from "./utils";

export const API_URL = process.env.API_URL || "http://192.168.15.9:8081";

export async function fetchApi<T = any>(
  endpoint: string,
  version: "v1",
  opts?: RequestInit
): Promise<T | null> {
  const resourceUrl = `${API_URL}${join("/api", version, endpoint)}`;

  const data = await fetchJSON<T>(resourceUrl, opts);

  if (isApiError(data)) {
    console.error(
      `fetch api error: ${opts?.method || "GET"} ${resourceUrl} ${
        data.status
      }: ${data.message}`
    );
    return null;
  }

  return data;
}
