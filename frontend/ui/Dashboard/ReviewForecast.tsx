import {
  Accordion,
  AccordionButton,
  AccordionIcon,
  AccordionItem,
  AccordionPanel,
  Box,
  Card,
  CardBody,
  CardHeader,
  Center,
  HStack,
  Image,
  Progress,
  StackDivider,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Tr,
  VStack,
} from "@chakra-ui/react";
import * as datefns from "date-fns";

import EmptyIllustration from "@/assets/illustration/girl-cheering.png";
import { useSRS } from "@/lib/hooks/srs";
import { SRSUserData } from "@/lib/models/srs";

type WithRequired<T, K extends keyof T> = T & { [P in K]-?: T[P] };
type CardAvailable = WithRequired<SRSUserData["cards"][number], "available_at">;

export default function ReviewForecast() {
  const { data } = useSRS();

  if (!data) {
    return <></>;
  }

  const { now, schedule } = makeForecast(data),
    weekForecast = forecastWeekHourly(schedule),
    availableNow = now.length;

  let totalReviews = availableNow,
    currentReviews = totalReviews;

  return (
    <Card
      shadow="none"
      rounded="2xl"
      // bg="black"
      // bg="#EBEBEB"
      // bg="white"
      // color={"gray.200"}
      w="full"
      h="full"
    >
      <CardHeader px={2}>
        <Text textTransform="capitalize" fontSize="lg">
          This week's schedule
        </Text>
      </CardHeader>
      <CardBody pt={0} px={2}>
        {Object.keys(weekForecast).length > 0 ? (
          <Accordion w="full" defaultIndex={[0]} allowMultiple>
            {Object.entries(weekForecast).map(([day, reviews], i) => {
              const reviewDate = new Date(day),
                reviewDayString = datefns.format(reviewDate, "dd"),
                reviewWeekDayString =
                  datefns.differenceInDays(
                    reviewDate,
                    new Date().setHours(0, 0, 0, 0)
                  ) <= 0
                    ? "Today"
                    : datefns.format(reviewDate, "EEEE");
              const fcReviews = Object.values(reviews).reduce(
                (curr, cs) => curr + cs.length,
                0
              );

              currentReviews = totalReviews;
              totalReviews += fcReviews;

              return (
                <AccordionItem
                  key={i}
                  _first={{ borderTop: "none" }}
                  _last={{ borderBottom: "none" }}
                >
                  {({ isExpanded }) => {
                    let totalDayReviews = currentReviews;

                    return (
                      <>
                        <AccordionButton px={2} pl={1}>
                          <AccordionIcon mr="1" />
                          <HStack justifyContent={"space-between"} w="full">
                            <Text as="span" fontSize={"lg"}>
                              {`${reviewDayString} - ${reviewWeekDayString}`}
                            </Text>
                            <HStack
                              fontSize={"md"}
                              divider={<StackDivider />}
                              hidden={isExpanded}
                              color={"gray.600"}
                            >
                              <Box as="span">+{fcReviews}</Box>
                              <Box as="span">{totalReviews}</Box>
                            </HStack>
                          </HStack>
                        </AccordionButton>
                        <AccordionPanel px={2}>
                          <TableContainer>
                            <Table
                              size="sm"
                              minW={"250px"}
                              sx={{
                                td: {
                                  borderBottom: "none",
                                },
                              }}
                            >
                              <Tbody>
                                {Object.entries(reviews).map(
                                  ([hour, fcCards], j) => {
                                    totalDayReviews += fcCards.length;

                                    const totRadical = fcCards.filter(
                                        (c) => c.subject.kind == "radical"
                                      ),
                                      totKanji = fcCards.filter(
                                        (c) => c.subject.kind == "kanji"
                                      ),
                                      totVocab = fcCards.filter(
                                        (c) => c.subject.kind == "vocabulary"
                                      );
                                    const pctRadical =
                                        (totRadical.length / fcCards.length) *
                                        100,
                                      pctKanji =
                                        (totKanji.length / fcCards.length) *
                                          100 +
                                        pctRadical,
                                      pctVocab =
                                        (totVocab.length / fcCards.length) *
                                          100 +
                                        pctKanji;

                                    return (
                                      <Tr key={j}>
                                        <Td pl={0} pr={2}>
                                          <Box as="span" fontWeight={"medium"}>
                                            {time24to12(parseInt(hour))}
                                          </Box>
                                        </Td>
                                        <Td w="full" pl={0} pr={5}>
                                          <Progress
                                            opacity={"0.8"}
                                            colorScheme={"orange"}
                                            value={
                                              (fcCards.length /
                                                (availableNow +
                                                  Object.values(weekForecast)
                                                    .length)) *
                                              100
                                            }
                                            sx={{
                                              div: {
                                                bgGradient: `linear(to-r,
                                                  blue.500 0% ${pctRadical}%,
                                                  pink.500 ${pctRadical}% ${pctKanji}%,
                                                  purple.500 ${pctKanji}% ${pctVocab}%
                                                )`,
                                                bgRepeat: "no-repeat",
                                              },
                                            }}
                                            // min={10}
                                            size="md"
                                            // roundedLeft="sm"
                                            borderRadius={"md"}
                                          />
                                        </Td>
                                        <Td
                                          w="max-content"
                                          px={2}
                                          textAlign="right"
                                          borderRightWidth={1}
                                        >
                                          <Box as="span" color={"gray.600"}>
                                            +{fcCards.length}
                                          </Box>
                                        </Td>
                                        <Td
                                          w="max-content"
                                          textAlign="left"
                                          pl={2}
                                          pr={0}
                                        >
                                          <Box as="span" color={"gray.600"}>
                                            {totalDayReviews}
                                          </Box>
                                        </Td>
                                      </Tr>
                                    );
                                  }
                                )}
                              </Tbody>
                            </Table>
                          </TableContainer>
                        </AccordionPanel>
                      </>
                    );
                  }}
                </AccordionItem>
              );
            })}
          </Accordion>
        ) : (
          <Center h="full">
            <VStack>
              <Image
                src={EmptyIllustration.src}
                maxW={"xs"}
                rounded="xl"
                filter="grayscale(1) brightness(1.0456)"
                px={8}
              />
              <Text px={3} textAlign="center">
                You have no reviews scheduled for this week
                {availableNow > 0
                  ? `, but you have ${availableNow} available to do right now!`
                  : "."}
              </Text>
            </VStack>
          </Center>
        )}
      </CardBody>
    </Card>
  );
}

