import { HttpStatus } from "@nestjs/common";

export class Response<T> {
	constructor(
		readonly data?: T,
		readonly statusCode: HttpStatus = HttpStatus.OK,
	) {}
}
