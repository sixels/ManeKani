import { Box, Center, Container, Heading, Stack, Text } from "@chakra-ui/react";

import { Section, SectionProps } from "@/ui/Subject/Section";
import { RemoteContent } from "@/lib/models/subject";
import { API_URL } from "@/lib/api/fetchApi";
import SVG from "../SVG";

export type CardPageProps = {
  card: {
    kind: string;
    name: string;
    level: number;
    value?: string;
    value_image?: RemoteContent;
    reading?: string;
  };
  sections: SectionProps[];
};

export default function CardPage({ card, sections }: CardPageProps) {
  function headingFontSize(text: string) {
    const x = text.length;
    return Math.max(1, ((1 / 4) * x) ** (1 / 2) - x / 2.3 + 2.58);
  }

  const highlightColor: { [key: string]: string } = {
    radical: "blue.500",
    kanji: "pink.500",
    vocabulary: "purple.500",
  };

  return (
    <Container maxW={"7xl"}>
      <Stack spacing={8} pt={0} align="center" position={"relative"}>
        <Center position={"relative"}>
          <Box
            position={"absolute"}
            width={"120%"}
            height={"60%"}
            transform="translate(-50%,-50%)"
            top={"50%"}
            left={"50%"}
            backgroundColor={
              (card.kind in highlightColor && highlightColor[card.kind]) ||
              "orange.500"
            }
            zIndex={-1}
            rounded={"2xl"}
            filter="auto"
            blur="2xl"
            opacity={"23%"}
          ></Box>
          <Stack
            spacing={1}
            direction={"column"}
            backgroundColor="white"
            shadow={"lg"}
            rounded={"2xl"}
            border={"2px"}
            borderColor={"gray.200"}
            px="10"
            w={"72"}
            py="5"
            align="center"
            fontSize={"1.5em"}
          >
            <Text textTransform={"uppercase"} fontSize="md" fontWeight="bold">
              level {card.level} {card.kind}
            </Text>
            {card.value ? (
              <Heading
                as="h1"
                fontSize={`${headingFontSize(card.value)}em`}
                noOfLines={1}
                lineHeight={"15rem"}
                lang="ja"
              >
                {card.value}
              </Heading>
            ) : (
              <Center height={"142px"} w="65%" py="8.458rem">
                <SVG
                  url={`${API_URL}/files/${card.value_image?.url}`}
                  style={{
                    width: "100%",
                    fill: "none",
                    stroke: "currentcolor",
                    strokeWidth: "60px",
                  }}
                />
              </Center>
            )}
            <Text
              textTransform={"capitalize"}
              textAlign="center"
              fontSize={"xl"}
            >
              {card.name}
            </Text>
            {card.reading && (
              <Text
                textTransform={"capitalize"}
                fontWeight="normal"
                fontSize={"lg"}
              >
                {card.reading}
              </Text>
            )}
          </Stack>
        </Center>
        {sections
          .filter((s) => s.sectionItems)
          .map((section, i) => {
            return (
              <Section
                key={i}
                sectionItems={section.sectionItems}
                title={section.title}
              />
            );
          })}
      </Stack>
    </Container>
  );
}
