import classNames from "classnames";
import React, { ComponentPropsWithRef } from "react";

interface IconButtonProps extends ComponentPropsWithRef<"button"> {
	/** The button icon */
	icon: React.ReactNode;
	/** The button size */
	size?: "sm" | "md" | "lg";
	/** Is this button primary */
	isPrimary?: boolean;
}

const getModeClasses = (isPrimary: boolean) =>
	isPrimary
		? "bg-wk-primary-400 hover:bg-wk-primary-500 active:bg-wk-primary-500 active:ring-wk-primary-500 focus:bg-wk-primary-500 focus:ring-wk-primary-500 focus:ring-2 active:ring-2 focus:ring-offset-2 active:ring-offset-2 text-wk-text"
		: "bg-wk-secondary-50 hover:bg-wk-secondary-100 active:bg-wk-secondary-100 active:ring-wk-text focus:bg-wk-secondary-100 focus:ring-wk-text active:ring-1 focus:ring-1 text-wk-text";

export const IconButton = ({
	icon,
	size = "md",
	isPrimary = false,
	className,
	type,
	...props
}: IconButtonProps) => {
	const modeClasses = getModeClasses(isPrimary);
	const sizeClasses = {
		sm: "p-2",
		md: "p-2.5",
		lg: "p-3",
	}[size];

	return (
		<button
			type={type}
			className={classNames(
				modeClasses,
				sizeClasses,
				"transition duration-200 rounded-full flex gap-2 justify-center items-center",
				className,
			)}
			{...props}
		>
			{icon}
		</button>
	);
};
