import { TSchema, Type } from "@sinclair/typebox";

export type Optional<T, K extends keyof T> = Omit<T, K> & Partial<T>;

export const UuidSchema = Type.String({ format: "uuid" });

export const TypeSlug = (options: Parameters<typeof Type.RegExp>[1]) =>
	Type.RegExp(/^[a-z0-9]+(?:-[a-z0-9]+)*$/, options);

export const Nullable = <T extends TSchema>(schema: T) =>
	Type.Optional(Type.Union([schema, Type.Null()], { default: null }));
