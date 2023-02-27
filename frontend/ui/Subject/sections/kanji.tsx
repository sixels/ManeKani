import { API_URL } from "@/lib/api/fetchApi";

import { PartialSubject, Subject } from "@/lib/models/subject";
import SVG from "@/ui/SVG";
import {
  Flex,
  Img,
  LinkBox,
  LinkOverlay,
  Text,
  VStack,
} from "@chakra-ui/react";

import { SectionProps } from "../Section";
import { ListProps } from "../Section/components/List";
import { studyData, studyItem, studyItems } from "./common";

export const kanjiSections = (
  kanji: Subject,
  vocabularies?: PartialSubject[],
  radicals?: PartialSubject[]
): SectionProps[] => {
  const primaryReading = studyItem(
    kanji,
    "reading",
    (si) => si.is_primary
  )?.value;

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
              {vocabularies.map(({ name, value: word, ...vocab }, i) => {
                const reading =
                  studyItem(vocab, "reading", (si) => si.is_primary)?.value ||
                  "-";
                return (
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
                );
              })}
            </VStack>
          ),
        },
      ],
    };
  }

  let radicalsSection: SectionProps = {} as SectionProps;
  if (radicals && radicals.length) {
    radicalsSection = {
      title: "Radical composition",
      sectionItems: [
        {
          horizontalList: {
            list: radicals.map((rad) => ({
              title: rad.name,
              value: rad.value ? (
                rad.value
              ) : (
                <SVG
                  url={`${API_URL}/files/${rad.value_image?.url}`}
                  style={{
                    height: "58px",
                    width: "50px",
                  }}
                />
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
    };
  }

  const meanings =
    studyData(kanji, "meaning")?.items.map((si) => si.value) || [];
  const meaningMnemonic = studyData(kanji, "meaning")?.mnemonic;

  const onyomi =
    studyItems(kanji, "reading", (si) => si.category == "onyomi")?.map(
      (si) => si.value
    ) || [];
  const kunyomi =
    studyItems(kanji, "reading", (si) => si.category == "kunyomi")?.map(
      (si) => si.value
    ) || [];
  const nanori =
    studyItems(kanji, "reading", (si) => si.category == "nanori")?.map(
      (si) => si.value
    ) || [];
  const readingMnemonic = studyData(kanji, "reading")?.mnemonic;

  return [
    { ...radicalsSection },
    {
      title: "Meaning",
      sectionItems: [
        {
          headlessTable: {
            table: [
              {
                title: "meanings",
                value: meanings.join(", "),
              },
            ],
          },
        },
        {
          mdDoc: {
            title: "Mnemonic",
            content: meaningMnemonic || "",
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
                value: onyomi.length ? onyomi.join(", ") : "None",
                ...styleReading(onyomi),
              },
              {
                title: "Kun'yomi",
                value: kunyomi.length ? kunyomi.join(", ") : "None",
                ...styleReading(kunyomi),
              },
              {
                title: "Nanori",
                value: nanori.length ? nanori.join(", ") : "None",
                ...styleReading(nanori),
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
            content: readingMnemonic || "",
          },
        },
      ],
    },
    { ...vocabulariesSection },
  ];
};
