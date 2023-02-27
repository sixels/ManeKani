import {
  Container,
  SimpleGrid,
  Box,
  Grid,
  Spinner,
  Text,
  VStack,
  Heading,
  Flex,
} from "@chakra-ui/react";

import AuthRoute from "@/lib/auth/wrappers/AuthRoute";
import ReviewForecast from "@/ui/Dashboard/ReviewForecast";
import LinkCard from "@/ui/Dashboard/LinkCard";

import ExtraStudy from "@/ui/Dashboard/ExtraStudy";
import { LevelProgress } from "@/ui/Dashboard/LevelProgress";
import { API_URL } from "@/lib/api/fetchApi";
import { useSRS, WANIKANI_DECK_ID } from "@/lib/hooks/srs";
import StudyModal from "@/ui/Dashboard/StudyModal";
import { differenceInHours } from "date-fns";

type CardProgress = {
  kind: "radical" | "kanji" | "vocabulary";
  name: string;
  value: string | { image: string };
  meaning: string;
  progress: number;
};

const LEVEL_PROGRESS: {
  level: number;
  cards: CardProgress[];
} = {
  level: 3,
  cards: [
    {
      kind: "radical",
      name: "spoon",
      value: "匕",
      meaning: "spoon",
      progress: 1,
    },
    { kind: "radical", name: "", value: "夂", meaning: "", progress: 2 },
    { kind: "radical", name: "", value: "干", meaning: "", progress: 3 },
    { kind: "radical", name: "", value: "广", meaning: "", progress: 5 },
    { kind: "radical", name: "", value: "扌", meaning: "", progress: 2 },
    { kind: "radical", name: "", value: "元", meaning: "", progress: 2 },
    { kind: "radical", name: "", value: "心", meaning: "", progress: 5 },
    { kind: "radical", name: "", value: "方", meaning: "", progress: 5 },
    { kind: "radical", name: "", value: "毛", meaning: "", progress: 2 },
    { kind: "radical", name: "", value: "父", meaning: "", progress: 2 },
    { kind: "radical", name: "", value: "古", meaning: "", progress: 2 },
    { kind: "radical", name: "", value: "用", meaning: "", progress: 2 },
    { kind: "radical", name: "", value: "矢", meaning: "", progress: 2 },
    { kind: "radical", name: "", value: "戸", meaning: "", progress: 5 },
    { kind: "radical", name: "", value: "幺", meaning: "", progress: 5 },
    { kind: "radical", name: "", value: "巾", meaning: "", progress: 2 },
    { kind: "radical", name: "", value: "生", meaning: "", progress: 2 },
    { kind: "radical", name: "", value: "今", meaning: "", progress: 2 },
    //
    { kind: "kanji", name: "", value: "万", meaning: "", progress: 1 },
    { kind: "kanji", name: "", value: "元", meaning: "", progress: 1 },
    { kind: "kanji", name: "", value: "内", meaning: "", progress: 1 },
    { kind: "kanji", name: "", value: "分", meaning: "", progress: 5 },
    { kind: "kanji", name: "", value: "切", meaning: "", progress: 5 },
    { kind: "kanji", name: "", value: "今", meaning: "", progress: 5 },
    { kind: "kanji", name: "", value: "午", meaning: "", progress: 5 },
    { kind: "kanji", name: "", value: "友", meaning: "", progress: 1 },
    { kind: "kanji", name: "", value: "太", meaning: "", progress: 1 },
    { kind: "kanji", name: "", value: "少", meaning: "", progress: 4 },
    { kind: "kanji", name: "", value: "引", meaning: "", progress: 1 },
    { kind: "kanji", name: "", value: "心", meaning: "", progress: 2 },
    { kind: "kanji", name: "", value: "戸", meaning: "", progress: 2 },
    { kind: "kanji", name: "", value: "方", meaning: "", progress: 3 },
    { kind: "kanji", name: "", value: "牛", meaning: "", progress: 3 },
    { kind: "kanji", name: "", value: "父", meaning: "", progress: 3 },
    { kind: "kanji", name: "", value: "毛", meaning: "", progress: 3 },
    { kind: "kanji", name: "", value: "止", meaning: "", progress: 3 },
    { kind: "kanji", name: "", value: "冬", meaning: "", progress: 1 },
    { kind: "kanji", name: "", value: "北", meaning: "", progress: 1 },
    { kind: "kanji", name: "", value: "半", meaning: "", progress: 5 },
    { kind: "kanji", name: "", value: "古", meaning: "", progress: 5 },
    { kind: "kanji", name: "", value: "台", meaning: "", progress: 5 },
    { kind: "kanji", name: "", value: "外", meaning: "", progress: 4 },
    { kind: "kanji", name: "", value: "市", meaning: "", progress: 2 },
    { kind: "kanji", name: "", value: "広", meaning: "", progress: 2 },
    { kind: "kanji", name: "", value: "母", meaning: "", progress: 4 },
    { kind: "kanji", name: "", value: "用", meaning: "", progress: 2 },
    { kind: "kanji", name: "", value: "矢", meaning: "", progress: 2 },
    { kind: "kanji", name: "", value: "生", meaning: "", progress: 2 },
  ],
};