function makeForecast(reviews: SRSUserData["cards"]) {
  const forecast: {
    now: CardAvailable[];
    schedule: CardAvailable[];
  } = {
    now: [],
    schedule: [],
  };

  const now = new Date();

  for (const review of reviews) {
    if (!review.available_at || !review.started_at) {
      continue;
    }

    if (datefns.differenceInHours(new Date(review.available_at), now) <= 0) {
      forecast.now.push(review as CardAvailable);
    } else {
      forecast.schedule.push(review as CardAvailable);
    }
  }

  return forecast;
}

function forecastWeekHourly(schedules: CardAvailable[]) {
  const now = new Date();
  const weekForecast: { [key: string]: { [key: number]: CardAvailable[] } } =
    {};

  for (const schedule of schedules) {
    let scheduleDate = new Date(schedule.available_at);
    scheduleDate.setMinutes(0, 0, 0);

    const scheduleDateString = new Date(
      new Date(scheduleDate).setHours(0, 0, 0, 0)
    ).toString();

    if (datefns.differenceInDays(scheduleDate, now) >= 7) {
      continue;
    }

    let schedules: { [key: number]: CardAvailable[] } = {};
    if (scheduleDateString in weekForecast) {
      schedules = weekForecast[scheduleDateString];
      if (scheduleDate.getHours() in schedules) {
        schedules[scheduleDate.getHours()].push(schedule);
      } else {
        schedules[scheduleDate.getHours()] = [schedule];
      }
    }

    weekForecast[scheduleDateString] = schedules;
  }

  console.log(weekForecast);
  return weekForecast;
}

function time24to12(n: number) {
  return datefns.format(new Date(n).setHours(n, 0, 0, 0), "hh aaa");
}
