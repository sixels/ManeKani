import { FileStorage } from "@manekani/infra-files/dist";
import { Injectable, OnModuleInit } from "@nestjs/common";

export const FileStorageProviderLabel = "FILES_REPOSITORY";

@Injectable()
export class FileStorageService extends FileStorage implements OnModuleInit {
	constructor() {
		super({
			proxyUrl: process.env.PROXY_URL || "",
			s3Endpoint: process.env.S3_URL || "",
			s3AccessKey: process.env.S3_ACCESS_KEY || "",
			s3SecretKey: process.env.S3_SECRET_KEY || "",
			s3Port: parseInt(process.env.S3_PORT || "") || 9000,
		});
	}

	async onModuleInit() {
		await this.createBuckets();
	}
}
