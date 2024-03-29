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
import { SubjectsResolver } from "./subjects.resolvers";
import {
	FileStorageProviderLabel,
	FileStorageService,
} from "@/infra/files/files.service";
import { FileStorageModule } from "@/infra/files/files.module";

@Module({
	imports: [AuthModule, DatabaseModule, FileStorageModule],
	controllers: [],
	providers: [
		{ provide: SubjectsProviderLabel, useExisting: SubjectsDatabaseService },
		{ provide: DecksProviderLabel, useExisting: DecksDatabaseService },
		{ provide: FileStorageProviderLabel, useClass: FileStorageService },
		SubjectsResolver,
	],
	exports: [SubjectsResolver],
})
export class SubjectsV1Module {}
