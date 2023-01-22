import { useUserData } from "@/lib/auth/context";
import {
  Menu,
  MenuButton,
  HStack,
  Avatar,
  Box,
  MenuList,
  MenuGroup,
  MenuDivider,
  MenuItem,
} from "@chakra-ui/react";
import { ChevronDownIcon } from "@radix-ui/react-icons";
import { ReactElement } from "react";

export type UserAvatarProps = {
  fallback: ReactElement;
};
export function UserAvatar({ fallback }: UserAvatarProps) {
  const user = useUserData();

  if (user.loading) {
    return null;
  }

  if (!user.user) {
    return <>{fallback}</>;
  }

  return (
    <Menu>
      <MenuButton py={2} transition="all 0.3s" _focus={{ boxShadow: "none" }}>
        <HStack spacing={1}>
          <Avatar
            size="sm"
            src={`https://api.dicebear.com/5.x/lorelei/svg?seed=${user.user.username}&flip=true&backgroundColor=F64D07`}
          />
          <Box display={{ base: "none", md: "block" }}>
            <ChevronDownIcon />
          </Box>
        </HStack>
      </MenuButton>
      <MenuList bg={"white"} borderColor={"gray.200"}>
        <MenuGroup
          title={`signed in as ${user.user.username}`}
          color={"gray.700"}
          textAlign={"center"}
        >
          <MenuDivider />
          <MenuItem>Profile</MenuItem>
          <MenuItem>Settings</MenuItem>
          <MenuItem>Billing</MenuItem>
          <MenuDivider />
          <MenuItem>Sign out</MenuItem>
        </MenuGroup>
      </MenuList>
    </Menu>
  );
}
