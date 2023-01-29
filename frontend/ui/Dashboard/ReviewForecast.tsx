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

import EmptyIllustration from "@/assets/open-doodles/svg/SittingDoodle.svg";

type ReviewForecastProps = {
  schedules: { schedule: Date; reviews: any[] }[];
};

export default function ReviewForecast({ schedules }: ReviewForecastProps) {
  const now = new Date();
  const [active, scheduled] = filterDivide(
    schedules,
    ({ schedule }) => datefns.differenceInHours(schedule, now) <= 0
  );

  const total = active.reduce((acc, cur) => acc + cur.reviews.length, 0);
  const forecast = forecastFromSchedules(scheduled, total);

  return (
    <Card shadow="sm" rounded="2xl" bg="white" w="full" h="full">
      <CardHeader>
        <Text textTransform="capitalize" fontSize="lg">
          This week's schedule
        </Text>
      </CardHeader>
      <CardBody px={2} pt={0}>
        {scheduled.length > 0 ? (
          <Accordion w="full" defaultIndex={[0]} allowMultiple>
            {Object.entries(forecast).map(([day, dayReviews], i) => {
              const reviewDate = new Date(day),
                reviewDay = datefns.format(reviewDate, "dd"),
                reviewWeekDay =
                  datefns.differenceInDays(
                    reviewDate,
                    new Date().setHours(0, 0, 0, 0)
                  ) <= 0
                    ? "Today"
                    : datefns.format(reviewDate, "EEEE");

              return (
                <AccordionItem
                  key={i}
                  _first={{ borderTop: "none" }}
                  _last={{ borderBottom: "none" }}
                >
                  {({ isExpanded }) => (
                    <>
                      <AccordionButton px={2} pl={1}>
                        <AccordionIcon mr="1" />
                        <HStack justifyContent={"space-between"} w="full">
                          <Text as="span" fontSize={"lg"}>
                            {`${reviewDay} - ${reviewWeekDay}`}
                          </Text>
                          <HStack
                            fontSize={"md"}
                            divider={<StackDivider />}
                            hidden={isExpanded}
                            color={"gray.600"}
                          >
                            <Box as="span">+{dayReviews.reviews}</Box>
                            <Box as="span">{dayReviews.total}</Box>
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
                              {dayReviews.schedules.map(
                                ({ schedule, reviews, total }, j) => (
                                  <Tr key={j}>
                                    <Td pl={0} pr={2}>
                                      <Box as="span" fontWeight={"medium"}>
                                        {time24to12(schedule)}
                                      </Box>
                                    </Td>
                                    <Td w="full" pl={0} pr={5}>
                                      <Progress
                                        w="100%"
                                        opacity={"0.8"}
                                        colorScheme={"orange"}
                                        value={
                                          (reviews.length /
                                            (dayReviews.max || 1)) *
                                          100
                                        }
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
                                        +{reviews.length}
                                      </Box>
                                    </Td>
                                    <Td
                                      w="max-content"
                                      textAlign="left"
                                      pl={2}
                                      pr={0}
                                    >
                                      <Box as="span" color={"gray.600"}>
                                        {total}
                                      </Box>
                                    </Td>
                                  </Tr>
                                )
                              )}
                            </Tbody>
                          </Table>
                        </TableContainer>
                      </AccordionPanel>
                    </>
                  )}
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
                filter="grayscale(1)"
                p={8}
              />
              <Text px={3} textAlign="center" pb={6}>
                You have no reviews scheduled for this week
                {active.length > 0
                  ? `, but you have ${active.length} available to do right now!`
                  : "."}
              </Text>
            </VStack>
          </Center>
        )}
      </CardBody>
    </Card>
  );
}

function forecastFromSchedules(
  schedules: ReviewForecastProps["schedules"],
  total: number = 0
) {
  const forecast_by_day: {
    [key: string]: {
      schedules: (ReviewForecastProps["schedules"][number] & {
        total: number;
      })[];
      reviews: number;
      total: number;
      max: number;
    };
  } = {};

  const today = new Date().setHours(0, 0, 0, 0);
  for (const s of schedules) {
    if (datefns.differenceInDays(today, s.schedule) >= 7) {
      continue;
    }

    const day = s.schedule.toDateString();
    if (!(day in forecast_by_day)) {
      forecast_by_day[day] = { schedules: [], reviews: 0, total: 0, max: 0 };
    }
    forecast_by_day[day].schedules.push({ ...s, total: 0 });
  }

  for (const s of Object.values(forecast_by_day)) {
    for (const f of s.schedules) {
      total += f.reviews.length;
      s.reviews += f.reviews.length;

      f.total = total;

      s.max = Math.max(f.reviews.length, s.max);
    }
    s.total = total;
  }

  return forecast_by_day;
}

function filterDivide<T>(array: T[], filter: (elm: T) => boolean): [T[], T[]] {
  return array.reduce(
    ([pass, fail], elm) =>
      filter(elm) ? [[...pass, elm], fail] : [pass, [...fail, elm]],
    [[] as T[], [] as T[]]
  );
}

function time24to12(n: Date) {
  return datefns.format(n, "hh aaa");
}
