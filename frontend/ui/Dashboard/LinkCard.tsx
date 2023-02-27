import {
  Box,
  Button,
  Card,
  CardBody,
  CardProps,
  Center,
  Flex,
  Heading,
  HStack,
  IconButton,
  Image,
  Img,
  Icon,
  LinkBox,
  LinkOverlay,
  Text,
  VStack,
} from "@chakra-ui/react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  ArrowRightIcon,
  ButtonIcon,
  ExternalLinkIcon,
} from "@radix-ui/react-icons";

export type LinkCardProps = {
  data: string;
  props?: CardProps;
};

export default function LinkCard({ data, props }: LinkCardProps) {
  return (
    <LinkBox>
      <Card
        rounded="2xl"
        shadow="sm"
        bgGradient="linear(to-br, whiteAlpha.600, whiteAlpha.600, whiteAlpha.300)"
        transition={"ease"}
        {...props}
      >
        <CardBody
          px={1.5}
          py={1.5}
          backgroundImage={`url("/assets/topograph.svg")`}
        >
          <VStack rounded="xl" spacing={14} py={4}>
            {/* <Box
              alignSelf={"flex-start"}
              bg="whiteAlpha.700"
              p={5}
              userSelect={"none"}
              rounded={"xl"}
              h="20px"
              w="20px"
              roundedTopLeft={"2xl"}
            >
              {image && (
                <Image
                  src={image.src}
                  h="20px"
                  w="20px"
                  scale={2.2}
                  transform="auto"
                />
              )}
            </Box> */}
            <HStack w="full" justifyContent="space-between" px={4}>
              <LinkOverlay href="#">
                <Heading size="sm">{data}</Heading>
              </LinkOverlay>
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
          </VStack>
        </CardBody>
      </Card>
    </LinkBox>
  );
}
