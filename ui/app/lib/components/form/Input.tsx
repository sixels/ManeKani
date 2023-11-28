import classNames from "classnames";
import {
	ComponentPropsWithRef,
	ComponentPropsWithoutRef,
	forwardRef,
} from "react";

type TextInputProps = ComponentPropsWithoutRef<"input"> & {
	label?: string;
	error?: string | null;
};
export const TextInput = forwardRef(
	(
		{ label, error, className, ...params }: TextInputProps,
		ref?: ComponentPropsWithRef<"input">["ref"],
	) => (
		<div className="w-full">
			<div className="relative">
				<input
					type="text"
					ref={ref}
					className={classNames(
						error
							? "border-red-600 focus:border-red-600"
							: "border-neutral-400 focus:border-neutral-900",
						"block px-2.5 pb-2.5 pt-4 bg-neutral-50 focus:bg-white placeholder-shown:bg-neutral-100 w-full text-sm text-gray-900  rounded-t-lg appearance-none focus:outline-none focus:ring-0  peer border-b-2",
						className,
					)}
					placeholder=" "
					{...params}
				/>
				<label
					htmlFor="display_name"
					className={classNames(
						error
							? "text-red-600"
							: "text-neutral-500 peer-focus:text-neutral-900",
						"absolute font-medium text-sm duration-300 transform -translate-y-4 scale-75 top-2 z-10 origin-[0] bg-transparent leading-none  px-2 peer-focus:px-1.5 peer-placeholder-shown:scale-100 peer-placeholder-shown:-translate-y-1/2 peer-placeholder-shown:top-1/2 peer-focus:top-2 peer-focus:scale-75 peer-focus:-translate-y-4 left-1",
					)}
				>
					{label}
					{params.required && (
						<>
							{" "}
							<span className="text-red-600">*</span>
						</>
					)}
				</label>
			</div>
			{error && <p className="mt-2 text-sm text-red-600">{error}</p>}
		</div>
	),
);
