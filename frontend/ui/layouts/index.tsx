import { NextPage } from "next";
import { ReactElement, ReactNode } from "react";

export type PageWithLayout<P = {}, IP = P> = NextPage<P, IP> & {
  getLayout?: (page: ReactElement) => ReactNode;
};
