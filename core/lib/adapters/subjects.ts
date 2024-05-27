import {
	CreateSubjectDto,
	CreateSubjectSchema,
	GetSubjectsFilters,
	GetSubjectsFiltersSchema,
	Subject,
	UpdateSubjectSchema,
} from "../domain/subject";

import { FileType } from "../domain";
import { InvalidRequestError, ResourceNotFoundError } from "../domain/error";
import { IFilesRepositoryV1 } from "../ports";
import { ISubjectRepositoryV1 } from "../ports/subjects";
import { Validator } from "../validator";
import { validateId } from "./common";
import { FilesAdapter } from "./files";

export const CreateSubjectValidator = new Validator(CreateSubjectSchema);
export const UpdateSubjectValidador = new Validator(UpdateSubjectSchema);
export const SubjectsFiltersValidator = new Validator(GetSubjectsFiltersSchema);

export class SubjectsAdapter<S extends ISubjectRepositoryV1> {
	private filesAdapter?: FilesAdapter<IFilesRepositoryV1>;

	constructor(private subjectsRepository: S) {}

	withFilesAdapter(filesRepository: IFilesRepositoryV1): this {
		this.filesAdapter = new FilesAdapter(filesRepository);
		return this;
	}

	v1GetSubjects(filters: GetSubjectsFilters): Promise<Subject[]> {
		SubjectsFiltersValidator.validate(filters);

		return this.subjectsRepository.v1GetSubjects(filters);
	}

	async v1GetSubject(subjectId: string): Promise<Subject> {
		validateId(subjectId);

		const foundSubject = await this.subjectsRepository.v1GetSubject(subjectId);
		if (!foundSubject) {
			throw new ResourceNotFoundError({
				cause: new Error("Subject not found"),
				context: { subjectId },
				description: `No subjects with id "${subjectId}" were found.`,
			});
		}
		return foundSubject;
	}

	async v1CreateSubject(
		userId: string,
		deckId: string,
		subject: CreateSubjectDto,
	): Promise<Subject> {
		validateId(deckId);
		CreateSubjectValidator.validate(subject);

		if (!subject.value && !subject.valueImage) {
			throw new InvalidRequestError({
				cause: new Error("Subjects must have a value or a value image."),
				context: { subject },
				description: "A subject must have a value or a value image.",
			});
		}

		// validate subject image
		if (subject.valueImage) {
			const subjectImageUrl = await this.filesAdapter?.v1GetResolvedFilePath({
				type: FileType.attachments,
				filePath: subject.valueImage,
			});

			if (!subjectImageUrl) {
				throw new InvalidRequestError({
					cause: new Error("Subject image not found"),
					context: { subject },
					description: "The provided subject image does not exists.",
				});
			}

			subject.valueImage = subjectImageUrl;
		}

		// TODO: check if user can modify the deck

		console.debug("creating subject:", { userId, subject });
		const createdSubject = await this.subjectsRepository.v1CreateSubject(
			userId,
			deckId,
			subject,
		);

		return createdSubject;
	}

	v1UpdateSubject(
		userId: string,
		subjectId: string,
		changes: Partial<Subject>,
	): Promise<Subject> {
		validateId(subjectId);
		UpdateSubjectValidador.validate(changes);

		// TODO: check if user can modify the subject

		console.debug("updating subject:", { subjectId, changes });
		return this.subjectsRepository.v1UpdateSubject(userId, subjectId, changes);
	}

	v1DeleteSubject(userId: string, subjectId: string): Promise<void> {
		validateId(subjectId);

		// TODO: check if user can modify the subject

		console.debug("deleting subject:", { subjectId });
		return this.subjectsRepository.v1DeleteSubject(userId, `${subjectId}`);
	}
}
