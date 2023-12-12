import { BaseError, Token } from "@manekani/core";
import type { ActionFunctionArgs, LoaderFunctionArgs } from "@remix-run/node";
import { json } from "@remix-run/node";
import { useFetcher, useLoaderData } from "@remix-run/react";
import classNames from "classnames";
import {
	compareDesc,
	differenceInDays,
	format,
	formatDistanceToNow,
} from "date-fns";
import { Suspense, useEffect, useRef, useState } from "react";
import { TextInput } from "~/lib/components/form/Input";
import { Button, buttonStyles } from "~/lib/components/general/Button";
import { NavigateBack } from "~/lib/components/general/NavigateBack";
import { Modal, ModalBody, ModalFooter } from "~/lib/components/layout/Modal";
import { useDisclosure } from "~/lib/hooks/disclosure";
import { tokens } from "~/lib/infra/db/db.server";
import { requireCompletedUserSession } from "~/lib/util/session";
import { action as manageTokenAction } from "../settings.api-tokens.$id/route";

type Jsonify<T> = T extends Date
	? string
	: T extends object
	  ? {
				[k in keyof T]: Jsonify<T[k]>;
		  }
	  : T;

export function meta() {
	return [{ title: "API Tokens | ManeKani" }];
}

export async function loader({ request }: LoaderFunctionArgs) {
	const { user } = await requireCompletedUserSession(request);

	// const tokensAdapter = new TokensAdapter(new TokensDatabase(tokens));
	const userTokens = await tokens.getTokens(user.id);

	// const tokensAdapter = new TokensAdapter(tokensDb);
	// return json(tokensAdapter.getTokens(userId));
	const userInfo = {
		username: user.username,
		email: user.email,
		displayName: user.displayName,
		isVerified: user.isVerified,
	};

	return json({ userInfo, userTokens });
}

export async function action({ request }: ActionFunctionArgs) {
	const { user } = await requireCompletedUserSession(request);

	const formData = await request.formData();

	const fields = {
		name: formData.get("name")?.toString(),
		permissions: {
			deckCreate: formData.has("deck--create"),
			deckUpdate: formData.has("deck--update"),
			deckDelete: formData.has("deck--delete"),
			subjectCreate: formData.has("subject--create"),
			subjectUpdate: formData.has("subject--update"),
			subjectDelete: formData.has("subject--delete"),
			studyMaterialCreate: formData.has("study_material--create"),
			studyMaterialUpdate: formData.has("study_material--update"),
			studyMaterialDelete: formData.has("study_material--delete"),
			reviewCreate: formData.has("review--create"),
		},
	};

	const errors: Record<string, string | null> = {
		username: fields.name ? null : "Token name is required",
		display_name: null,
		other: null,
	};

	if (fields.name) {
		try {
			const token = await tokens.generateToken(user.id, {
				name: fields.name,
				claims: Object.assign(fields.permissions, { userUpdate: false }),
			});
			return json({ token: token.generatedToken, errors: null });
		} catch (error) {
			if (error instanceof BaseError) {
				// TODO: handle error
				console.error(error);
				errors.other =
					"An error occurred while creating the token. Please try again later.";
			}
		}
	}

	// await tokens.generateToken(user.id, {

	// });

	return json({ token: null, errors });
}

export default function Component() {
	const { userInfo, userTokens } = useLoaderData<typeof loader>();

	// TODO: separate layout mockup from component
	return (
		<section className="px-5 w-full">
			<div>
				<NavigateBack href="/settings" to="settings" />
				<h1 className="font-bold text-3xl text-neutral-900 py-2">API Tokens</h1>
				<p className="max-w-3xl">
					An API token can be used by third-party applications to improve your
					experience.
				</p>
			</div>

			<div className="w-full mt-6">
				<div className="bg-white rounded-sm border border-neutral-200 w-full ">
					<div className="flex items-center border-b border-neutral-200 py-3.5 px-6">
						<div className="flex-grow">
							<h2 className="font-medium text-xl text-neutral-900">
								ManeKani API Tokens
							</h2>
						</div>
						<CreateTokenButton />
					</div>
					<UserTokensTable tokens={userTokens} />
				</div>
			</div>
		</section>
	);
}

