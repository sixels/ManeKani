// import { AuthContext, useSession } from "@/lib/auth/context";
// import { useUser } from "@/lib/hooks/user";
// import {
// 	Menu,
// 	MenuButton,
// 	HStack,
// 	Avatar,
// 	Box,
// 	MenuList,
// 	MenuGroup,
// 	MenuDivider,
// 	MenuItem,
// } from "@chakra-ui/react";
// import { ChevronDownIcon } from "@radix-ui/react-icons";
// import { useRouter } from "next/router";
// import { ReactElement, useContext } from "react";

// export type UserAvatarProps = {
// 	fallback: ReactElement;
// };
// export function UserAvatar({ fallback }: UserAvatarProps) {
// 	// const { user, isLoading, isError } = useUser();

// 	// if (isLoading) {
// 	//   return null;
// 	// }

// 	// if (isError || !user) {
// 	//   return <>{fallback}</>;
// 	// }

// 	const router = useRouter();

// 	const authSession = useSession();
// 	const username = authSession?.username;

// 	console.log(authSession);

// 	return (
// 		<Menu>
// 			<MenuButton py={2} transition="all 0.3s" _focus={{ boxShadow: "none" }}>
// 				<HStack spacing={1}>
// 					<Avatar
// 						// size="sm"
// 						h={"38px"}
// 						w={"38px"}
// 						// height="10%"
// 						src={`https://api.dicebear.com/5.x/lorelei/svg?seed=${username}&flip=true&backgroundColor=F64D07`}
// 					/>
// 					<Box display={{ base: "none", md: "block" }}>
// 						<ChevronDownIcon />
// 					</Box>
// 				</HStack>
// 			</MenuButton>
// 			<MenuList bg={"white"} borderColor={"gray.200"}>
// 				<MenuGroup
// 					title={`signed in as ${username}`}
// 					color={"gray.700"}
// 					textAlign={"center"}
// 				>
// 					<MenuDivider />
// 					<MenuItem>Profile</MenuItem>
// 					<MenuItem>Settings</MenuItem>
// 					<MenuItem>Billing</MenuItem>
// 					<MenuDivider />
// 					<MenuItem
// 						onClick={() => {
// 							let logoutURL = authSession?.logout_url;
// 							if (logoutURL) {
// 								console.debug(logoutURL);
// 								router.push(logoutURL);
// 							}
// 						}}
// 					>
// 						Sign out
// 					</MenuItem>
// 				</MenuGroup>
// 			</MenuList>
// 		</Menu>
// 	);
// }
