"use client";

import { Dialog, Transition } from "@headlessui/react";
import classNames from "classnames";
import { ComponentPropsWithoutRef, Fragment } from "react";

interface ModalProps extends ComponentPropsWithoutRef<"div"> {
	button: React.ReactNode;
	title?: string;
	isOpen?: boolean;
	onClose: () => void;
}

export const Modal = ({
	button,
	isOpen = false,
	title,
	children,
	className,
	onClose,
	...props
}: ModalProps) => {
	return (
		<>
			{button}
			<Transition appear show={isOpen} as={Fragment}>
				<Dialog as="div" className="relative z-30" onClose={onClose}>
					<Overlay />
					<div className="fixed inset-0 overflow-y-auto">
						<div className="flex min-h-full items-center justify-center p-4 text-center">
							<Transition.Child
								as={Fragment}
								enter="ease-out duration-300"
								enterFrom="opacity-0 scale-95"
								enterTo="opacity-100 scale-100"
								leave="ease-in duration-200"
								leaveFrom="opacity-100 scale-100"
								leaveTo="opacity-0 scale-95"
							>
								<Dialog.Panel
									as="div"
									className={classNames(
										"w-full container max-w-5xl transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all",
										className,
									)}
									{...props}
								>
									{title && (
										<Dialog.Title
											as="h3"
											className="text-lg font-medium leading-6 text-wk-text"
										>
											{title}
										</Dialog.Title>
									)}
									{children}
								</Dialog.Panel>
							</Transition.Child>
						</div>
					</div>
				</Dialog>
			</Transition>
		</>
	);
};

const Overlay = () => {
	return (
		<Transition.Child
			as={Fragment}
			enter="ease-out duration-300"
			enterFrom="opacity-0"
			enterTo="opacity-100"
			leave="ease-in duration-200"
			leaveFrom="opacity-100"
			leaveTo="opacity-0"
		>
			<div className="fixed inset-0 bg-black bg-opacity-30" />
		</Transition.Child>
	);
};
