"use client";

import { useDisclosure } from "@/lib/component/utils";
import { Modal } from "@/stories/Client/Modal";
import { Button } from "@/stories/Primitives/Button";
import { Spinner } from "@/stories/Primitives/Spinner";
import { Token } from "@manekani/core";
import { useState } from "react";

export const UserTokens = ({ tokens }: { tokens: Token[] }) => {
	if (tokens.length == 0) {
		return (
			<div className="w-full items-center inline-flex justify-center pt-6 text-lg font-medium text-gray-500">
				Your API tokens will appear here after you create them.
			</div>
		);
	}

	return (
		<div className="relative overflow-x-auto  sm:rounded-lg">
			<table className="w-full text-sm text-left text-gray-500">
				<thead className="text-xs text-gray-700 uppercase bg-wk-secondary-50">
					<tr>
						<th scope="col" className="px-6 py-3">
							Name
						</th>
						<th scope="col" className="px-6 py-3">
							Prefix
						</th>
						<th scope="col" className="px-6 py-3">
							Created At
						</th>
						<th scope="col" className="px-6 py-3">
							Status
						</th>
						<th scope="col" className="px-6 py-3">
							Last Used
						</th>
						<th scope="col" className="px-6 py-3">
							<span className="sr-only">More</span>
						</th>
					</tr>
				</thead>
				<tbody className="divide-y">
					{tokens.map((tk) => (
						<tr
							key={tk.id}
							className="bg-white hover:bg-wk-secondary-100 text-wk-text"
						>
							<th
								scope="row"
								className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
							>
								{tk.name}
							</th>
							<td className="px-6 py-4 text-wk-primary-800 font-mono">
								<pre>{tk.prefix}</pre>
							</td>
							<td className="px-6 py-4">Just now</td>
							<td className="px-6 py-4 capitalize">
								{tk.isActive ? "active" : "paused"}
							</td>
							<td className="px-6 py-4">Never</td>
							<td className="px-6 py-4">
								<TokenDetails token={tk} />
							</td>
						</tr>
					))}
				</tbody>
			</table>
		</div>
	);
};

export const Skeleton = ({ rows }: { rows: number }) => {
	const skeletonRows = new Array(rows).fill(0).map((_, i) => (
		<tr key={i}>
			<th scope="row" className="px-6 py-4">
				<div className="h-2.5 bg-gray-300 rounded-full w-20 mb-2.5"></div>
			</th>
			<td className="px-6 py-4">
				<div className="h-2.5 bg-wk-primary-400 rounded-full w-24 mb-2.5"></div>
			</td>
			<td className="px-6 py-4">
				<div className="h-2.5 bg-gray-200 rounded-full w-20 mb-2.5"></div>
			</td>
			<td className="px-6 py-4">
				<div className="h-2.5 bg-gray-200 rounded-full w-20 mb-2.5"></div>
			</td>
			<td className="px-6 py-4">
				<div className="h-2.5 bg-gray-200 rounded-full w-20 mb-2.5"></div>
			</td>
			<td className="px-6 py-4">
				<div className="h-2.5 bg-wk-accent-300 rounded-full w-20 mb-2.5"></div>
			</td>
		</tr>
	));

	return (
		<div
			role="status"
			className="relative sm:rounded-lg animate-pulse overflow-x-hidden"
		>
			<table className="w-full bg-white text-sm text-left text-gray-500">
				<thead className="bg-wk-secondary-50">
					<tr>
						<th scope="col" className="px-6 py-3">
							<div className="h-2.5 bg-gray-300 rounded-full w-24 mb-2.5"></div>
						</th>
						<th scope="col" className="px-6 py-3">
							<div className="h-2.5 bg-gray-300 rounded-full w-24 mb-2.5"></div>
						</th>
						<th scope="col" className="px-6 py-3">
							<div className="h-2.5 bg-gray-300 rounded-full w-24 mb-2.5"></div>
						</th>
						<th scope="col" className="px-6 py-3">
							<div className="h-2.5 bg-gray-300 rounded-full w-24 mb-2.5"></div>
						</th>
						<th scope="col" className="px-6 py-3">
							<div className="h-2.5 bg-gray-300 rounded-full w-24 mb-2.5"></div>
						</th>
						<th scope="col" className="px-6 py-3"></th>
					</tr>
				</thead>
				<tbody className="divide-y">{skeletonRows}</tbody>
			</table>
		</div>
	);
};

const TokenDetails = ({ token }: { token: Token }) => {
	// const { deleteToken } = useTokens();
	const { isOpen, onOpen, onClose } = useDisclosure();

	const [isDeleting, setIsDeleting] = useState(false);

	const deleteTokenWrapper = () => {
		// TODO: show confirmation dialog
		setIsDeleting(true);
		if (token.id) {
			// deleteToken(token.id)
			// .then(() => {
			//   onClose();
			// })
			// .catch((e) => {
			//   console.error(e);
			// })
			// .finally(() => {
			//   setIsDeleting(false);
			// });
		}
	};

	return (
		<Modal
			button={
				<button
					className="font-medium text-wk-accent-600 dark:text-wk-accent-500 hover:underline"
					onClick={onOpen}
				>
					More
				</button>
			}
			isOpen={isOpen}
			onClose={onClose}
			title={`Details for "${token.name}"`}
		>
			<section className="space-y-4 mt-4">
				<div className="flex gap-1.5">
					<span className="font-medium">Status</span>
					<span className="text-gray-700 capitalize">
						{token.isActive ? "active" : "paused"}
					</span>
				</div>
				<div className="flex flex-col gap-1.5">
					<span className="font-medium">Claims</span>
					<pre className="w-full font-mono overflow-auto rounded-lg bg-gray-50 border border-gray-300  p-2">
						{JSON.stringify(token, null, 2)}
					</pre>
				</div>
				<div className="inline-flex justify-between w-full">
					<Button text="Close" onClick={onClose} />
					<div className="inline-flex gap-2">
						<Button text="Freeze" />
						<Button
							text="Delete"
							icon={
								isDeleting ? (
									<Spinner thumbFillColor="fill-red-950" />
								) : undefined
							}
							iconOnLeft
							className="font-medium !text-white !bg-red-500 hover:!bg-red-600"
							onClick={deleteTokenWrapper}
						/>
					</div>
				</div>
			</section>
		</Modal>
	);
};
