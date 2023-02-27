import { Spinner } from "@chakra-ui/react";
import { CSSProperties, useEffect, useState } from "react";

export interface SVGProps {
  url: string;
  style: CSSProperties;
}
export default function SVG({ url, style }: SVGProps) {
  const [isLoading, setIsLoading] = useState(false);
  const [content, setContent] = useState<string>();

  useEffect(() => {
    setIsLoading(true);
    try {
      fetch(url, { cache: "force-cache" })
        .then((res) => res.text())
        .then((content) => {
          setContent(content);
          setIsLoading(false);
        });
    } catch (e) {
      console.warn(e);
    }
  }, []);

  if (isLoading) {
    return <Spinner />;
  }

  if (!content) {
    return <>A</>;
  }

  return <div style={style} dangerouslySetInnerHTML={{ __html: content }} />;
}
