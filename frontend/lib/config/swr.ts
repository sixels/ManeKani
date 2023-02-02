import { SWRConfiguration } from "swr";

const config: SWRConfiguration = {
  fetcher: async (url, init) => {
    const res = await fetch(url, init).then((r) => r);

    const data = await res.json().catch((_) => ({ message: res.statusText }));
    if (!res.ok) {
      // get error message from body or default to response statusText
      const error = {
        status: res.status,
        message: data.message,
      };
      throw error;
    }

    return data;
  },
};

export default config;
