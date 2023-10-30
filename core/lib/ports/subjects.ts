import {
  CreateSubjectDto,
  GetSubjectsFilters,
  Subject,
  UpdateSubjectDto,
} from '../domain/subject';

export interface ISubjectRepositoryV1 {
  v1GetSubjects(filters: GetSubjectsFilters): Promise<Subject[]>;
  v1GetSubject(subjectId: string): Promise<Subject | null>;
  v1CreateSubject(
    userId: string,
    deckId: string,
    subject: CreateSubjectDto,
  ): Promise<Subject>;
  v1UpdateSubject(
    ownerId: string,
    subjectId: string,
    changes: UpdateSubjectDto,
  ): Promise<Subject>;
  v1DeleteSubject(userId: string, subjectId: string): Promise<void>;
}
