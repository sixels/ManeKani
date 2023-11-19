import {
  Button,
  Drawer,
  DrawerBody,
  DrawerCloseButton,
  DrawerContent,
  DrawerFooter,
  DrawerOverlay,
  DrawerProps,
  Flex,
  FlexProps,
  useDisclosure,
} from "@chakra-ui/react";
import React, { PropsWithChildren, ReactElement, useRef } from "react";

type MyDrawerProps = PropsWithChildren<
  {
    drawerIcon: ReactElement;
    footer?: ReactElement;
  } & Pick<DrawerProps, "placement"> &
    Pick<FlexProps, "width">
>;

export default function MyDrawer({
  placement = "right",
  drawerIcon,
  footer,
  children,
}: MyDrawerProps) {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const btnRef = useRef<any>();

  return (
    <>
      <Button ref={btnRef} onClick={onOpen}>
        {drawerIcon}
      </Button>
      <Flex>
        <Drawer
          isOpen={isOpen}
          placement={placement}
          onClose={onClose}
          finalFocusRef={btnRef}
          closeOnEsc={true}
          size={{ base: "xs", md: "md" }}
        >
          <DrawerOverlay />
          <DrawerContent>
            <DrawerCloseButton
              size={"lg"}
              my={2}
              mx={4}
              bgColor="white"
              zIndex={100}
            />
            <DrawerBody px={0} pt={0}>
              {children}
            </DrawerBody>
            {footer && <DrawerFooter>{footer}</DrawerFooter>}
          </DrawerContent>
        </Drawer>
      </Flex>
    </>
  );
}