function CreateTokenButton() {
	// TODO: token creation logic
	const { isOpen, onClose, onOpen } = useDisclosure();

	const fetcher = useFetcher<typeof action>();

	const generatedToken = fetcher.data?.token;
	// Todo: display errors
	const errors = fetcher.data?.errors;

	const onOpenWrapper = () => {
		fetcher.data = undefined;
		onOpen();
	};

	// biome-ignore lint/correctness/useExhaustiveDependencies: we only need to do this when action data changes
	useEffect(() => {
		if (fetcher.data) {
			onOpen();
		}
	}, [fetcher.data]);

	// TODO: show errors

	return (
		<>
			<Button isPrimary onClick={onOpenWrapper}>
				Generate new API token
			</Button>
			<Modal
				isOpen={isOpen}
				onClose={onClose}
				title={generatedToken ? "Token created" : "Create a new API token"}
				className="px-0"
			>
				{generatedToken ? (
					<TokenCreatedModal token={generatedToken} onClose={onClose} />
				) : (
					<CreateTokenModal onClose={onClose} fetcher={fetcher} />
				)}
			</Modal>
		</>
	);
}

function TokenCreatedModal({
	onClose,
	token,
}: {
	onClose: () => void;
	token: string;
}) {
	const [copied, setCopied] = useState(false);
	const [copyTimeout, setCopyTimeout] = useState<NodeJS.Timeout | null>(null);
	const tokenInputRef = useRef<HTMLInputElement>(null);

	const copy = () => {
		setCopied(true);
		if (copyTimeout) {
			clearTimeout(copyTimeout);
		}
		setCopyTimeout(
			setTimeout(() => {
				setCopied(false);
			}, 8_000),
		);
	};

	const copyToken = () => {
		if (tokenInputRef.current) {
			try {
				tokenInputRef.current.focus();
				tokenInputRef.current.select();
				tokenInputRef.current.setSelectionRange(0, token.length);
			} catch {
				console.error("could not select the token content");
			}
		}

		navigator.clipboard
			.writeText(token)
			.then(() => {
				copy();
			})
			.catch((e) => {
				console.error(e);
			});
	};

	return (
		<>
			<ModalBody>
				<p className="text-neutral-600">
					You have successfully created a new API token! For security reasons,
					this token will only be shown once, make sure to copy it and store in
					a safe place.
				</p>
				<div className="flex w-full gap-2 flex-wrap md:flex-nowrap">
					<TextInput
						ref={tokenInputRef}
						className="font-mono select-text"
						defaultValue={token}
						onClick={copyToken}
						readOnly
					/>
				</div>
			</ModalBody>
			<ModalFooter>
				<Button isSecondary onClick={onClose}>
					Close
				</Button>
				<Button
					isPrimary
					onClick={copyToken}
					className={copied ? "bg-neutral-900" : undefined}
				>
					<span className="material-symbols-outlined text-lg">
						{copied ? "check_circle" : "content_copy"}
					</span>
					<span className="ml-2">{copied ? "Copied" : "Copy"}</span>
				</Button>
			</ModalFooter>
		</>
	);
}

function CreateTokenModal({
	fetcher,
	onClose,
}: {
	fetcher: ReturnType<typeof useFetcher<typeof action>>;
	onClose: () => void;
}) {
	// TODO: Add a tooltip for each permission resource

	return (
		<>
			<p className="text-neutral-600 pt-6 px-6">
				{/* Create a description telling what is an API token and its dangers */}
				API tokens are used to authenticate your application with the server.
				They are used to grant access to your data, so be careful with what
				permissions you give to each token.
			</p>
			<fetcher.Form method="post" className="w-full">
				<ModalBody>
					<div className="flex w-full flex-col gap-4">
						<TextInput id="name" name="name" required label="Token name" />
						<section className="w-full">
							<h3 className="block mb-2 text-sm font-medium text-gray-900">
								Permissions
							</h3>
							<p className="text-gray-500">
								Every token have read access to your data (being able to see
								your decks, cards, study materials and so on).
							</p>
							<p className="text-gray-500">
								By checking the options below, applications may have the power
								to create, update, and delete your data, according to the
								permissions you give to it.
							</p>
							<div className="mt-3 overflow-x-auto w-full">
								<PermissionsTable />
							</div>
						</section>
					</div>
				</ModalBody>

				<ModalFooter>
					<Button isSecondary onClick={onClose}>
						Cancel
					</Button>
					<Button isPrimary type="submit">
						Create
					</Button>
				</ModalFooter>
			</fetcher.Form>
		</>
	);
}

