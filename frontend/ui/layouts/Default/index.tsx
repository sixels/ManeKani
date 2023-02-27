import Footer from "@/ui/Footer";
import Navbar from "@/ui/Navbar";
import { Grid } from "@chakra-ui/react";
import { PropsWithChildren } from "react";

import styles from "./default.module.css";

export default function Layout({ children }: PropsWithChildren) {
  return (
    <>
      <Navbar />
      <Grid gridTemplateRows={"1fr auto"} mt="60px">
        <main> {children} </main>
        <Footer />
      </Grid>
    </>
  );
}
