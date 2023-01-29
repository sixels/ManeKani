import { API_URL } from "@/lib/api/fetchApi";
import { cardColors } from "@/lib/theme";
import {
  Card,
  CardBody,
  CardHeader,
  Center,
  HStack,
  Progress,
  Stack,
  Text,
  VStack,
} from "@chakra-ui/react";
import Image from "next/image";

export type LevelProgressProps = {
  userLevel: number;
  completeCardProgress: number;
  cards: {
    name: string;
    value: string | { image: string };
    meaning: string;
    kind: "radical" | "vocabulary" | "kanji";
    progress: number;
  }[];
};
export function LevelProgress({
  userLevel,
  completeCardProgress,
  cards,
}: LevelProgressProps) {
  const radicals = cards.filter((c) => c.kind == "radical");
  const kanjis = cards.filter((c) => c.kind == "kanji");
  // const vocabs = cards.filter(c => c.kind == "vocabulary")

  // we don't consider vocabularies to pass levels
  const totalCards = radicals.length + kanjis.length;

  const kanjisDone = kanjis.reduce((tot, { progress }) => {
    return tot + (progress >= completeCardProgress ? 1 : 0);
  }, 0);

  let totalProgress = kanjis.reduce((tot, { progress }) => {
    return tot + (progress >= completeCardProgress ? 1 : 0);
  }, 0);

  totalProgress = clamp(
    totalCards == 0 ? 0 : totalProgress / kanjis.length,
    0.0,
    1.0
  );

  // apply mask when `progress != 1.0`
  const progressMask =
    totalProgress > 0 && totalProgress < 1.0
      ? `
    radial-gradient(7.25px at calc(100% - 10.25px) 50%,#000 99%,#0000 101%) 0 calc(50% - 10px)/100% 20px,
    radial-gradient(7.25px at calc(100% + 5.25px) 50%,#0000 99%,#000 101%) calc(100% - 5px) 50%/100% 20px repeat-y;
  `
      : undefined;

  return (
    <Card shadow="sm" bg="white" rounded="2xl">
      <CardHeader
        position="relative"
        zIndex={1}
        roundedTop="2xl"
        overflow="hidden"
        _before={{
          content: `""`,
          zIndex: -1,
          bgColor: "blue.100",
          w: `${totalProgress * 100}%`,
          h: "100%",
          position: "absolute",
          top: 0,
          left: 0,
          mask: progressMask,
        }}
        _after={{
          content: `""`,
          zIndex: -1,
          bgColor: "blue.200",
          w: `calc(${totalProgress * 100}% - ${
            kanjis.length - kanjisDone == 0 ? 0 : 10
          }px)`,
          h: "100%",
          position: "absolute",
          top: 0,
          left: 0,
          mask: progressMask,
        }}
      >
        <HStack justifyContent="space-between">
          <Text textTransform="capitalize" fontSize="lg">
            Level {userLevel} Progress
          </Text>
          <Text as="span">
            {kanjisDone} / {kanjis.length} kanji
          </Text>
        </HStack>
      </CardHeader>
      <CardBody>
        <Stack align="flex-start" spacing={6}>
          <VStack align="flex-start">
            <Text fontWeight={"medium"}>Radicals</Text>
            <HStack flexWrap={"wrap"} gap={2} rowGap={3} spacing={0}>
              {radicals.map((r, i) => (
                <CardProgress key={i} card={r} />
              ))}
            </HStack>
          </VStack>
          <VStack align="flex-start">
            <Text fontWeight={"medium"}>Kanji</Text>
            <HStack flexWrap={"wrap"} gap={2} rowGap={3} spacing={0}>
              {kanjis.map((r, i) => (
                <CardProgress key={i} card={r} />
              ))}
            </HStack>
          </VStack>
        </Stack>
      </CardBody>
    </Card>
  );
}

function CardProgress({ card }: { card: LevelProgressProps["cards"][number] }) {
  return (
    <Stack spacing={1} align="center">
      <Center
        w={10}
        h={10}
        bg={cardColors[card.kind]}
        color="white"
        rounded={"lg"}
      >
        {typeof card.value == "string" ? (
          card.value
        ) : (
          <Image
            src={`${API_URL}/files/image/${card.value}`}
            alt={`${card.meaning} progress`}
          />
        )}
      </Center>
      <Progress
        colorScheme={cardColors[card.kind].split(".")[0]}
        w="85%"
        size={"xs"}
        value={(card.progress / 5) * 100}
        rounded="sm"
        hasStripe
      />
    </Stack>
  );
}

function clamp(value: number, min: number, max: number): number {
  return Math.max(min, Math.min(max, value));
}
