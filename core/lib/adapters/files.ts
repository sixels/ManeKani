import { randomUUID } from "crypto";
import base58 from "bs58";
import { InvalidRequestError } from "../domain";
import { CreatedFile, GetFileUrlDto, UploadFileDto } from "../domain/files";
import { IFilesRepositoryV1 } from "../ports/files";

export class FilesAdapter<R extends IFilesRepositoryV1> {
	constructor(private filesRepository: R) {}

	v1UploadFile(userId: string, data: UploadFileDto): Promise<CreatedFile> {
		if (data.size > 10000000) {
			throw new InvalidRequestError({
				description: "File size too big",
				context: { size: data.size },
			});
		}

		const encodedNamespace = encodeB58Uuid(userId);

		const fileId = randomUUID();
		const encodedFileId = encodeB58Uuid(fileId);

		return this.filesRepository.v1CreateFile({
			type: data.type,
			objectName: `${encodedNamespace}/${encodedFileId}`,
			metadata: data.metadata,
		});
	}

	v1GetFileUrl(data: GetFileUrlDto): Promise<string | null> {
		return this.filesRepository.v1GetFileUrl(data);
	}
}

function encodeB58Uuid(uuid: string): string {
	return base58.encode(Buffer.from(uuid.replaceAll("-", ""), "hex"));
}
