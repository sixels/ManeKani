import { CreateFileDto, CreatedFile, GetFileUrlDto } from "../domain/files";

export interface IFilesRepositoryV1 {
	v1CreateFile(data: CreateFileDto): Promise<CreatedFile>;
	v1GetResolvedFilePath(data: GetFileUrlDto): Promise<string | null>;
}