export function PermissionsTable({
	defaultValues,
	readonly,
}: {
	defaultValues?: {
		deck: { create: boolean; update: boolean; delete: boolean };
		subject: { create: boolean; update: boolean; delete: boolean };
		study_material: { create: boolean; update: boolean; delete: boolean };
		review: { create: boolean };
	};
	readonly?: boolean;
}) {
	const permissions = {
		deck: { create: true, update: true, delete: true },
		subject: { create: true, update: true, delete: true },
		study_material: { create: true, update: true, delete: true },
		review: { create: true, update: false, delete: false },
	} as const;
	return (
		<table className="text-sm text-left w-full rounded-sm text-neutral-500">
			<thead className="text-xs text-neutral-700 font-medium bg-neutral-50 border-b border-neutral-200">
				<tr>
					<th scope="col" className="px-6 py-4 whitespace-nowrap">
						Resource
					</th>
					<th scope="col" className="px-6 py-4 whitespace-nowrap text-center">
						Create
					</th>
					<th scope="col" className="px-6 py-4 whitespace-nowrap text-center">
						Modify
					</th>
					<th scope="col" className="px-6 py-4 whitespace-nowrap text-center">
						Delete
					</th>
				</tr>
			</thead>
			<tbody className="divide-y">
				{Object.entries(permissions).map(([resource, actions]) => (
					<tr
						key={resource}
						className="bg-white hover:bg-neutral-100 text-neutral-400"
					>
						<th
							scope="row"
							className="px-6 py-4 capitalize font-medium text-neutral-900 "
						>
							{resource.replace("_", " ")}
						</th>
						{Object.entries(actions).map(([action, allowed]) => (
							<td key={action} className="px-6 py-4 text-center">
								{" "}
								{allowed && (
									<input
										id={!readonly ? `${resource}--${action}` : undefined}
										name={!readonly ? `${resource}--${action}` : undefined}
										type="checkbox"
										className="w-5 h-5"
										disabled={readonly}
										readOnly={readonly}
										defaultChecked={
											defaultValues &&
											resource in defaultValues &&
											(
												defaultValues as Record<string, Record<string, boolean>>
											)[resource as keyof typeof defaultValues][action]
										}
									/>
								)}
							</td>
						))}
					</tr>
				))}
			</tbody>
		</table>
	);
}

function UserTokensTable({ tokens }: { tokens: Jsonify<Token>[] }) {
	// TODO: sort tokens dinamically

	const sortedTokens = tokens.sort((a, b) => {
		return compareDesc(new Date(a.createdAt), new Date(b.createdAt));
	});

	return (
		<div className="w-full overflow-x-auto">
			<div className="h-auto overflow-y-visible">
				<table className="w-full  text-left text-sm text-neutral-500">
					<thead className="text-xs text-neutral-700 font-medium bg-neutral-50 border-b border-neutral-200">
						<tr>
							<th scope="col" className="px-6 py-4 whitespace-nowrap">
								Name
							</th>
							<th scope="col" className="px-6 py-4 whitespace-nowrap">
								Prefix
							</th>
							<th scope="col" className="px-6 py-4 whitespace-nowrap">
								Created At
							</th>
							<th scope="col" className="px-6 py-4 whitespace-nowrap">
								Status
							</th>
							<th scope="col" className="px-6 py-4 whitespace-nowrap">
								Last Used
							</th>
							<th scope="col" className="px-6 py-4 whitespace-nowrap">
								<span className="sr-only">More</span>
							</th>
						</tr>
					</thead>
					<tbody className="divide-y">
						{tokens.length == 0 ? (
							<tr className="bg-white">
								<td
									colSpan={6}
									className="px-6 py-8 text-base font-medium text-neutral-500 whitespace-nowrap text-center"
								>
									Your API tokens will appear here after you create them.
								</td>
							</tr>
						) : (
							sortedTokens.map((tk) => (
								<tr
									key={tk.id}
									className="bg-white text-left hover:bg-neutral-100 text-neutral-600"
								>
									<th
										scope="row"
										className="px-6 py-3 font-medium text-neutral-900 whitespace-nowrap"
									>
										{tk.name}
									</th>
									<td className="px-6 py-3 text-neutral-900 font-mono">
										<pre>{tk.prefix}</pre>
									</td>
									<td className="px-6 py-3 first-letter:uppercase">
										<Suspense>
											{differenceInDays(new Date(), new Date(tk.createdAt)) > 0
												? format(new Date(tk.createdAt), "dd MMM yyyy")
												: `${formatDistanceToNow(new Date(tk.createdAt))} ago`}
										</Suspense>
										{/* {tk.createdAt} */}
									</td>
									<td className="px-6 py-3">
										<span
											className={classNames(
												tk.isActive
													? "text-green-800 bg-green-100"
													: "text-amber-800 bg-amber-100",
												"rounded-sm px-2.5  font-medium py-1.5",
											)}
										>
											{tk.isActive ? "Active" : "Paused"}
										</span>
									</td>
									<td className="px-6 py-3">{tk.usedAt ?? "Never"}</td>
									<td className="px-6 py-3 relative">
										<ManageTokenButton token={tk} />
									</td>
								</tr>
							))
						)}
					</tbody>
				</table>
			</div>
		</div>
	);
}

