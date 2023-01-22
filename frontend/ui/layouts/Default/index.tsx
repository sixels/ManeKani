import Footer from "@/ui/Footer";
import Navbar from "@/ui/Navbar";
import { PropsWithChildren } from "react";

import styles from "./default.module.css";

export default function Layout({ children }: PropsWithChildren) {
  return (
    <>
      <Navbar />
      <main className={styles.pageContent}> {children} </main>
      <Footer />
    </>
  );
}
