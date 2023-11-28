import { SubjectsDatabase } from "@manekani/infra-db";
import { Injectable } from "@nestjs/common";
import { DatabaseService } from "./database.service";

export const SubjectsProviderLabel = "SUBJECTS_REPOSITORY";

@Injectable()
export class SubjectsDatabaseService extends SubjectsDatabase {
	// biome-ignore lint/complexity/noUselessConstructor: This is required by Nest
	constructor(db: DatabaseService) {
		super(db);
	}
}
