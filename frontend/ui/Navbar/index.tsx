import {
  Accordion,
  AccordionButton,
  AccordionIcon,
  AccordionItem,
  AccordionPanel,
  Box,
  Button,
  Divider,
  Flex,
  HStack,
  IconButton,
  Img,
  Popover,
  PopoverArrow,
  PopoverContent,
  PopoverTrigger,
  Text,
  useDisclosure,
} from "@chakra-ui/react";
import React from "react";
import { HamburgerMenuIcon, MagnifyingGlassIcon } from "@radix-ui/react-icons";
import Link from "next/link";

import MyDrawer from "./components/Drawer";

import Logo from "@/assets/logo-aside.svg";
import LogoSmall from "@/assets/icon.svg";

import { menuItems } from "./items";
import { UserAvatar } from "./components/UserAvatar";
import { SignInButtons } from "./components/SignIn";

export default function Navbar() {
  return (
    <Flex
      alignItems={"center"}
      justifyContent={"space-between"}
      px={{ base: 4, md: 6, lg: 10 }}
      py={1}
      position="fixed"
      top={0}
      left={0}
      borderBottom={"1px"}
      borderColor="gray.200"
      w={"full"}
      h={"60px"}
      zIndex={100}
      backgroundColor={"white"}
    >
      <Link href="/">
        <Box display={{ base: "none", md: "block" }}>
          <Img src={Logo.src} alt="ManeKani Logo" width={"auto"} height={35} />
        </Box>
        <Box display={{ base: "block", md: "none" }}>
          <Img
            src={LogoSmall.src}
            alt="ManeKani Logo"
            width={"auto"}
            height={35}
          />
        </Box>
      </Link>

      <HStack spacing={2} height="40px">
        <Box display={{ base: "none", md: "block" }}>
          <NavLinks />
        </Box>
        <IconButton
          aria-label="Search"
          rounded={"full"}
          variant={"ghost"}
          icon={<MagnifyingGlassIcon />}
        />
        <Divider
          orientation="vertical"
          borderColor={"gray.300"}
          // border={"1px"}
          // height="10px"
        />
        <Flex alignItems={"center"}>
          <UserAvatar fallback={<SignInButtons />} />
        </Flex>
        <Box display={{ base: "block", md: "none" }}>
          <MobileMenu />
        </Box>
      </HStack>
    </Flex>
  );
}

function MobileMenu() {
  const colors: { [key: string]: [string, string] } = {
    radical: ["blue.500", "blue.50"],
    kanji: ["pink.500", "pink.50"],
    vocabulary: ["purple.500", "purple.50"],
  };

  return (
    <MyDrawer placement="right" drawerIcon={<HamburgerMenuIcon />}>
      <Box w="full" h="full" bgColor="whitesmoke">
        <Accordion allowMultiple>
          {menuItems.map((item, i) => {
            const [bgColor, fgColor] =
              item.text in colors ? colors[item.text] : ["gray.700", "gray.50"];

            return (
              <AccordionItem key={i}>
                {({ isExpanded }) => (
                  <>
                    <AccordionButton fontSize="lg">
                      <Flex w={"full"} h="3em" align="center" px={5} gap={1}>
                        <AccordionIcon />
                        <Box as="span" textTransform={"capitalize"}>
                          {item.text}
                        </Box>
                      </Flex>
                    </AccordionButton>
                    <AccordionPanel pb={2} bg={bgColor} color={fgColor}>
                      {isExpanded ? item.menu : <></>}
                    </AccordionPanel>
                  </>
                )}
              </AccordionItem>
            );
          })}
        </Accordion>
      </Box>
    </MyDrawer>
  );
}

function NavLinks() {
  const colors: { [key: string]: [string, string, string, string] } = {
    radical: ["blue.100", "blue.800", "blue.500", "blue.50"],
    kanji: ["pink.100", "pink.800", "pink.500", "pink.50"],
    vocabulary: ["purple.100", "purple.800", "purple.500", "purple.50"],
    default: ["gray.100", "gray.800", "gray.700", "gray.50"],
  };

  const navLinkElements = menuItems
    .filter((item) => !item.mobileOnly)
    .map((item, i) => {
      const { onOpen, onClose, isOpen } = useDisclosure();

      const [bgColor, fgColor, menuBg, menuFg] =
        item.text in colors ? colors[item.text] : colors["default"];

      const element = (
        <Button
          variant="ghost"
          _hover={{
            backgroundColor: bgColor,
            color: fgColor,
          }}
          py={2}
          px={2}
          rounded="sm"
          fontWeight="normal"
          textTransform={"capitalize"}
          color="gray.700"
        >
          {item.text}
        </Button>
      );

      if (item.isLink) {
        return <Link href={item.href}>{element}</Link>;
      }
      return (
        <Popover
          key={i}
          isOpen={isOpen}
          onOpen={onOpen}
          onClose={onClose}
          placement="bottom"
          closeOnBlur={true}
        >
          <PopoverTrigger>{element}</PopoverTrigger>
          <PopoverContent
            mt="1"
            px={2}
            py={3}
            bgColor={menuBg}
            color={menuFg}
            border={"none"}
            boxShadow="2xl"
            w="full"
          >
            <PopoverArrow bgColor={menuBg} />
            <Box maxH="calc(100vh - 90px)" overflowY={"auto"}>
              {item.menu}
            </Box>
          </PopoverContent>
        </Popover>
      );
    });

  return <HStack spacing={2}>{navLinkElements}</HStack>;
}
