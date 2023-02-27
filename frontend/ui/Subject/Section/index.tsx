import { Flex, Heading, Stack, StackDivider } from "@chakra-ui/react";
import { HorizontalList, HorizontalListProps } from "./components/List";
import { Markdown, MarkdownProps } from "./components/Markdown";
import { HeadlessTable, HeadlessTableProps } from "./components/Table";

export interface SectionItem {
  component?: JSX.Element;
  headlessTable?: HeadlessTableProps;
  horizontalList?: HorizontalListProps;
  mdDoc?: MarkdownProps;
}

export type SectionProps = {
  title: string;
  sectionItems: SectionItem[];
};

export function Section({ title, sectionItems }: SectionProps) {
  return (
    <Flex as="section" w="full" py={5}>
      <Stack flex={1}>
        <Heading as="h2" size={"lg"} textTransform={"capitalize"}>
          {title}
        </Heading>
        <Stack divider={<StackDivider />}>
          {sectionItems.map(renderSectionItem)}
        </Stack>
      </Stack>
    </Flex>
  );
}

function renderSectionItem(item: SectionItem, key: number) {
  if (item.component) {
    return <div key={key}>{item.component}</div>;
  }
  if (item.headlessTable) {
    return <HeadlessTable key={key} {...item.headlessTable} />;
  }
  if (item.horizontalList) {
    return <HorizontalList key={key} {...item.horizontalList} />;
  }
  if (item.mdDoc) {
    return <Markdown key={key} {...item.mdDoc} />;
  }
  return null;
}
