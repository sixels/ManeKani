import classNames from "classnames";
import { ComponentPropsWithoutRef, useEffect } from "react";

type ModalProps = ComponentPropsWithoutRef<"div"> & {
	title?: string;
	isOpen?: boolean;
	onClose: () => void;
};
export function Modal({
	title,
	children,
	isOpen = false,
	onClose,
	className,
	...props
}: ModalProps) {
	useEffect(() => {
		document.body.classList.toggle("overflow-hidden", isOpen);
	}, [isOpen]);

	return (
		<>
			<div
				tabIndex={-1}
				aria-hidden={isOpen}
				className={classNames(
					isOpen ? "block" : "hidden",
					"fixed top-0 left-0 right-0 z-50 w-full p-4 bg-neutral-800/60 grid place-items-center overflow-x-hidden overflow-y-auto md:inset-0 h-full",
				)}
			>
				<div className="relative w-full max-w-2xl max-h-full">
					<div className="relative bg-white rounded-lg shadow w-full">
						<div className="flex items-center justify-between p-4 px-6 border-b rounded-t">
							<h3 className="text-xl font-medium text-neutral-900 ">{title}</h3>
							<button
								type="button"
								className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ml-auto inline-flex justify-center items-center "
								data-modal-hide="static-modal"
								onClick={onClose}
							>
								<svg
									className="w-3 h-3"
									aria-hidden="true"
									xmlns="http://www.w3.org/2000/svg"
									fill="none"
									viewBox="0 0 14 14"
								>
									<title>close icon</title>
									<path
										stroke="currentColor"
										strokeLinecap="round"
										strokeLinejoin="round"
										strokeWidth="2"
										d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
									/>
								</svg>
								<span className="sr-only">Close modal</span>
							</button>
						</div>
						{children}
					</div>
				</div>
			</div>

			{/* <Transition appear show={isOpen} as={Fragment}>
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
                    'w-full container max-w-5xl transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all',
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
      </Transition> */}
		</>
	);
}

export function ModalBody({
	className,
	children,
}: ComponentPropsWithoutRef<"div">) {
	return (
		<div className={classNames("w-full p-6 space-y-6", className)}>
			{children}
		</div>
	);
}

export function ModalFooter({
	className,
	children,
}: ComponentPropsWithoutRef<"div">) {
	return (
		<div
			className={classNames(
				"flex items-center justify-end p-6 space-x-2 border-t border-gray-200 rounded-b",
				className,
			)}
		>
			{children}
		</div>
	);
}
