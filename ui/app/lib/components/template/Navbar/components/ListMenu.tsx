// import {
// 	Flex,
// 	LinkBox,
// 	LinkOverlay,
// 	Stack,
// 	StackDivider,
// 	Text,
// } from "@chakra-ui/react";
// import React from "react";

// import { LevelList } from "../items";

// export type ListMenuProps = {
// 	kind: "radical" | "kanji" | "vocabulary";
// 	levels: LevelList;
// };

// function ListMenu({ kind, levels }: ListMenuProps) {
// 	function prefixLevel(n: number): string {
// 		return n.toString().padStart(2, "0");
// 	}

// 	return (
// 		<Stack divider={<StackDivider borderColor={"whiteAlpha.400"} />}>
// 			{levels.map((level, i) => {
// 				const levelName = level.name.split(" ")[1];
// 				return (
// 					<LinkBox
// 						key={i}
// 						_hover={{ bgColor: "whiteAlpha.300" }}
// 						px={2}
// 						rounded="md"
// 					>
// 						<LinkOverlay href={`/${kind}?difficulty=${levelName}`}>
// 							<Flex
// 								align={"left"}
// 								py={2}
// 								gap={8}
// 								justifyContent="space-between"
// 							>
// 								<Text as="span" textTransform={"uppercase"} fontSize="sm">
// 									{level.name}
// 								</Text>
// 								<Text as="span" textTransform={"capitalize"} fontSize="md">
// 									Levels {prefixLevel(level.range[0])}-
// 									{prefixLevel(level.range[1])}
// 								</Text>
// 							</Flex>
// 						</LinkOverlay>
// 					</LinkBox>
// 				);
// 			})}
// 		</Stack>
// 	);
// }

// export default React.memo(ListMenu);
