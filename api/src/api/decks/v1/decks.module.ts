import {
	DecksDatabaseService,
	DecksProviderLabel,
} from "@/infra/database/decksDatabase.service";
import {
	SubjectsDatabaseService,
	SubjectsProviderLabel,
} from "@/infra/database/subjectsDatabase.service";

import { AuthModule } from "@/api/auth/auth.module";
import { DatabaseModule } from "@/infra/database/database.module";
import { Module } from "@nestjs/common";
import { DecksResolver } from "./decks.resolvers";

@Module({
	imports: [AuthModule, DatabaseModule],
	controllers: [],
	providers: [
		{ provide: DecksProviderLabel, useExisting: DecksDatabaseService },
		{ provide: SubjectsProviderLabel, useExisting: SubjectsDatabaseService },
		DecksResolver,
	],
	exports: [DecksResolver],
})
export class DecksV1Module {}
