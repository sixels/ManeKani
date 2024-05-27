import {
	CreateFileDto,
	CreatedFile,
	FileType,
	GetFileUrlDto,
} from "@manekani/core";
import {HttpProxyAgent} from "http-proxy-agent"
import { IFilesRepositoryV1 } from "@manekani/core";
import { Client } from "minio";

const FileTypeBucket: Record<FileType, string> = {
	[FileType.attachments]: "attachments",
	[FileType.avatars]: "avatars",
};

type RepoConfig = {
	proxyUrl: string;

	s3Endpoint: string;
	s3Port: number;
	s3AccessKey: string;
	s3SecretKey: string;
};

export class FileStorage implements IFilesRepositoryV1 {
	private minio: Client;
	// private config: RepoConfig;

	constructor(config: RepoConfig) {
		this.minio = new Client({
			endPoint: config.s3Endpoint || "",
			port: config.s3Port,
			accessKey: config.s3AccessKey || "",
			secretKey: config.s3SecretKey || "",
			useSSL: process.env.NODE_ENV === "production",
			transportAgent: new HttpProxyAgent(config.proxyUrl),
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
		let uploadUrl = await this.minio.presignedUrl(
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

	async v1GetResolvedFilePath(data: GetFileUrlDto): Promise<string | null> {
		try {
			const object = await this.minio.statObject(
				FileTypeBucket[data.type],
				data.filePath,
			);
			if (!object) {
				return null;
			}

			return `${FileTypeBucket[data.type]}/${data.filePath}`
		} catch (e) {
			console.error(e);
			return null;
		}
	}
}
