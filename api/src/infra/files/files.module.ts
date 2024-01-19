import { Module } from "@nestjs/common";
import { FileStorageService } from "./files.service";

@Module({
	providers: [FileStorageService],
	exports: [FileStorageService],
})
export class FileStorageModule {}
