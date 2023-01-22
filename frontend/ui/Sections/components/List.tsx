import { Flex, FlexProps, Text, TextProps } from "@chakra-ui/react";

export type ListProps = {
  title?: string;
  value: string | JSX.Element;
  valueProps?: Partial<TextProps>;
  titleProps?: Partial<TextProps>;
};

export type HorizontalListProps = {
  list: ListProps[];
  listProps?: FlexProps;
  itemProps?: FlexProps;
};

export function HorizontalList({
  list,
  listProps,
  itemProps,
}: HorizontalListProps) {
  const items = list.map(({ title, value, valueProps, titleProps }, i) => {
    return (
      <Flex
        direction={"column"}
        gap={"10px"}
        align={"center"}
        py={4}
        key={i}
        {...itemProps}
      >
        {title ? (
          <Text
            as="span"
            color="GrayText"
            textTransform={"uppercase"}
            fontWeight={"bold"}
            {...titleProps}
          >
            {title}
          </Text>
        ) : (
          <></>
        )}
        <Text as="div" color={"GrayText"} {...valueProps}>
          {value}
        </Text>
      </Flex>
    );
  });

  return (
    <Flex direction={"row"} {...listProps}>
      {items}
    </Flex>
  );
}
