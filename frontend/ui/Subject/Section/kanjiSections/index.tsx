import { API_URL } from "@/lib/api/fetchApi";
import {
  Kanji,
  PartialRadicalResponse,
  PartialVocabularyResponse,
} from "@/lib/models/cards";
import { PartialSubject, Subject } from "@/lib/models/subject";
import {
  Box,
  Flex,
  Img,
  LinkBox,
  LinkOverlay,
  StackDivider,
  Text,
  VStack,
} from "@chakra-ui/react";
import { SectionProps } from "../../Section";
import { ListProps } from "../../Section/components/List";

export const kanjiSections = (
  kanji: Subject,
  vocabularies?: PartialSubject[],
  radicals?: PartialSubject[]
): SectionProps[] => {
  const primaryReading = kanji.study_data
    .find((sd) => sd.kind == "reading")
    ?.items.find((si) => si.is_primary)?.value;

  function readingIncludes(readings: string[]): boolean {
    return primaryReading != undefined
      ? readings.includes(primaryReading)
      : false;
  }

  function styleReading(readings: string[]): Partial<ListProps> {
    const valueProps = (readingIncludes(readings) && {
      fontWeight: "bold",
      color: "",
      lang: "ja",
    }) || { lang: "ja" };

    return {
      valueProps,
      titleProps:
        (readingIncludes(readings) && {
          color: "",
        }) ||
        undefined,
    };
  }

  let vocabulariesSection: SectionProps = {} as SectionProps;
  if (vocabularies && vocabularies.length) {
    vocabulariesSection = {
      title: "Found in vocabulary",
      sectionItems: [
        {
          component: (
            <VStack mt={2} gap={1}>
              {vocabularies.map(({ name, reading, word }, i) => (
                <LinkBox key={i} w="full">
                  <LinkOverlay href={`/vocabulary/${word}`}>
                    <Flex
                      alignItems="center"
                      px={3}
                      py={2}
                      bg="purple.500"
                      color="purple.50"
                      rounded="md"
                      shadow={"md"}
                    >
                      <Text as="span" lang="ja" flex={1}>
                        {word}
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
    radicals && radicals.length
      ? {
          title: "Radical composition",
          sectionItems: [
            {
              horizontalList: {
                list: radicals.map((rad) => ({
                  title: rad.name,
                  value: rad.symbol.includes("/") ? (
                    <Img
                      src={`${API_URL}/files/${rad.symbol}`}
                      alt={"radical symbol"}
                      h={"58px"}
                      w={"50px"}
                      filter="brightness(0.12) "
                    />
                  ) : (
                    rad.symbol
                  ),
                  valueProps: {
                    color: undefined,
                    fontSize: "5xl",
                    fontWeight: "normal",
                    flex: 1,
                  },
                  titleProps: { color: undefined, fontWeight: "bold" },
                })),
                itemProps: {
                  bgColor: "white",
                  minWidth: 32,
                  rounded: "md",
                  borderWidth: 1,
                  boxShadow: "md",
                  position: "relative",
                  h: 40,
                  justifyContent: "space-around",
                  _before: {
                    content: '""',
                    w: "90%",
                    h: "70%",
                    top: "50%",
                    left: "50%",
                    rounded: "lg",
                    filter: "auto",
                    blur: "lg",
                    opacity: "26%",
                    zIndex: -1,
                    transform: "translate(-50%,-50%)",
                    position: "absolute",
                    bgColor: "blue.500",
                  },
                },
                listProps: {
                  pt: 3,
                  justifyContent: "left",
                  gap: "4",
                  flexWrap: "wrap",
                },
              },
            },
          ],
        }
      : ({} as any),
    {
      title: "Meaning",
      sectionItems: [
        {
          headlessTable: {
            table: [
              {
                title: "meanings",
                value: [kanji.name, ...(kanji.alt_names || [])].join(", "),
              },
            ],
          },
        },
        {
          mdDoc: {
            title: "Mnemonic",
            content: kanji.meaning_mnemonic,
          },
        },
      ],
    },
    {
      title: "Reading",
      sectionItems: [
        {
          horizontalList: {
            list: [
              {
                title: "On'yomi",
                value:
                  (kanji.onyomi.length && kanji.onyomi.join(", ")) || "None",
                ...styleReading(kanji.onyomi),
              },
              {
                title: "Kun'yomi",
                value:
                  (kanji.kunyomi.length && kanji.kunyomi.join(", ")) || "None",
                ...styleReading(kanji.kunyomi),
              },
              {
                title: "Nanori",
                value:
                  (kanji.nanori.length && kanji.nanori.join(", ")) || "None",
                ...styleReading(kanji.nanori),
              },
            ],
            listProps: {
              justify: "space-around",
            },
          },
        },
        {
          mdDoc: {
            title: "Mnemonic",
            content: kanji.reading_mnemonic,
          },
        },
      ],
    },
    { ...vocabulariesSection },
  ];
};
