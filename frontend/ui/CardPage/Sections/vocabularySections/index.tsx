import { PartialKanjiResponse, Vocabulary } from "@/lib/models/cards";
import {
  Box,
  Flex,
  Heading,
  HStack,
  LinkBox,
  LinkOverlay,
  Stack,
  StackDivider,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Tabs,
  Text,
  VStack,
} from "@chakra-ui/react";
import * as wanakana from "wanakana";

import { SectionProps } from "..";
import { ListProps } from "../components/List";
import { PlayIcon } from "@radix-ui/react-icons";

export const vocabularySections = (
  vocabulary: Vocabulary,
  kanjis?: PartialKanjiResponse[]
): SectionProps[] => {
  // const patterns = vocabulary.patterns ? vocabulary.patterns.map((pattern) => ([pattern.name, pattern.sentences])) : null

  let kanjisSection: SectionProps = {} as SectionProps;
  if (kanjis && kanjis.length) {
    kanjisSection = {
      title: "Kanji composition",
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
                value: [vocabulary.name, ...(vocabulary.alt_names || [])].join(
                  ", "
                ),
              },
              {
                title: "word type",
                value: vocabulary.word_type.join(", "),
              },
            ],
          },
        },
        {
          mdDoc: {
            title: "Mnemonic",
            content: vocabulary.meaning_mnemonic,
          },
        },
      ],
    },
    {
      title: "Reading",
      sectionItems: [
        {
          horizontalList: {
            list: [vocabulary.reading, ...(vocabulary.alt_readings || [])].map(
              (reading) => ({
                value: (
                  <HStack gap={3} flexWrap={"wrap"}>
                    <VStack spacing={0} pl={5}>
                      <Text as="span" fontSize="xl">
                        {reading}
                      </Text>
                      <Text as="span" color={"gray.600"} fontSize="sm">
                        {wanakana.toRomaji(reading)}
                      </Text>
                    </VStack>
                    <Flex
                      gap={1}
                      as="span"
                      align="center"
                      fontSize={"xs"}
                      flexWrap={"wrap"}
                    >
                      <PlayIcon />
                      <Text as="span" fontWeight="bold" fontSize={"sm"}>
                        KYOKO
                      </Text>
                      <>(TOKYO ACCENT, FEMALE)</>
                    </Flex>
                    <Flex
                      gap={1}
                      as="span"
                      fontSize={"xs"}
                      align="center"
                      flexWrap={"wrap"}
                    >
                      <PlayIcon />
                      <Text as="span" fontWeight="bold" fontSize={"sm"}>
                        KENICH
                      </Text>
                      <>(TOKYO ACCENT, MALE)</>
                    </Flex>
                  </HStack>
                ),
                valueProps: {
                  fontSize: "xl",
                  textTransform: undefined,
                  fontWeight: "normal",
                  color: undefined,
                },
              })
            ),
          },
        },
        {
          mdDoc: {
            title: "Mnemonic",
            content: vocabulary.reading_mnemonic,
          },
        },
      ],
    },
    {
      title: "Context",
      sectionItems: [
        vocabulary.patterns
          ? {
              component: (
                <Tabs
                  variant="line"
                  my={5}
                  orientation="vertical"
                  defaultIndex={1}
                  isLazy
                >
                  <TabList
                    gap={2}
                    borderLeft="none"
                    borderRightWidth="2px"
                    px={1}
                  >
                    <Tab
                      key={0}
                      isDisabled
                      cursor={"default !important"}
                      userSelect="text"
                      opacity="1 !important"
                      pt={0}
                      pb={2}
                    >
                      <Heading
                        as="h3"
                        size={{ base: "sm", md: "md" }}
                        color="GrayText"
                        whiteSpace="nowrap"
                        fontWeight={"normal"}
                        textTransform="capitalize"
                        cursor={"text"}
                      >
                        Patterns of Use
                      </Heading>
                    </Tab>
                    {vocabulary.patterns.map((pattern, i) => (
                      <Tab
                        key={i + 1}
                        bg="white"
                        border="none"
                        rounded="md"
                        _selected={{
                          bgColor: "gray.700",
                          color: "white",
                        }}
                      >
                        {pattern.name}
                      </Tab>
                    ))}
                  </TabList>
                  <TabPanels>
                    <TabPanel />
                    {vocabulary.patterns.map((pattern, i) => (
                      <TabPanel key={i + 1} py={0} pl={3} pr={2}>
                        <Heading
                          as="h3"
                          size={{ base: "sm", md: "md" }}
                          color="GrayText"
                          fontWeight={"normal"}
                          textTransform="capitalize"
                        >
                          Common word combinations
                        </Heading>
                        {pattern.sentences.map((sentence, i) => (
                          <VStack key={i} py="4" spacing={0} alignItems="start">
                            <Text color="gray.900" lang="jp">
                              {sentence.sentence}
                            </Text>
                            <Text color="gray.700">{sentence.meaning}</Text>
                          </VStack>
                        ))}
                      </TabPanel>
                    ))}
                  </TabPanels>
                </Tabs>
              ),
            }
          : ({} as any),
        vocabulary.sentences
          ? {
              component: (
                <Box>
                  <Heading
                    as="h4"
                    size="md"
                    py={3}
                    color="GrayText"
                    textTransform="capitalize"
                  >
                    Sentences
                  </Heading>
                  <Stack
                    align="flex-start"
                    spacing={6}
                    divider={<StackDivider w={"80%"} />}
                  >
                    {vocabulary.sentences.map((sentence, i) => (
                      <VStack align="flex-start" spacing={0}>
                        <Text color="gray.900" lang={"jp"}>
                          {sentence.sentence}
                        </Text>
                        <Text color="gray.700">{sentence.meaning}</Text>
                      </VStack>
                    ))}
                  </Stack>
                </Box>
              ),
            }
          : ({} as any),
      ],
    },
    { ...kanjisSection },
  ];
};
