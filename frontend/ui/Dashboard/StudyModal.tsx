import {
  Box,
  Button,
  ButtonGroup,
  Card,
  CardBody,
  CardProps,
  Center,
  Heading,
  HStack,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  useDisclosure,
} from "@chakra-ui/react";
import { ArrowRightIcon } from "@radix-ui/react-icons";
import Link from "next/link";

export type StudyModalProps = {
  title: string;
  body: JSX.Element;
  buttonText: string | number;
  props?: CardProps;
  startURL: string;
};

export default function StudyModal({
  buttonText,
  title,
  body,
  props,
  startURL,
}: StudyModalProps) {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <>
      <Button onClick={onOpen} h="auto" variant={"unstyled"}>
        <Card
          rounded="2xl"
          shadow="md"
          bgGradient="linear(to-br, whiteAlpha.600, whiteAlpha.600, whiteAlpha.200)"
          transition={"ease"}
          {...props}
        >
          <CardBody
            px={1.5}
            py={1.5}
            backgroundImage={`url("/assets/topograph.svg")`}
            backgroundPosition="center"
          >
            <HStack w="full" justifyContent="space-between" px={4} py={4}>
              <Heading size="sm">{buttonText}</Heading>
              <Box
                rounded="lg"
                bg="whiteAlpha.700"
                _hover={{ bg: "whiteAlpha.800" }}
                _active={{ bg: "whiteAlpha.800" }}
                p={3}
              >
                <ArrowRightIcon />
              </Box>
            </HStack>
          </CardBody>
        </Card>
      </Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>{title}</ModalHeader>
          <ModalCloseButton />
          <ModalBody>{body}</ModalBody>

          <ModalFooter>
            <ButtonGroup spacing={3}>
              <Button colorScheme="blue" onClick={onClose}>
                Close
              </Button>
              <Link href={startURL}>
                <Button variant="ghost">Start</Button>
              </Link>
            </ButtonGroup>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
}
