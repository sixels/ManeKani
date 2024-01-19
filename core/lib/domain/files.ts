import { Static, Type } from "@sinclair/typebox";

export enum FileType {
	attachments = "attachments",
	avatars = "avatars",
}

export type UploadFileDto = Static<typeof UploadFileSchema>;
export const UploadFileSchema = Type.Object({
	type: Type.Enum(FileType),
	filename: Type.String(),
	size: Type.Integer(),
	metadata: Type.Optional(Type.Record(Type.String(), Type.Any())),
});

export type CreatedFile = Static<typeof CreatedFileSchema>;
export const CreatedFileSchema = Type.Object({
	uploadUrl: Type.String(),
	uploadFilename: Type.String(),
});

export type CreateFileDto = Static<typeof CreateFileSchema>;
export const CreateFileSchema = Type.Object({
	type: Type.Enum(FileType),
	objectName: Type.String(),
	metadata: Type.Optional(Type.Record(Type.String(), Type.Any())),
});

export type GetFileUrlDto = Static<typeof GetFileUrlSchema>;
export const GetFileUrlSchema = Type.Object({
	type: Type.Enum(FileType),
	filePath: Type.String(),
});
