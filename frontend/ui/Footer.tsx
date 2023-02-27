import { Box, Container, Link, Stack, Text } from "@chakra-ui/react";
import Image from "next/image";

import Logo from "@/assets/logo-aside-dark.svg";

export default function Footer() {
  return (
    <Box bgColor={"#121D2B"} color={"gray.300"} mt="28" w="full" h="270px">
      <Container
        as={Stack}
        maxW={"6xl"}
        py={4}
        pt={8}
        spacing={4}
        justify={"center"}
        align={"center"}
      >
        <Box py={6}>
          <Image
            src={Logo}
            alt={"ManeKani logo"}
            width={192}
            height={0}
            color="white"
          />
        </Box>
        <Stack direction={"row"} spacing={6}>
          <Link href={"/"}>Home</Link>
          <Link href={"/dashboard"}>Dashboard</Link>
          <Link href={"/about"}>About</Link>
          <Link href={"/contact"}>Contact</Link>
        </Stack>
      </Container>

      <Box borderTopWidth={1} borderStyle={"solid"} borderColor={"gray.700"}>
        <Container
          as={Stack}
          maxW={"5xl"}
          py={4}
          direction={{ base: "column", md: "row" }}
          spacing={4}
          justify={"space-between"}
          align={"center"}
        >
          <Text>© 2023 ManeKani. All rights reserved</Text>
          <Stack direction={"row"} spacing={6}></Stack>
        </Container>
      </Box>
    </Box>
  );
}
