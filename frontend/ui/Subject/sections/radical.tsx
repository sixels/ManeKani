import { PartialKanjiResponse, Radical, Vocabulary } from "@/lib/models/cards";
import { PartialSubject, StudyData, Subject } from "@/lib/models/subject";
import { Flex, LinkBox, LinkOverlay, Text, VStack } from "@chakra-ui/react";
import { SectionProps } from "../Section";
import { studyData } from "./common";

export const radicalSections = (
  radical: Subject,
  kanjis?: PartialSubject[]
): SectionProps[] => {
  let kanjisSection: SectionProps = {} as SectionProps;

  const meaningMnemonic = studyData(radical, "meaning");

  if (kanjis && kanjis.length) {
    kanjisSection = {
      title: "Found in kanji",
      sectionItems: [
        {
          component: (
            <VStack mt={2} gap={1}>
              {kanjis.map((kanji, i) => (
                <LinkBox key={i} w="full">
                  <LinkOverlay href={`/kanji/${kanji.value}`}>
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
                        {kanji.value}
                      </Text>
                      <VStack lineHeight={"1em"} align="end">
                        <Text as="span" lang="ja">
                          <>
                            {"    "}
                            {
                              studyData(kanji, "reading")?.items.find(
                                (item) => item.is_primary
                              )?.value
                            }
                            {"    "}
                          </>
                        </Text>
                        <Text as="span">{kanji.name}</Text>
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
            content: meaningMnemonic?.mnemonic || "",
          },
        },
      ],
    },
    { ...kanjisSection },
  ];
};
