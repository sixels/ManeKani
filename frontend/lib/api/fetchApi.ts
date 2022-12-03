import { fetchJSON } from './utils';

export async function fetchApi<T>(
  endpoint: string,
  opts?: RequestInit,
): Promise<T> {
  return await fetchJSON(`${process.env.API_URL}/api/${endpoint}`, opts);
}
