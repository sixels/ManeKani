import { GetServerSidePropsContext } from "next";

import { fetchApi } from "@/lib/api/fetchApi";

import { kanjiSections } from "@/ui/Subject/sections/kanji";
import CardPage from "@/ui/Subject";
import { PartialSubject, Subject } from "@/lib/models/subject";

interface KanjiProps {
  kanji: Subject;
  kanjiRadicals?: PartialSubject[];
  kanjiVocabularies?: PartialSubject[];
}

export async function getServerSideProps({
  params,
}: GetServerSidePropsContext) {
  if (params == undefined) {
    return { notFound: true };
  }

  const { kanjiSymbol } = params;
  const kanjis = await fetchApi<PartialSubject[]>(
    `/subject?kinds=kanji&slugs=${kanjiSymbol}`,
    "v1"
  );

  if (kanjis == null || kanjis.length == 0) {
    return { notFound: true };
  }
  console.log(`found ${kanjis.length} kanji with slug ${kanjiSymbol}`);

  const [kanji, kanjiRadicals, kanjiVocabularies] = await Promise.all([
    fetchApi<Subject[]>(`/subject/${kanjis[0].id}`, "v1"),
    fetchApi<PartialSubject[]>(
      `/subject?ids=${kanjis[0].dependencies.join(",")}`,
      "v1"
    ),
    fetchApi<PartialSubject[]>(
      `/subject?ids=${kanjis[0].dependents.join(",")}`,
      "v1"
    ),
  ]);

  return { props: { kanji, kanjiRadicals, kanjiVocabularies } };
}

export default function Page({
  kanji,
  kanjiRadicals,
  kanjiVocabularies,
}: KanjiProps) {
  const sections = kanjiSections(kanji, kanjiVocabularies, kanjiRadicals);

  return (
    <CardPage
      card={{
        kind: "kanji",
        level: kanji.level,
        name: kanji.name,
        value: kanji.value,
        reading: kanji.study_data
          .find((sd) => sd.kind == "reading")
          ?.items.find((si) => si.is_primary)?.value,
      }}
      sections={sections}
    />
  );
}
