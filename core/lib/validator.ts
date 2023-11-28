import { Static, TSchema, Type } from "@sinclair/typebox";
import Ajv, { ValidateFunction } from "ajv";

import addErrors from "ajv-errors";
import addFormats from "ajv-formats";
import { InvalidRequestError } from "./domain/error";

const FAIL_MESSAGE = "Schema validation failed";

export class Validator<S extends TSchema, T = Static<S>> {
	private readonly check: ValidateFunction<S>;

	constructor(schema: S) {
		const ajv = addFormats(
			addErrors(
				new Ajv({
					useDefaults: true,
					allErrors: true,
				}),
			),
			["date", "date-time", "regex", "uuid", "time", "email"],
		);
		this.check = ajv.compile(schema);
	}

	isValid(data: unknown): data is T {
		return this.check(data);
	}

	validate(data: unknown): T {
		if (this.isValid(data)) {
			return data;
		}

		const errors = this.check.errors;
		const errorsString =
			errors?.map(({ message }) => message).join("; ") ?? FAIL_MESSAGE;

		throw new InvalidRequestError({
			cause: new Error(FAIL_MESSAGE),
			context: { errors },
			description: errorsString,
		});
	}
}

const UuidSchema = Type.String({ format: "uuid" });
export const UuidValidator = new Validator(UuidSchema);
