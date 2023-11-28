import { DecksDatabase } from "@manekani/infra-db";
import { Injectable } from "@nestjs/common";
import { DatabaseService } from "./database.service";

export const DecksProviderLabel = "DECKS_REPOSITORY";

@Injectable()
export class DecksDatabaseService extends DecksDatabase {
	// biome-ignore lint/complexity/noUselessConstructor: This is required by Nest
	constructor(db: DatabaseService) {
		super(db);
	}
}
