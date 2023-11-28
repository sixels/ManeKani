import classNames from "classnames";
import { ComponentPropsWithoutRef } from "react";

type NavigateBackProps = ComponentPropsWithoutRef<"a"> & { to: string };
export function NavigateBack({ to, className, ...props }: NavigateBackProps) {
	return (
		<a
			className={classNames(
				"text-gray-600 font-medium text-sm hover:border-opacity-100 border-b-2 px-1 py-0.5 border-opacity-0 border-gray-900 hover:text-gray-900 transition-colors",
				className,
			)}
			{...props}
		>
			<svg
				className="w-7 h-4 inline mr-2 transition-colors"
				xmlns="http://www.w3.org/2000/svg"
				viewBox="0 0 28 16"
				fill="none"
			>
				<title>back arrow</title>
				<path
					stroke="currentColor"
					strokeLinecap="round"
					strokeLinejoin="round"
					strokeWidth="1.5"
					d="m28 7H1m4-4-4 4 4 4"
				/>
			</svg>
			Back to {to}
		</a>
	);
}
