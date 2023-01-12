// import "../styles/globals.css";
import React from "react";
import { AppProps } from "next/app";
import SuperTokensReact, { SuperTokensWrapper } from "supertokens-auth-react";

import { frontendConfig } from "@/lib/supertokens/frontendConfig";

if (typeof window !== "undefined") {
  SuperTokensReact.init(frontendConfig());
}

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <SuperTokensWrapper>
      <Component {...pageProps} />
    </SuperTokensWrapper>
  );
}

export default MyApp;
