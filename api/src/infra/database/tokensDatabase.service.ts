import { TokensDatabase } from "@manekani/infra-db";
import { Injectable } from "@nestjs/common";
import { DatabaseService } from "./database.service";

@Injectable()
export class TokensDatabaseService extends TokensDatabase {
	// biome-ignore lint/complexity/noUselessConstructor: This is required by Nest
	constructor(db: DatabaseService) {
		super(db);
	}
}
