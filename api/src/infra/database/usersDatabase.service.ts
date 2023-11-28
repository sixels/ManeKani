import { UsersDatabase } from "@manekani/infra-db";
import { Injectable } from "@nestjs/common";
import { DatabaseService } from "./database.service";

export const UsersProviderLabel = "USERS_REPOSITORY";

@Injectable()
export class UsersDatabaseService extends UsersDatabase {
	// biome-ignore lint/complexity/noUselessConstructor: This is required by Nest
	constructor(db: DatabaseService) {
		super(db);
	}
}
