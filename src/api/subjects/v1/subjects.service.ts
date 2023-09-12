import { SubjectsAdapter } from '@/core/adapters/subjects';
import { ISubjectRepositoryV1 } from '@/core/ports/subjects';
import { Inject, Injectable } from '@nestjs/common';

@Injectable()
export class SubjectsService<
  R extends ISubjectRepositoryV1,
> extends SubjectsAdapter<R> {
  constructor(@Inject('SUBJECTS_REPOSITORY') subjectsRepository: R) {
    super(subjectsRepository);
  }
}
