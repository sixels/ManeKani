import { useEffect, useState } from "react";
import { Stack, StackProps, Text, TextProps } from "@chakra-ui/react";

import MarkdownParser from "@/lib/markdown";

export type MarkdownProps = {
  title: string;
  content: string;
  props?: StackProps;
  titleProps?: TextProps;
  valueProps?: TextProps;
};

export function Markdown({
  title,
  content,
  props,
  titleProps,
  valueProps,
}: MarkdownProps) {
  const [markdown, setMarkdown] = useState("");
  const [isLoadingMarkdow, setLoadingMarkdow] = useState(false);

  useEffect(() => {
    setLoadingMarkdow(true);
    MarkdownParser.toHTML(content).then((parsed) => {
      setMarkdown(parsed);
      setLoadingMarkdow(false);
    });
  }, []);

  if (isLoadingMarkdow) {
    return <Text>Loading {title}...</Text>;
  }
  if (!markdown) {
    return null;
  }

  return (
    <Stack direction={"column"} {...props}>
      <Text
        as="span"
        fontWeight={"bold"}
        color={"GrayText"}
        textTransform={"uppercase"}
        {...titleProps}
      >
        Mnemonic
      </Text>
      <Text
        whiteSpace={"normal"}
        dangerouslySetInnerHTML={{ __html: markdown }}
        {...valueProps}
      />
    </Stack>
  );
}
