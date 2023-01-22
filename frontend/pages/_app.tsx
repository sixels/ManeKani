// import "../styles/globals.css";
import React from "react";
import { AppProps } from "next/app";
import SuperTokensReact, { SuperTokensWrapper } from "supertokens-auth-react";
import { ChakraProvider } from "@chakra-ui/react";
import Head from "next/head";

import { frontendConfig } from "@/lib/supertokens/frontendConfig";
import { SEO } from "@/lib/config";

import Footer from "@/ui/Footer";
import Navbar from "@/ui/Navbar";

import Favicon from "@/assets/icon.svg";

import "./kanji/style.css";
import { UserDataProvider } from "@/lib/auth/context";

if (typeof window !== "undefined") {
  SuperTokensReact.init(frontendConfig());
}

function App({ Component, pageProps }: AppProps) {
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

      <ChakraProvider>
        <SuperTokensWrapper>
          <UserDataProvider>
            <Navbar />
            <Component {...pageProps} />
            <Footer />
          </UserDataProvider>
        </SuperTokensWrapper>
      </ChakraProvider>
    </>
  );
}

export default App;
