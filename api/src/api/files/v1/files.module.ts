import { AuthModule } from "@/api/auth/auth.module";
import { DatabaseModule } from "@/infra/database/database.module";
import { FileStorageModule } from "@/infra/files/files.module";
import {
	FileStorageProviderLabel,
	FileStorageService,
} from "@/infra/files/files.service";
import { Module } from "@nestjs/common";
import { FilesController } from "./files.controller";
import { FilesService } from "./files.service";

@Module({
	imports: [AuthModule, DatabaseModule, FileStorageModule],
	controllers: [FilesController],
	providers: [
		{ provide: FileStorageProviderLabel, useExisting: FileStorageService },
		FilesService,
	],
	exports: [],
})
export class FilesV1Module {}
