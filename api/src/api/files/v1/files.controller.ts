import { Authorize, UserData } from "@/api/auth/auth.decorator";
import { FileType, UploadFileDto } from "@manekani/core";
import { Body, Controller, Post } from "@nestjs/common";
import { FilesService } from "./files.service";

type UploadInfo = {
	files: {
		id: number;
		filename: string;
		size: number;
	}[];
};

@Controller()
export class FilesController {
	// biome-ignore lint/suspicious/noExplicitAny: any is just a placeholder type
	constructor(private readonly filesService: FilesService<any>) {}

	@Post("/v1/attachments")
	@Authorize({ loginOnly: false })
	async uploadAttachment(
		@UserData("userId") userId: string,
		@Body() data: UploadInfo,
	) {
		const uploadPromises = data.files.map(async (file) => {
			const uploadData: UploadFileDto = {
				filename: file.filename,
				size: file.size,
				type: FileType.attachments,
			};
			return this.filesService
				.v1UploadFile(userId, uploadData)
				.then((result) => Object.assign(result, { id: file.id }));
		});

		const uploadFiles = await Promise.all(uploadPromises);
		console.log(uploadFiles);

		return {
			files: uploadFiles,
		};
	}

	@Post("/v1/avatars")
	@Authorize()
	async uploadAvatars(
		@UserData("userId") userId: string,
		@Body() data: UploadInfo,
	) {
		const uploadPromises = data.files.map((file) => {
			const uploadData: UploadFileDto = {
				filename: file.filename,
				size: file.size,
				type: FileType.avatars,
			};
			return this.filesService
				.v1UploadFile(userId, uploadData)
				.then((result) => Object.assign(result, { id: file.id }));
		});

		const uploadFiles = await Promise.all(uploadPromises);

		return {
			files: uploadFiles,
		};
	}
}
