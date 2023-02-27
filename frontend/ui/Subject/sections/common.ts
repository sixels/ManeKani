import { StudyData, StudyItem } from "@/lib/models/subject";

export function studyData(subject: { study_data: StudyData[] }, kind: string) {
  return subject.study_data.find((data) => data.kind == kind);
}

export function studyItem(
  subject: { study_data: StudyData[] },
  kind: string,
  findItem: (item: StudyItem) => boolean
) {
  return studyData(subject, kind)?.items.find((si) =>
    findItem(si) ? si : null
  );
}

export function studyItems(
  subject: { study_data: StudyData[] },
  kind: string,
  findItem: (item: StudyItem) => boolean
) {
  return studyData(subject, kind)?.items.filter((si) =>
    findItem(si) ? si : null
  );
}