function ManageTokenButton({ token }: { token: Jsonify<Token> }) {
	const statusFetcher = useFetcher<typeof manageTokenAction>();
	const deleteFetcher = useFetcher<typeof manageTokenAction>();

	const toggleToken = async () => {
		const data: {
			action: "save";
			token_name: string;
			token_is_active?: "on";
		} = {
			action: "save",
			token_name: token.name,
			...(token.isActive ? {} : { token_is_active: "on" }),
		};

		statusFetcher.submit(data, {
			action: `/settings/api-tokens/${token.id}`,
			method: "post",
		});
	};
	const isUpdatingStatus = statusFetcher.state != "idle";

	const deleteToken = async () => {
		const data: {
			action: "delete";
		} = {
			action: "delete",
		};

		deleteFetcher.submit(data, {
			action: `/settings/api-tokens/${token.id}`,
			method: "post",
		});
	};
	const isDeleting = deleteFetcher.state != "idle";

	// TODO: error handling
	// TODO: delete confirmation

	return (
		<div className="flex gap-0.5">
			<a
				href={`/settings/api-tokens/${token.id}`}
				className={classNames(
					buttonStyles({ isSecondary: true }),
					"!h-8 focus:z-[0]",
				)}
			>
				Manage
			</a>
			<div className="group relative">
				<Button
					isSecondary
					className="rounded-l-none !h-8 !px-2 z-0 !ring-0"
					aria-label="Token quick actions"
				>
					<svg
						className="w-3 h-3"
						viewBox="0 0 24 24"
						fill="none"
						xmlns="http://www.w3.org/2000/svg"
					>
						<title>more actions icon</title>
						<g id="SVGRepo_bgCarrier" strokeWidth="0" />
						<g
							id="SVGRepo_tracurrentColorerCarrier"
							stroke-linecurrentcap="round"
							strokeLinejoin="round"
						/>
						<g id="SVGRepo_icurrentColoronCarrier">
							<path
								d="M4 9.5C5.38071 9.5 6.5 10.6193 6.5 12C6.5 13.3807 5.38071 14.5 4 14.5C2.61929 14.5 1.5 13.3807 1.5 12C1.5 10.6193 2.61929 9.5 4 9.5Z"
								fill="currentColor"
							/>
							<path
								d="M12 9.5C13.3807 9.5 14.5 10.6193 14.5 12C14.5 13.3807 13.3807 14.5 12 14.5C10.6193 14.5 9.5 13.3807 9.5 12C9.5 10.6193 10.6193 9.5 12 9.5Z"
								fill="currentColor"
							/>
							<path
								d="M22.5 12C22.5 10.6193 21.3807 9.5 20 9.5C18.6193 9.5 17.5 10.6193 17.5 12C17.5 13.3807 18.6193 14.5 20 14.5C21.3807 14.5 22.5 13.3807 22.5 12Z"
								fill="currentColor"
							/>
						</g>
					</svg>
				</Button>

				<div className="absolute rounded-l-sm px-0 right-0 bottom-[calc(100%_+_4px)] transition-all invisible opacity-0  group-focus-within:visible group-focus-within:opacity-100 z-10">
					<ul className="flex justify-right bg-neutral-100 rounded-sm shadow  text-sm text-neutral-800">
						<li className="flex-grow h-full">
							<Button
								disabled={isUpdatingStatus}
								className={classNames(
									buttonStyles({ isSecondary: true }),
									"flex w-full py-1.5 pl-2 pr-3 items-center focus:!ring-0 border-0 bg-white !h-full hover:bg-neutral-200 transition-colors",
								)}
								onClick={toggleToken}
								aria-label="toggle token status"
							>
								<span className="material-symbols-outlined mr-2">
									{token.isActive ? "pause" : "play_arrow"}
								</span>
								{token.isActive ? "Pause" : "Resume"}
							</Button>
						</li>
						<li className="h-full">
							<Button
								disabled={isDeleting}
								className={classNames(
									buttonStyles({ isSecondary: true }),
									"flex items-center !px-2 py-1.5 !w-auto !h-full focus:!ring-0 border-0 !bg-red-600 !text-white hover:!bg-red-700 transition-colors",
								)}
								onClick={deleteToken}
								aria-label="delete token"
							>
								<span className="material-symbols-outlined">
									delete_forever
								</span>
							</Button>
						</li>
					</ul>
				</div>
			</div>
		</div>
	);
}
