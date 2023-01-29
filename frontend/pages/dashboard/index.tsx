import { Container, SimpleGrid, Stack, Box, Grid } from "@chakra-ui/react";
import { subHours, setHours, addDays } from "date-fns";

import AuthRoute from "@/lib/auth/wrappers/AuthRoute";
import { Calendar } from "@/ui/Calendar";
import ReviewForecast from "@/ui/Dashboard/ReviewForecast";
import LinkCard from "@/ui/Dashboard/LinkCard";

import LessonImage from "@/assets/images/Lesson.svg";
import ReviewImage from "@/assets/images/Review.svg";
import ExtraStudy from "@/ui/Dashboard/ExtraStudy";
import { LevelProgress } from "@/ui/Dashboard/LevelProgress";

const N_LESSONS = 13;
const N_REVIEWS = 30;

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

const USER_PROGRESS: {
  apprendice: { radical: number; kanji: number; vocabulary: number };
  guru: { radical: number; kanji: number; vocabulary: number };
  master: { radical: number; kanji: number; vocabulary: number };
  enlightened: { radical: number; kanji: number; vocabulary: number };
  burned: { radical: number; kanji: number; vocabulary: number };
} = {} as any;

// TODO: SORT
const REVIEW_SCHEDULE: {
  schedule: Date;
  reviews: {
    kind: "kanji" | "radical" | "vocabulary";
    name: string;
    level: number;
    value: string;
    meaning: string;
  }[];
}[] = [
  // now
  { schedule: subHours(new Date(), 2), reviews: [] },
  { schedule: new Date(), reviews: [] },
  // today
  { schedule: setHours(new Date(), 21), reviews: [] },
  { schedule: setHours(new Date(), 22), reviews: [] },
  { schedule: setHours(new Date(), 23), reviews: [] },
  // tomorrow
  { schedule: addDays(new Date().setHours(21), 1), reviews: [] },
  { schedule: addDays(new Date().setHours(22), 1), reviews: [] },
  { schedule: addDays(new Date().setHours(23), 1), reviews: [] },
  // the day after tomorrow
  { schedule: addDays(new Date().setHours(0), 2), reviews: [] },
  // after 5 days
  { schedule: addDays(new Date().setHours(10), 5), reviews: [] },
];

function Dashboard() {
  return (
    <Container maxW={"8xl"} mt={6}>
      <Grid
        rowGap={4}
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
          <LinkCard
            data={`${N_LESSONS} lessons`}
            image={LessonImage}
            props={{
              bgColor: "hsl(14.52, 100%, 80.59%)",
              color: "hsl(14.52, 100%, 20.59%)",
            }}
          />
          {/* reviews */}
          <LinkCard
            data={`${N_REVIEWS} reviews`}
            image={ReviewImage}
            props={{
              bgColor: "blue.200",
              color: "blue.800",
            }}
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
          <ReviewForecast schedules={REVIEW_SCHEDULE} />
        </Box>
        {/* level progress */}
        <Box gridArea={"progress"}>
          <LevelProgress
            cards={LEVEL_PROGRESS.cards}
            userLevel={LEVEL_PROGRESS.level}
            completeCardProgress={5}
          />
        </Box>
      </Grid>
    </Container>
  );
}

export default function () {
  return (
    // <AuthRoute>
    <Dashboard />
    // </AuthRoute>
  );
}
