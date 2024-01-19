import { TSchema, Type } from "@sinclair/typebox";

export type Optional<T, K extends keyof T> = Omit<T, K> & Partial<T>;

export const UuidSchema = Type.String({ format: "uuid" });

export const TypeSlug = (options: Parameters<typeof Type.RegExp>[1]) =>
	Type.RegExp(/^(?:[^\0-\x7F]|\w)+(?:-([^\0-\x7F]|\w|\d)+)*$/, options);

export const Nullable = <T extends TSchema>(schema: T) =>
	Type.Optional(Type.Union([schema, Type.Null()], { default: null }));
