export interface ApiError {
  status: number;
  message: string;
}

export function isApiError(obj: any): obj is ApiError {
  return obj && "status" in obj && "message" in obj;
}

export async function fetchJSON<T>(
  input: RequestInfo,
  init?: RequestInit
): Promise<T | ApiError | null> {
  const res = await fetch(input, init)
    .catch((e) => {
      console.error(e);
      return null;
    })
    .then((r) => r);

  if (!res) {
    console.error("failed to send the request");
    return null;
  }

  const data: { [key: string | number | symbol]: any } = await res
    .json()
    .catch((_) => ({ message: res.statusText }));

  // check for error response
  if (!res.ok) {
    // get error message from body or default to response statusText
    const error = {
      status: res.status,
      message: data.message,
    };

    return error;
  }

  return data as T;
}
