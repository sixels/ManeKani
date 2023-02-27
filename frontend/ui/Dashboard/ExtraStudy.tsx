import {
  Button,
  ButtonGroup,
  Card,
  CardBody,
  CardHeader,
  Flex,
  Image,
  Text,
  VStack,
} from "@chakra-ui/react";

import Illustration from "@/assets/illustration/boy-studying.png";

export type ExtraStudyProps = {};
export default function ExtraStudy(_: ExtraStudyProps) {
  return (
    <Card shadow="xl" bg="#000" color={"gray.200"} rounded="2xl">
      <CardHeader>
        <Text textTransform="capitalize" fontSize="lg">
          Extra Study
        </Text>
      </CardHeader>
      <CardBody pt={0}>
        <Flex
          flexDir={{ base: "column", lg: "row-reverse" }}
          rowGap={6}
          columnGap={4}
          alignItems={"center"}
        >
          <Image
            src={Illustration.src}
            flex={{ lg: 1 }}
            maxW={"xs"}
            // p={8}
            filter="grayscale(1)"
            rounded="xl"
            bgColor="black"
          />
          <VStack align="flex-start" gridArea={"content"} flex={{ lg: 1 }}>
            <Text fontSize="md">
              Extra Study lets you practice your radicals, kanji, and vocabulary
              outside of your regular Reviews. Practice done here won't count
              toward regular Reviews, but will count toward making your memory
              stronger!
            </Text>
            <ButtonGroup w="100%" orientation="vertical">
              <Button
                fontWeight="normal"
                variant="outline"
                disabled={false}
                colorScheme={"orange"}
                justifyContent={"flex-start"}
                _hover={{ backgroundColor: "orange.800", color: "orange.400" }}
              >
                3 Recent Lessons
              </Button>
              <Button
                fontWeight="normal"
                variant="outline"
                // colorScheme={"orange"}
                disabled
                justifyContent={"flex-start"}
              >
                0 Recent Mistakes
              </Button>
              <Button
                fontWeight="normal"
                variant="outline"
                // colorScheme={"orange"}
                disabled
                justifyContent={"flex-start"}
              >
                0 Burned Items
              </Button>
            </ButtonGroup>
          </VStack>
        </Flex>
      </CardBody>
    </Card>
  );
}
