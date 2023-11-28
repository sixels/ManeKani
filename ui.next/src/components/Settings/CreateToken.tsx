"use client";

import { useDisclosure } from "@/lib/component/utils";
import { Modal } from "@/stories/Client/Modal";
import { Button } from "@/stories/Primitives/Button";
import { IconButton } from "@/stories/Primitives/IconButton";
import { Combobox, Dialog, Transition } from "@headlessui/react";
import {
	IconCheck,
	IconCopy,
	IconSquare,
	IconSquareCheck,
} from "@tabler/icons-react";
import classNames from "classnames";
import { Token } from "@manekani/core";
import { Fragment, useEffect, useState } from "react";

const PERMISSIONS = [
	{ id: 1, name: "deck:create" },
	{ id: 2, name: "deck:delete" },
	{ id: 3, name: "deck:update" },
	{ id: 4, name: "subject:create" },
	{ id: 5, name: "subject:update" },
	{ id: 6, name: "subject:delete" },
	{ id: 7, name: "review:create" },
	{ id: 8, name: "study_material:create" },
	{ id: 9, name: "study_material:update" },
	{ id: 10, name: "user:update" },
];

export const CreateToken = () => {
	// const { createToken } = useTokens();

	const { isOpen, onClose, onOpen } = useDisclosure();
	const [token, setToken] = useState("");

	const dimiss = () => {
		onClose();
		setToken("");
	};

	const createTokenWrapper = async (
		name: string,
		permissions: typeof PERMISSIONS,
	) => {
		setToken("");
		try {
			// const created = await createToken(name, parsePermissions(permissions));
			// setToken(created.token || '');
			setToken("");
		} catch (e) {
			Promise.reject(e);
		}
	};

	return (
		<Modal
			button={<Button text="Create new token" isPrimary onClick={onOpen} />}
			isOpen={isOpen}
			onClose={onClose}
			title={token == "" ? "Create a new token" : "Token created"}
		>
			{token == "" ? (
				<CreateTokenForm createToken={createTokenWrapper} />
			) : (
				<ShowToken token={token} dimiss={dimiss} />
			)}
		</Modal>
	);
};

const CreateTokenForm = ({
	createToken,
}: {
	createToken: (name: string, permissions: typeof PERMISSIONS) => Promise<void>;
}) => {
	const [selectedPermissions, setSelectedPermissions] = useState<
		typeof PERMISSIONS
	>([]);
	const [tokenName, setTokenName] = useState("");
	const [query, setQuery] = useState("");
	const filteredPerms =
		query == ""
			? PERMISSIONS
			: PERMISSIONS.filter((perm) =>
					perm.name
						.toLowerCase()
						.replace(/\s+/g, "")
						.includes(query.toLowerCase().replace(/\s+/g, "")),
			  );

	const [isLoading, setIsLoading] = useState(false);

	const submitToken = () => {
		setIsLoading(true);
		createToken(tokenName, selectedPermissions)
			.catch((e) => {
				console.error(e);
			})
			.finally(() => {
				setIsLoading(false);
			});
	};

	return (
		<form
			onSubmit={(e) => {
				e.preventDefault();
				submitToken();
			}}
			className="mt-6"
		>
			<div className="flex flex-col gap-4">
				<div>
					<label
						htmlFor="name"
						className="block mb-2 text-sm font-medium text-gray-900"
					>
						Name
					</label>
					<input
						type="text"
						id="name"
						className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
						placeholder="What is this token for?"
						onChange={(e) => setTokenName(e.target.value)}
						required
					/>
				</div>
				<Combobox
					value={selectedPermissions}
					onChange={setSelectedPermissions}
					// @ts-ignore
					multiple
				>
					<div className="relative">
						<label
							htmlFor="name"
							className="block mb-2 text-sm font-medium text-gray-900"
						>
							Permissions
						</label>
						<p className="text-gray-500">
							Every token have read access to your data (being able to see your
							decks, cards, study materials and so on).
						</p>
						<p className="text-gray-500">
							By checking the options below, applications may have the power to
							create, update, and delete some of your data, according to the
							permissions you give to it. Only give access to applications you
							trust.
						</p>
						<p className="text-gray-500">
							View the{" "}
							<a
								href="#"
								className="text-wk-accent-400 hover:text-wk-accent-500 transition-colors hover:underline"
							>
								API documentation
							</a>{" "}
							to learn more about these endpoints and what information they
							return.
						</p>
						<Combobox.Input
							id="permissions"
							placeholder="Filter permissions"
							className="w-full mt-2 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5"
							onChange={(e) => setQuery(e.target.value)}
							autoComplete="off"
						/>
						<Combobox.Options
							static
							className=" mt-1 max-h-60 h-60 w-full overflow-auto py-1 text-base focus:outline-none sm:text-sm"
						>
							{filteredPerms.map((perm) => (
								<Combobox.Option
									key={perm.id}
									value={perm}
									className={({ active }) =>
										classNames(
											"relative cursor-default select-none py-2 pl-10 pr-4",
											active && "bg-wk-secondary-50",
										)
									}
								>
									{({ selected }) => (
										<>
											<span
												className={`block truncate ${
													selected ? "font-medium" : "font-normal"
												}`}
											>
												{perm.name}
											</span>
											<span className="absolute inset-y-0 left-0 flex items-center pl-3 ">
												{selected ? (
													<IconSquareCheck
														size={20}
														className="stroke-wk-accent-500"
														aria-hidden="true"
													/>
												) : (
													<IconSquare
														size={20}
														className="stroke-gray-500"
														aria-hidden="true"
													/>
												)}
											</span>
										</>
									)}
								</Combobox.Option>
							))}
						</Combobox.Options>
					</div>
				</Combobox>
			</div>
			<div className="mt-6">
				<Button
					// type="submit"
					isPrimary
					fillWidth
					iconOnLeft
					icon={
						isLoading ? (
							<svg
								aria-hidden="true"
								role="status"
								className="inline w-4 h-4 mr-3 text-gray-200 animate-spin"
								viewBox="0 0 100 101"
								fill="none"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
									fill="currentColor"
								/>
								<path
									d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
									className="fill-wk-primary-950"
								/>
							</svg>
						) : undefined
					}
					text={isLoading ? "Creating..." : "Create"}
				/>
			</div>
		</form>
	);
};

