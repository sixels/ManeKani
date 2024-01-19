import { SubjectsProviderLabel } from "@/infra/database/subjectsDatabase.service";
import { ISubjectRepositoryV1, SubjectsAdapter } from "@manekani/core";
import { Inject, Injectable } from "@nestjs/common";

@Injectable()
export class SubjectsService<
	R extends ISubjectRepositoryV1,
> extends SubjectsAdapter<R> {
	constructor(@Inject(SubjectsProviderLabel) subjectsRepository: R) {
		super(subjectsRepository);
	}
}
