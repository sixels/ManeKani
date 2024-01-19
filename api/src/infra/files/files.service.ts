import { FileStorage } from "@manekani/infra-files/dist";
import { Injectable, OnModuleInit } from "@nestjs/common";

export const FileStorageProviderLabel = "FILES_REPOSITORY";

@Injectable()
export class FileStorageService extends FileStorage implements OnModuleInit {
	constructor() {
		super({
			cdnUrl: process.env.CDN_URL || "",
			s3Endpoint: process.env.S3_URL || "",
			s3AccessKey: process.env.S3_ACCESS_KEY || "",
			s3SecretKey: process.env.S3_SECRET_KEY || "",
			s3Port: 9000,
		});
	}

	async onModuleInit() {
		await this.createBuckets();
	}
}
