import {
	CreateFileDto,
	CreatedFile,
	FileType,
	GetFileUrlDto,
} from "@manekani/core";
import { IFilesRepositoryV1 } from "@manekani/core";
import { Client } from "minio";

const FileTypeBucket: Record<FileType, string> = {
	[FileType.attachments]: "attachments",
	[FileType.avatars]: "avatars",
};

type RepoConfig = {
	cdnUrl: string;

	s3Endpoint: string;
	s3Port: 9000;
	s3AccessKey: string;
	s3SecretKey: string;
};

export class FileStorage implements IFilesRepositoryV1 {
	private minio: Client;
	private cdnUrl: string;

	constructor(config: RepoConfig) {
		this.cdnUrl = config.cdnUrl;
		this.minio = new Client({
			endPoint: config.s3Endpoint || "",
			port: config.s3Port,
			accessKey: config.s3AccessKey || "",
			secretKey: config.s3SecretKey || "",
			useSSL: process.env.NODE_ENV === "production",
		});
	}

	async createBuckets() {
		for (const bucket of Object.values(FileTypeBucket)) {
			const bucketExists = await this.minio.bucketExists(bucket);
			if (!bucketExists) {
				console.log(`Creating s3 bucket ${bucket}`);
				await this.minio.makeBucket(bucket, "sa-east1", {
					ObjectLocking: true,
				});
			}
		}
	}

	async v1CreateFile(data: CreateFileDto): Promise<CreatedFile> {
		const uploadFilename = data.objectName;
		const uploadUrl = await this.minio.presignedUrl(
			"PUT",
			FileTypeBucket[data.type],
			uploadFilename,
			10 * 60,
		);

		return {
			uploadFilename,
			uploadUrl,
		};
	}

	async v1GetFileUrl(data: GetFileUrlDto): Promise<string | null> {
		try {
			const object = await this.minio.statObject(
				FileTypeBucket[data.type],
				data.filePath,
			);
			if (!object) {
				return null;
			}

			const url = new URL(
				`${FileTypeBucket[data.type]}/${data.filePath}`,
				this.cdnUrl,
			);

			return url.href;
		} catch (e) {
			console.error(e);
			return null;
		}
	}
}
