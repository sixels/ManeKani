// import {
// 	Grid,
// 	GridItem,
// 	LinkBox,
// 	LinkOverlay,
// 	Text,
// 	VStack,
// } from "@chakra-ui/react";
// import React from "react";

// import { LevelList } from "../items";

// export type LevelsMenuProps = {
// 	levels: LevelList;
// };
// function LevelsMenu({ levels }: LevelsMenuProps) {
// 	const levelButtons = levels.map((level) => {
// 		return (
// 			<VStack key={level.name} align={"left"} py={2}>
// 				<Text as="span" textTransform={"uppercase"} fontSize="sm">
// 					{level.name}
// 				</Text>
// 				<Grid
// 					templateColumns={{ base: "repeat(5, 3em)", md: "repeat(10, 2.5em)" }}
// 					columnGap={2}
// 					rowGap={1.5}
// 					flexWrap={"wrap"}
// 				>
// 					{[...Array(level.range[1] + 1 - level.range[0]).keys()].map((l) => {
// 						const levelNr = l + level.range[0];
// 						return (
// 							<GridItem key={l}>
// 								<LinkBox
// 									alignItems={"center"}
// 									justifyContent={"center"}
// 									display={"flex"}
// 									h={"2.5em"}
// 									w="100%"
// 									color="white"
// 									bg={"whiteAlpha.50"}
// 									rounded="md"
// 									_hover={{ bgColor: "whiteAlpha.200" }}
// 								>
// 									<LinkOverlay href={`/level/${levelNr}`}>
// 										{levelNr.toString().padStart(2, "0")}
// 									</LinkOverlay>
// 								</LinkBox>
// 							</GridItem>
// 						);
// 					})}
// 				</Grid>
// 			</VStack>
// 		);
// 	});

// 	return <>{levelButtons}</>;
// }

// export default React.memo(LevelsMenu);
