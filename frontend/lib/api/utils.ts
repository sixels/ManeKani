export async function fetchJSON<T>(
  input: RequestInfo,
  init?: RequestInit,
): Promise<T> {
  const res = await fetch(input, init)
    .catch((e) => {
      console.log(e);
      return null;
    })
    .then((r) => r);

  if (!res) {
    return Promise.reject('');
  }
  const data = await res.json();

  // check for error response
  if (!res.ok) {
    // get error message from body or default to response statusText
    const error = new Error(
      data && 'message' in data ? data.message : res.statusText,
    );
    return Promise.reject(error);
  }

  return data;
}
