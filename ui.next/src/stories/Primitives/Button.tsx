// import { ComponentProps } from "@/lib/component/interfaces";
import classNames from "classnames";
import Link from "next/link";
import React, { ComponentPropsWithRef } from "react";

interface ButtonProps extends ComponentPropsWithRef<"button"> {
  /**  The button text */
  text: string;
  /** Is this button primary? */
  isPrimary?: boolean;
  /** Optional icon to display next to the text */
  icon?: React.ReactNode;
  /** The icon side */
  iconOnLeft?: boolean;
  /** Should the button fill the container width? */
  fillWidth?: boolean;
  /** Should the button be a link */
  link?: string;
  /** The button size */
  size?: "xs" | "sm" | "md" | "lg" | "xl";
  /** Should the button be disabled? */
  disabled?: boolean;
  /** Optional callback function */
  onClick?: () => void;
}

const getModeClasses = (isPrimary: boolean) =>
  isPrimary
    ? "bg-wk-primary-500 hover:bg-wk-primary-600 active:bg-wk-primary-600 active:ring-wk-primary-600 focus:bg-wk-primary-600 focus:ring-wk-primary-600 focus:ring-2 active:ring-2 focus:ring-offset-2 active:ring-offset-2 text-wk-text"
    : "bg-wk-secondary-50 hover:bg-wk-secondary-100 active:bg-wk-secondary-100 active:ring-wk-text focus:bg-wk-secondary-100 focus:ring-wk-text active:ring-1 focus:ring-1 text-wk-text";

const getIconClasses = (iconOnLeft: boolean) =>
  iconOnLeft ? "flex-row-reverse" : "flex-row";

const getSizeClasses = (size: "xs" | "sm" | "md" | "lg" | "xl") =>
  ({
    xs: "text-xs py-2 px-3",
    sm: "text-sm py-2 px-4",
    md: "text-md py-2.5 px-6",
    lg: "text-lg py-2.5 px-6",
    xl: "text-xl py-3 px-6",
  }[size]);

export const Button = ({
  text = "",
  isPrimary = false,
  iconOnLeft = false,
  fillWidth = false,
  disabled = false,
  link,
  size = "md",
  icon,
  className,
  ...props
}: ButtonProps) => {
  const modeStyle = getModeClasses(isPrimary),
    iconStyle = getIconClasses(iconOnLeft),
    widthStyle = fillWidth ? "w-full justify-center font-medium" : "w-auto",
    sizeStyle = getSizeClasses(size);

  const classes = classNames(
    className,
    modeStyle,
    iconStyle,
    widthStyle,
    sizeStyle,
    "transition duration-200 rounded-full flex gap-2 items-center"
  );

  return link ? (
    <Link className={classes} aria-disabled={disabled} href={link}>
      {text}
      {icon && <> {icon}</>}
    </Link>
  ) : (
    <button className={classes} disabled={disabled} {...props}>
      <span>{text}</span>
      {icon && <> {icon}</>}
    </button>
  );
};
