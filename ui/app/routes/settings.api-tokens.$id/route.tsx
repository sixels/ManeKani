import { ForbiddenError, ResourceNotFoundError } from "@manekani/core";
import { LoaderFunctionArgs, json, redirect } from "@remix-run/node";
import { useFetcher, useLoaderData } from "@remix-run/react";
import classNames from "classnames";
import { useEffect, useState } from "react";
import { TextInput } from "~/lib/components/form/Input";
import { Button } from "~/lib/components/general/Button";
import { NavigateBack } from "~/lib/components/general/NavigateBack";
import { tokens } from "~/lib/infra/db/db.server";
import { requireCompleteUserSession } from "~/lib/util/session";
import { PermissionsTable } from "../settings.api-tokens._index/route";

export async function loader({ request, params }: LoaderFunctionArgs) {
	const { user } = await requireCompleteUserSession(request);

	try {
		const userToken = await tokens.getToken(user.id, params.id || "");
		return json({ userToken });
	} catch (error) {
		console.error(error);

		if (
			error instanceof ResourceNotFoundError ||
			error instanceof ForbiddenError
		) {
			throw new Response(null, { status: 404 });
		}
		throw new Response(null, { status: 500 });
	}
}

export async function action({ request, params }: LoaderFunctionArgs) {
	const { user } = await requireCompleteUserSession(request);

	const formData = await request.formData();

	const action = formData.get("action");
	const formFields = {
		tokenName: formData.get("token_name")?.toString(),
		tokenIsActive: formData.get("token_is_active") != null,
	};
	console.log(formFields);

	console.log(
		typeof formData.get("token_is_active"),
		formData.get("token_is_active") == null,
		formData.get("token_is_active"),
	);

	if (!action) {
		throw new Response(null, { status: 400 });
	}

	if (action == "save") {
		if (!formFields.tokenName) {
			return json(
				{ error: "Token name is required", userToken: null },
				{ status: 400 },
			);
		}

		try {
			const userToken = await tokens.updateToken(user.id, params.id || "", {
				name: formFields.tokenName,
				isActive: formFields.tokenIsActive,
			});
			return json({ error: null, userToken });
		} catch (error) {
			console.error(error);
			return json(
				{ error: "Failed to update token", userToken: null },
				{ status: 500 },
			);
		}
	} else if (action == "delete") {
		try {
			await tokens.deleteToken(user.id, params.id || "");
			// return json({
			//   error: null,
			//   userToken: null,
			// });
			throw redirect("/settings/api-tokens", { status: 303 });
		} catch (error) {
			console.error(error);
			return json(
				{ error: "Failed to delete token", userToken: null },
				{ status: 500 },
			);
		}
	}
}

export default function Component() {
	// const { userToken } = useLoaderData<typeof loader>();
	const fetcher = useFetcher<typeof action>();

	const loaderData = useLoaderData<typeof loader>();
	const [userToken, setUserToken] = useState(loaderData.userToken);

	const [tokenName, setTokenName] = useState(userToken.name);
	const [tokenIsActive, setTokenIsActive] = useState(userToken.isActive);

	const isUpdating = fetcher.state != "idle";

	useEffect(() => {
		const data = fetcher.data;
		if (data?.userToken) {
			setUserToken((prev) => ({
				...prev,
				name: data.userToken.name,
				isActive: data.userToken.isActive,
			}));
		}
	}, [fetcher.data]);

	// TODO: delete confirmation
	// TODO: error handling

	return (
		<section className="px-5 w-full">
			<fetcher.Form method="post" className="space-y-6">
				<div className="flex gap-4">
					<div className="w-full">
						<NavigateBack href="/settings/api-tokens" to="API tokens" />
						<h1 className="font-bold text-3xl text-neutral-900 py-2 ">
							Manage API Token
						</h1>
						<p className="max-w-3xl">View or update token data</p>
					</div>
					<div className="flex w-max items-center float-right">
						<div className="flex items-center justify-end gap-6">
							<Button
								type="submit"
								name="action"
								value="save"
								isPrimary
								className="transition-opacity"
								disabled={
									isUpdating ||
									(tokenIsActive == userToken.isActive &&
										(tokenName.length == 0 || tokenName == userToken.name))
								}
							>
								Save Changes
							</Button>
						</div>
					</div>
				</div>
				{/* <div className="bg-white rounded-sm border border-neutral-200 w-full "> */}
				<TextInput
					label="Token name"
					id="token_name"
					name="token_name"
					value={tokenName}
					onChange={(e) => {
						setTokenName(e.target.value);
					}}
				/>
				<div>
					<span className="block text-sm font-medium mb-2 text-neutral-500 rounded-sm overflow-hidden">
						Status
					</span>
					<label className="relative inline-flex h-10 select-none items-center cursor-pointer gap-6 bg-neutral-100 z-[0]">
						<input
							id="token_is_active"
							name="token_is_active"
							type="checkbox"
							defaultValue=""
							checked={tokenIsActive}
							className="sr-only peer"
							onChange={(e) => {
								setTokenIsActive(e.target.checked);
							}}
						/>
						<div className="absolute w-full h-10  rounded-sm peer peer-checked:after:-translate-x-full after:content-[''] after:absolute after:right-0 after:bg-neutral-800 peer-active:after:bg-neutral-900 peer-focus:ring peer-active:ring ring-offset-1 ring-neutral-200 peer-hover:after:bg-neutral-900 -z-[1]  after:rounded-sm after:h-full after:w-1/2 after:transition-all" />
						<span
							className={classNames(
								tokenIsActive ? "text-white" : "text-neutral-700",
								"text-sm flex-grow font-medium px-3",
							)}
						>
							Active
						</span>
						<span
							className={classNames(
								!tokenIsActive ? "text-white" : "text-neutral-700",
								"px-3 text-sm flex-grow font-medium",
							)}
						>
							Paused
						</span>
					</label>
				</div>
				<div>
					<span className="block text-sm font-medium mb-2 text-neutral-500 rounded-sm overflow-hidden">
						Permissions (read-only)
					</span>
					<PermissionsTable
						readonly
						defaultValues={{
							deck: {
								create: userToken.claims.deckCreate,
								update: userToken.claims.deckUpdate,
								delete: userToken.claims.deckDelete,
							},
							subject: {
								create: userToken.claims.subjectCreate,
								update: userToken.claims.subjectUpdate,
								delete: userToken.claims.subjectDelete,
							},
							study_material: {
								create: userToken.claims.studyMaterialCreate,
								update: userToken.claims.studyMaterialUpdate,
								delete: userToken.claims.studyMaterialDelete,
							},
							review: {
								create: userToken.claims.reviewCreate,
							},
						}}
					/>
				</div>

				<Button
					isPrimary
					type="submit"
					name="action"
					value="delete"
					disabled={isUpdating}
					className="!mt-12 bg-red-600 hover:bg-red-700  !ring-red-200 active:bg-red-700 "
				>
					Delete Token
				</Button>
			</fetcher.Form>
		</section>
	);
}