function Dashboard() {
  const { data, isLoading, isError, mutate } = useSRS();

  if (isLoading) {
    return (
      <Box h="50vh" minH={"400px"}>
        <VStack spacing={4}>
          <Text>Loading data</Text>
          <Spinner />
        </VStack>
      </Box>
    );
  }

  if (isError) {
    return (
      <Box h="50vh" minH={"400px"}>
        <VStack>
          <Heading>Ooops..</Heading>
          <Text>{JSON.stringify(isError)}</Text>
        </VStack>
      </Box>
    );
  }

  if (!data) {
    return <>*TODO*</>;
  }

  if (data.length == 0) {
    fetch(`${API_URL}/api/v1/deck/${WANIKANI_DECK_ID}/subscribe`, {
      method: "PATCH",
      credentials: "include",
    })
      .then((res) => {
        if (res.ok) mutate();
      })
      .catch((e) =>
        console.error(`could not subscribe user to wanikani deck: ${e}`)
      );
  }

  const now = new Date();

  const lessons = data.filter(
    (c) =>
      c.available_at &&
      differenceInHours(new Date(c.available_at), now) <= 0 &&
      !c.started_at
  );
  const reviews = data.filter(
    (c) =>
      c.available_at &&
      differenceInHours(new Date(c.available_at), now) <= 0 &&
      c.started_at
  );

  return (
    <Container maxW={"8xl"} py={2} mt={4}>
      <Grid
        rowGap={4}
        py={0}
        columnGap={3}
        gridTemplateRows={{ md: "auto  1fr" }}
        gridTemplateColumns={{ base: "1fr", md: "3fr 1.2fr" }}
        gridTemplateAreas={{
          base: `
            "cards"
            "extra"
            "agenda"
            "progress"
          `,
          md: `
          "cards    agenda"
          "extra    agenda"
          "progress agenda"
        `,
        }}
      >
        <SimpleGrid columns={2} gridArea={"cards"} gap={3}>
          {/* lessons */}
          <StudyModal
            body={<></>}
            title={"Lessons Overview"}
            buttonText={`${lessons.length} lessons`}
            props={{
              bgColor: "orange.300",
              color: "orange.800",
            }}
            startURL="/lesson"
          />
          {/* reviews */}
          <StudyModal
            body={<></>}
            title={"Reviews Overview"}
            buttonText={`${reviews.length} reviews`}
            props={{
              bgColor: "purple.300",
              color: "purple.800",
            }}
            startURL="/review"
          />
        </SimpleGrid>
        {/* extra study */}
        <Box gridArea={"extra"}>
          <ExtraStudy />
        </Box>
        <Box
          h="full"
          w="full"
          placeSelf={"center"}
          gridArea={"agenda"}
          // maxW="550px"
          px={0}
        >
          {/* <Calendar /> */}
          <ReviewForecast />
        </Box>
        {/* level progress */}
        <Box gridArea={"progress"}>
          <LevelProgress completeCardProgress={5} />
        </Box>
      </Grid>
    </Container>
  );
}

export default function () {
  return (
    <AuthRoute>
      <Dashboard />
    </AuthRoute>
  );
}
