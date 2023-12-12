import { Button } from "../../general/Button";

export default function Navbar({
	user,
}: {
	user: {
		username: string;
		email: string;
		displayName?: string;
		profilePicture?: string;
	};
}) {
	return (
		<div className="flex items-center justify-between px-4 md:px-6 py-1 top-0 left-0 w-full h-[66px] z-[20]">
			<div>
				
			</div>

			<div className="flex items-center justify-center gap-4 h-10">
				<Button
					className="!h-10 !w-10 !bg-transparent hover:!bg-neutral-200 !border-0 px-0 py-0 flex justify-center items-center"
					isSecondary
				>
					<span className="material-symbols-outlined">notifications</span>
				</Button>
				<UserAvatar user={user} />
				{/* <div className="hidden md:block"> */}
				{/* Nav links */}
				{/* <NavLinks /> */}
				{/* </div> */}
				{/* <IconButton
          aria-label="Search"
          rounded={'full'}
          variant={'ghost'}
          icon={<MagnifyingGlassIcon />}
        /> */}
				{/* <Divider
          orientation="vertical"
          borderColor={'gray.300'}
          // border={"1px"}
          // height="10px"
        /> */}
				{/* <Flex alignItems={'center'}>
          <UserAvatar fallback={<SignInButtons />} />
        </Flex> */}
				{/* <Box display={{ base: 'block', md: 'none' }}>
          <MobileMenu />
        </Box> */}
			</div>
		</div>
	);
}

function UserAvatar({ user }: Parameters<typeof Navbar>[0]) {
	return (
		<div className="flex items-center gap-3">
			{/* TODO: change this */}
			<Button
				isSecondary
				type="button"
				className="!px-1 !border-0 !bg-transparent flex items-center space-x-2"
			>
				<img
					// temporary profile picture image
					src="https://avatars.githubusercontent.com/u/68879242?v=4"
					alt="user avatar"
					className="w-10 h-10 rounded-full shadow"
				/>

				<div className="text-left hidden md:block pr-3">
					{user.displayName && (
						<span className="block text-neutral-700 text-xs font-medium">
							{user.displayName}
						</span>
					)}
					<span className="block text-neutral-900 text-xs font-medium">
						{user.username}
					</span>
				</div>

				<div>
					{/* chevron down svg */}
					<svg
						xmlns="http://www.w3.org/2000/svg"
						className="h-5 w-5 text-neutral-700"
						viewBox="0 0 20 20"
						fill="currentColor"
					>
						<title>Expand menu</title>
						<path
							fillRule="evenodd"
							// eslint-disable-next-line max-len
							d="M10 12.586L4.707 7.293a1 1 0 011.414-1.414L10 10.758l4.879-4.879a1 1 0 111.414 1.414L10 12.586z"
							clipRule="evenodd"
						/>
					</svg>
				</div>
			</Button>
		</div>
	);
}

function MobileMenu() {
	const colors: { [key: string]: [string, string] } = {
		radical: ["blue.500", "blue.50"],
		kanji: ["pink.500", "pink.50"],
		vocabulary: ["purple.500", "purple.50"],
	};

	return (
		<></>
		// <MyDrawer placement="right" drawerIcon={<HamburgerMenuIcon />}>
		//   <Box w="full" h="full" bgColor="whitesmoke">
		//     <Accordion allowMultiple>
		//       {menuItems.map((item, i) => {
		//         const [bgColor, fgColor] =
		//           item.text in colors ? colors[item.text] : ['gray.700', 'gray.50'];

		//         return (
		//           <AccordionItem key={i}>
		//             {({ isExpanded }) => (
		//               <>
		//                 <AccordionButton fontSize="lg">
		//                   <Flex w={'full'} h="3em" align="center" px={5} gap={1}>
		//                     <AccordionIcon />
		//                     <Box as="span" textTransform={'capitalize'}>
		//                       {item.text}
		//                     </Box>
		//                   </Flex>
		//                 </AccordionButton>
		//                 <AccordionPanel pb={2} bg={bgColor} color={fgColor}>
		//                   {isExpanded ? item.menu : <></>}
		//                 </AccordionPanel>
		//               </>
		//             )}
		//           </AccordionItem>
		//         );
		//       })}
		//     </Accordion>
		//   </Box>
		// </MyDrawer>
	);
}

// function NavLinks() {
//   const colors: { [key: string]: [string, string, string, string] } = {
//     radical: ['blue.100', 'blue.800', 'blue.500', 'blue.50'],
//     kanji: ['pink.100', 'pink.800', 'pink.500', 'pink.50'],
//     vocabulary: ['purple.100', 'purple.800', 'purple.500', 'purple.50'],
//     default: ['gray.100', 'gray.800', 'gray.700', 'gray.50'],
//   };

//   const navLinkElements = menuItems
//     .filter((item) => !item.mobileOnly)
//     .map((item, i) => {
//       const { onOpen, onClose, isOpen } = useDisclosure();

//       const [bgColor, fgColor, menuBg, menuFg] =
//         item.text in colors ? colors[item.text] : colors['default'];

//       const element = (
//         <Button
//           variant="ghost"
//           _hover={{
//             backgroundColor: bgColor,
//             color: fgColor,
//           }}
//           py={2}
//           px={2}
//           rounded="sm"
//           fontWeight="normal"
//           textTransform={'capitalize'}
//           color="gray.700"
//         >
//           {item.text}
//         </Button>
//       );

//       if (item.isLink) {
//         return <Link href={item.href}>{element}</Link>;
//       }
//       return (
//         <Popover
//           key={i}
//           isOpen={isOpen}
//           onOpen={onOpen}
//           onClose={onClose}
//           placement="bottom"
//           closeOnBlur={true}
//         >
//           <PopoverTrigger>{element}</PopoverTrigger>
//           <PopoverContent
//             mt="1"
//             px={2}
//             py={3}
//             bgColor={menuBg}
//             color={menuFg}
//             border={'none'}
//             boxShadow="2xl"
//             w="full"
//           >
//             <PopoverArrow bgColor={menuBg} />
//             <Box maxH="calc(100vh - 90px)" overflowY={'auto'}>
//               {item.menu}
//             </Box>
//           </PopoverContent>
//         </Popover>
//       );
//     });

//   return <div className="flex flex-col" spacing={2}>{navLinkElements}</div className="flex flex-col">;
// }
