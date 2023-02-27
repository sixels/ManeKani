// import "../styles/globals.css";
import React from "react";
import { AppProps } from "next/app";
import SuperTokensReact, { SuperTokensWrapper } from "supertokens-auth-react";
import { ChakraProvider } from "@chakra-ui/react";
import Head from "next/head";

import { frontendConfig } from "@/lib/supertokens/frontendConfig";
import SEO from "@/lib/config/seo";
import swrConfiguration from "@/lib/config/swr";
import { theme } from "@/lib/theme";
import { SWRConfig } from "swr";
import { PageWithLayout } from "@/ui/layouts";
import DefaultLayout from "@/ui/layouts/Default";

import Footer from "@/ui/Footer";
import Navbar from "@/ui/Navbar";

import Favicon from "@/assets/icon.svg";

import "@fontsource/inter/400.css";
import "@fontsource/inter/500.css";
import "@fontsource/inter/700.css";
import "@fontsource/noto-sans-jp/400.css";
import "@/assets/styles/calendar.css";
import "@/assets/styles/markdown.css";

type AppPropsWithLayout = AppProps & {
  Component: PageWithLayout;
};

if (typeof window !== "undefined") {
  SuperTokensReact.init(frontendConfig());
}

function App({ Component, pageProps }: AppPropsWithLayout) {
  const getLayout =
    Component.getLayout ?? ((page) => <DefaultLayout>{page}</DefaultLayout>);

  return (
    <>
      <Head>
        <title>{SEO.title}</title>
        <link rel="manifest" href="/manifest.json" />
        <link rel="icon" href={Favicon.src} />
        <meta charSet="utf-8" />
        <meta name="theme-color" content={SEO.themeColor} />
        <meta name="description" content={SEO.description} />
        <meta property="og:description" content={SEO.description} />
        <meta property="og:site_name" content={SEO.title} />
        <meta property="og:image" content={Favicon.src} />
      </Head>

      <ChakraProvider theme={theme}>
        <SuperTokensWrapper>
          <SWRConfig value={swrConfiguration}>
            {getLayout(<Component {...pageProps} />)}
          </SWRConfig>
        </SuperTokensWrapper>
      </ChakraProvider>
    </>
  );
}

export default App;