const ShowToken = ({
	token,
	dimiss,
}: {
	token: string;
	dimiss?: () => void;
}) => {
	const [copied, setCopied] = useState(false);
	const copyIcon = copied ? <IconCheck /> : <IconCopy />;
	const copyClasses = copied
		? "!text-wk-accent-500 !bg-wk-accent-50 hover:!bg-wk-accent-100"
		: "!text-gray-700";

	const copyToken = () => {
		navigator.clipboard
			.writeText(token.trim())
			.then(() => {
				setCopied(true);
			})
			.catch((e) => {
				console.error(e);
			});
	};

	useEffect(() => {
		if (copied) {
			setTimeout(() => {
				setCopied(false);
			}, 5000);
		}
	}, [copied]);

	return (
		<section className="flex flex-col gap-3 mt-2">
			<p className="text-gray-500">
				We do not store your tokens anywhere in our servers, please copy the
				token or write it down so can use it later.
			</p>
			<div className="inline-flex gap-2 flex-wrap md:flex-nowrap">
				<pre className="font-mono text-sm w-full md:text-base block p-2.5 bg-gray-50 border border-gray-300 text-gray-800 rounded-lg overflow-x-scroll sm:overflow-x-auto">
					{token.split("-").map((part, i) => (
						<span key={i}>
							<span
								className={classNames(
									"whitespace-nowrap",
									i == 0 && "text-wk-primary-800",
								)}
							>
								{part}
							</span>
							{i == 0 && "-"}
						</span>
					))}
				</pre>
				<IconButton
					className={classNames("rounded-lg hidden md:block", copyClasses)}
					icon={copyIcon}
					aria-label="Copy token"
					onClick={copyToken}
				/>
				<Button
					className={classNames(
						"font-medium rounded-lg md:hidden",
						copyClasses,
					)}
					icon={copyIcon}
					text="Copy"
					iconOnLeft
					fillWidth
					onClick={copyToken}
				/>
			</div>
			<div className="mt-1.5">
				<Button text="Close" onClick={dimiss} />
			</div>
		</section>
	);
};

const parsePermissions = (perms: typeof PERMISSIONS) => {
	const permMap = {
		"deck:create": "deckCreate",
		"deck:delete": "deckDelete",
		"deck:update": "deckUpdate",
		"subject:create": "subjectCreate",
		"subject:update": "subjectUpdate",
		"subject:delete": "subjectDelete",
		"review:create": "reviewCreate",
		"study_material:create": "studyMaterialCreate",
		"study_material:update": "studyMaterialUpdate",
		"user:update": "userUpdate",
	};

	let permissions: Token["claims"] = {} as Token["claims"];
	for (const perm of perms) {
		const permField = permMap[
			perm.name as keyof typeof permMap
		] as keyof typeof permissions;
		permissions[permField] = true;
	}
	return permissions;
};
