import { FileStorageProviderLabel } from "@/infra/files/files.service";
import { FilesAdapter, IFilesRepositoryV1 } from "@manekani/core";
import { Inject } from "@nestjs/common";

export class FilesService<
	R extends IFilesRepositoryV1,
> extends FilesAdapter<R> {
	constructor(@Inject(FileStorageProviderLabel) filesRepository: R) {
		super(filesRepository);
	}
}
