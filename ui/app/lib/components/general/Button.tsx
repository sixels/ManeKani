import classNames from 'classnames';
import { ComponentPropsWithRef } from 'react';

type ButtonProps = ComponentPropsWithRef<'button'> & {
  isPrimary?: boolean;
  isSecondary?: boolean;
};

export function buttonStyles({
  isPrimary = false,
  isSecondary = false,
}: Pick<ButtonProps, 'isPrimary' | 'isSecondary'>) {
  return classNames(
    isPrimary
      ? 'text-white bg-neutral-800  hover:bg-neutral-900  disabled:hover:bg-neutral-800 focus:ring-4 focus:ring-neutral-300 font-medium px-5'
      : isSecondary
      ? 'text-neutral-600 bg-white hover:bg-neutral-100 disabled:hover:bg-white hover:text-neutral-600 border border-neutral-200 focus:ring-4 focus:ring-neutral-300 font-medium px-5'
      : 'TODO',

    'whitespace-nowrap h-11 flex items-center ring-offset-1 rounded-sm text-sm focus:outline-none disabled:opacity-50 disabled:cursor-not-allowed',
  );
}

export function Button({
  isPrimary = false,
  isSecondary = false,
  children,
  className,
  ...props
}: ButtonProps) {
  return (
    <button
      type="button"
      className={classNames(
        className,
        buttonStyles({ isPrimary, isSecondary }),
      )}
      {...props}
    >
      {children}
    </button>
  );
}
