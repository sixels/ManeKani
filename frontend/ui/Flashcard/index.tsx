import { MinimalSubject } from "@/lib/models/subject";
import {
  Center,
  Flex,
  FormControl,
  Heading,
  IconButton,
  Input,
  InputGroup,
  InputRightElement,
  Mark,
  VStack,
} from "@chakra-ui/react";
import { ArrowRightIcon } from "@radix-ui/react-icons";
import {
  FormEvent,
  PropsWithChildren,
  useCallback,
  useEffect,
  useRef,
  useState,
} from "react";
import * as wanakana from "wanakana";

export type FlashCardProps = {
  card: string;
  subject: MinimalSubject;
  studyDataType: string;
  onAttempt(
    attempt: string
  ): Promise<
    | ({ res: "valid" | "invalid" } | { res: "blacklisted"; message: string })
    | null
  >;
};

export type AnswerType = "valid" | "invalid" | "blacklisted";

export default function FlashCard({
  card,
  subject,
  studyDataType,
  onAttempt,
}: FlashCardProps) {
  const valueText = subject.value;
  const inputElem = useRef<any>();

  const [isInvalid, setIsInvalid] = useState(false);
  // const valueText = "A";

  let valueTextSize = `0px`;
  if (valueText) {
    valueTextSize = `clamp(1.8em, 40vw / (${valueText.length}), 6.3em)`;
  }

  useEffect(() => {
    if (studyDataType == "reading" && inputElem && inputElem.current) {
      wanakana.bind(inputElem.current);
    }
  }, [inputElem]);

  const onSubmit = useCallback(
    async (e: FormEvent) => {
      e.preventDefault();

      const answer = (inputElem.current.value as string).trim();
      if (answer.length == 0) {
        setIsInvalid(true);
        return;
      }

      setIsInvalid(false);

      const attemptRes = await onAttempt(answer);
      if (!attemptRes) {
        // TODO: set unexpected message
        return;
      }

      if (attemptRes.res == "valid") {
        inputElem.current.value = "";
      } else {
        setIsInvalid(true);
        // TODO: set error message
      }
    },
    [onAttempt]
  );

  return (
    <Flex
      flexDir={"column"}
      maxW="full"
      h="full"
      px={2}
      py={2}
      bgGradient={"linear(to-b, blue.300, blue.300, blue.400, blue.500)"}
      overflow={"hidden"}
    >
      <Center
        color="white"
        flex={1}
        w="full"
        roundedTop="2xl"
        position={"relative"}
      >
        <Heading
          as="h1"
          fontWeight={"normal"}
          fontSize={valueTextSize}
          wordBreak="break-all"
          fontFamily={`"Noto Sans JP", sans-serif`}
        >
          {valueText}
        </Heading>
      </Center>
      <ContentCard>
        <Heading
          as="h2"
          fontSize="2xl"
          textTransform={"capitalize"}
          fontWeight={"normal"}
          py={3}
          pt={5}
          px={4}
        >
          {subject.kind} <Mark fontWeight={"medium"}>{studyDataType}</Mark>
        </Heading>

        <FormControl onSubmit={onSubmit} as="form">
          <InputGroup>
            <Input
              variant={"filled"}
              shadow="none"
              placeholder="Your Answer"
              fontSize={"xl"}
              colorScheme="blackAlpha"
              bg="#E4EAEA"
              size="lg"
              h={"16"}
              textAlign={"center"}
              maxLength={30}
              isInvalid={isInvalid}
              ref={inputElem}
              _focus={{ borderColor: isInvalid && "red.500" }}
              pr="14"
            />
            <InputRightElement my="3" mr="2">
              <IconButton
                aria-label="next"
                colorScheme={"blackAlpha"}
                variant={"ghost"}
                type="submit"
                icon={<ArrowRightIcon />}
              />
            </InputRightElement>
          </InputGroup>
        </FormControl>
      </ContentCard>
    </Flex>
  );
}

function ContentCard({ children }: PropsWithChildren<{}>) {
  return (
    <VStack bg="whiteAlpha.900" gap={2} px={2} pb="3" spacing={0} rounded="xl">
      {children}
    </VStack>
  );
}
