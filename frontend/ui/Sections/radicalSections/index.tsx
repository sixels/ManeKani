import { PartialKanjiResponse, Radical, Vocabulary } from "@/lib/models/cards";
import { Flex, LinkBox, LinkOverlay, Text, VStack } from "@chakra-ui/react";
import { SectionProps } from "..";
import { ListProps } from "../components/List";

export const radicalSections = (
  radical: Radical,
  kanjis?: PartialKanjiResponse[]
): SectionProps[] => {
  let kanjisSection: SectionProps = {} as SectionProps;
  if (kanjis && kanjis.length) {
    kanjisSection = {
      title: "Found in kanji",
      sectionItems: [
        {
          component: (
            <VStack mt={2} gap={1}>
              {kanjis.map(({ name, reading, symbol }, i) => (
                <LinkBox key={i} w="full">
                  <LinkOverlay href={`/kanji/${symbol}`}>
                    <Flex
                      alignItems="center"
                      px={3}
                      py={2}
                      bg="pink.500"
                      color="pink.50"
                      rounded="md"
                      shadow={"md"}
                    >
                      <Text as="span" lang="ja" flex={1}>
                        {symbol}
                      </Text>
                      <VStack lineHeight={"1em"} align="end">
                        <Text as="span" lang="ja">
                          {"    "}
                          {reading}
                          {"    "}
                        </Text>
                        <Text as="span">{name}</Text>
                      </VStack>
                    </Flex>
                  </LinkOverlay>
                </LinkBox>
              ))}
            </VStack>
          ),
        },
      ],
    };
  }

  return [
    {
      title: "Meaning",
      sectionItems: [
        {
          headlessTable: {
            table: [
              {
                title: "meanings",
                value: radical.name,
              },
            ],
          },
        },
        {
          mdDoc: {
            title: "Mnemonic",
            content: radical.meaning_mnemonic,
          },
        },
      ],
    },
    { ...kanjisSection },
  ];
};
