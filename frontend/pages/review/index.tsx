import { API_URL } from "@/lib/api/fetchApi";
import AuthRoute from "@/lib/auth/wrappers/AuthRoute";
import { useSRSSession } from "@/lib/hooks/srs";
import { CreateReviewRequest } from "@/lib/models/review";
import { QueueItemAnswers, SessionQueue } from "@/lib/models/srs";
import { MinimalSubject } from "@/lib/models/subject";
import FlashCard, { AnswerType, FlashCardProps } from "@/ui/Flashcard";
import { PageWithLayout } from "@/ui/layouts";
import {
  Box,
  Container,
  Flex,
  IconButton,
  Spinner,
  Text,
} from "@chakra-ui/react";
import { HomeIcon } from "@radix-ui/react-icons";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";

interface OrganizedSessionQueue {
  done: OrganizedSessionQueueItem[];
  queue: OrganizedSessionQueueItem[];
}
interface OrganizedSessionQueueItem {
  card_id: string;
  subject: MinimalSubject;
  answers: QueueItemAnswers;
  next?: OrganizedSessionQueueItem;
  fails: number;
}

function Review() {
  const sessionType = "review";
  const { session, isLoading, isError } = useSRSSession("review");

  const [queue, setQueue] = useState<OrganizedSessionQueue["queue"]>(),
    [done, setDone] = useState<OrganizedSessionQueue["done"]>([]);

  const onAttempt: FlashCardProps["onAttempt"] = async (attempt) => {
    if (!queue || queue.length == 0) {
      return null;
    }

    const card = queue[0];
    const res = checkAnswer(
      attempt,
      card.answers.expected,
      card.answers.blacklisted
    );

    let failed = res == "invalid";
    if (res == "blacklisted") {
      // TODO: show hint
      alert("BL");

      return { res, message: "TODO" };
    } else if (res == "invalid") {
      // TODO
      console.log("Invalid");
      card.fails += 1;
      setQueue(nextQueueState(queue, card));
    } else {
      // TODO: play sound
      // TODO: show item info
      console.log("Valid");

      if (!card.next) {
        const errors = {
          [card.answers.study_item_type]: card.fails,
        };
        for (const related of done) {
          const itemType = related.answers.study_item_type;
          errors[itemType] = related.fails;
        }

        const reviewRequest: CreateReviewRequest = {
          card: card.card_id,
          errors,
          session_type: sessionType,
        };

        try {
          const response = await fetch(`${API_URL}/api/v1/review`, {
            method: "POST",
            credentials: "include",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(reviewRequest),
          });

          if (!response.ok) {
            // TODO: notify user
            console.error(
              "create review returned an error:",
              await response.text()
            );
            return null;
          }
        } catch (e) {
          console.error(e);
          return null;
        }
      }

      setDone([...done, card]);
      // TODO: update queue only after user review
      setQueue(nextQueueState(queue, card.next));
    }

    return { res };
  };

  useEffect(() => {
    if (!session) {
      return;
    }

    setQueue(createOrganizedQueue(session));
  }, [session]);

  if (isLoading) {
    return (
      <>
        <Spinner />
      </>
    );
  }

  if (isError || !session) {
    console.error(isError);
    return <></>;
  }

  // TODO: check if it is already in the state, then retrieve it from there, case positive. otherwise, generate it as below
  // TODO: save to state and set expiration time to 10 minutes

  if (!queue || queue.length == 0) {
    return <></>;
  }

  return (
    <Container
      background="white"
      p={0}
      maxW="full"
      h="100vh"
      minH="500px"
      position="relative"
    >
      <Flex
        justifyContent="space-between"
        alignItems={"center"}
        position={"absolute"}
        top={"0"}
        left={"0"}
        zIndex={10}
        px={1}
        py={1}
        rounded="2xl"
      >
        <Link href={"/"}>
          <IconButton
            variant="ghost"
            rounded={"full"}
            aria-label="Home"
            color="white"
            _hover={{ backgroundColor: "whiteAlpha.300" }}
            icon={<HomeIcon transform={"scale(1.2)"} />}
          />
        </Link>
      </Flex>
      <FlashCard
        card={queue[0].card_id}
        subject={queue[0].subject}
        studyDataType={queue[0].answers.study_item_type}
        onAttempt={onAttempt}
      />
      {JSON.stringify(queue)}
    </Container>
  );
}

const ReviewWrapper: PageWithLayout = () => {
  return (
    <AuthRoute>
      <Review />
    </AuthRoute>
  );
};

ReviewWrapper.getLayout = (page) => page;

export default ReviewWrapper;

function nextQueueState(
  queue: OrganizedSessionQueueItem[],
  include?: OrganizedSessionQueueItem
) {
  let nextQueueState = shifted(queue);
  if (include) {
    nextQueueState = blindInsert(nextQueueState, include);
  }
  return nextQueueState;
}

function createOrganizedQueue(queue: SessionQueue) {
  const sq: OrganizedSessionQueue = { done: [], queue: [] };
  for (const card of queue.cards) {
    let prev: OrganizedSessionQueueItem | undefined = undefined;
    card.answers.reverse().forEach((answers) => {
      const item: OrganizedSessionQueueItem = {
        card_id: card.card_id,
        subject: card.subject,
        answers: answers,
        next: prev,
        fails: 0,
      };
      sq.queue.push(item);
      prev = item;
    });
  }
  return shuffleArray(sq.queue);
}

function checkAnswer(
  answer: string,
  valid: string[],
  blacklisted: string[]
): AnswerType {
  const ans = answer.toLowerCase();

  const validComp = valid.map<[AnswerType, string]>((v) => [
      "valid",
      v.toLowerCase(),
    ]),
    blacklistedComp = blacklisted.map<[AnswerType, string]>((b) => [
      "blacklisted",
      b.toLowerCase(),
    ]);

  for (const comp of validComp.concat(blacklistedComp)) {
    if (ans == comp[1]) {
      return comp[0];
    }
  }

  return "invalid";
}

function shifted<T>(array: T[]): T[] {
  if (array.length == 0) {
    return [];
  }
  const [_x, ...xs] = array;
  return xs;
}

function shuffleArray<T>(array: T[]): T[] {
  const copy = [...array];
  for (let i = copy.length - 1; i > 0; i--) {
    const j = randInt(0, i + 1);
    [copy[i], copy[j]] = [copy[j], copy[i]];
  }

  return copy;
}

function blindInsert<T>(array: T[], item: T): T[] {
  const index = randInt(Math.min(array.length, 1), array.length + 1);
  const copy = [...array.slice(0, index), item, ...array.slice(index)];
  return copy;
}

function randInt(min: number, max: number): number {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min) + min);
}
