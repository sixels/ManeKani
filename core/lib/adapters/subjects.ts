import {
	CreateSubjectDto,
	CreateSubjectSchema,
	GetSubjectsFilters,
	GetSubjectsFiltersSchema,
	Subject,
	UpdateSubjectSchema,
} from "../domain/subject";

import { ResourceNotFoundError } from "../domain/error";
import { ISubjectRepositoryV1 } from "../ports/subjects";
import { Validator } from "../validator";
import { validateId } from "./common";

export const CreateSubjectValidator = new Validator(CreateSubjectSchema);
export const UpdateSubjectValidador = new Validator(UpdateSubjectSchema);
export const SubjectsFiltersValidator = new Validator(GetSubjectsFiltersSchema);

export class SubjectsAdapter<R extends ISubjectRepositoryV1> {
	constructor(private subjectsRepository: R) {}

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

	v1CreateSubject(
		userId: string,
		deckId: string,
		subject: CreateSubjectDto,
	): Promise<Subject> {
		validateId(deckId);
		CreateSubjectValidator.validate(subject);

		// TODO: check if user can modify the deck

		console.debug("creating subject:", { userId, subject });
		return this.subjectsRepository.v1CreateSubject(userId, deckId, subject);
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
