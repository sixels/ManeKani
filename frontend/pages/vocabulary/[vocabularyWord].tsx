import { GetServerSidePropsContext } from "next";

import { PartialKanjiResponse, Vocabulary } from "@/lib/models/cards";
import { vocabularySections } from "@/ui/Subject/sections/vocabulary";
import { fetchApi } from "@/lib/api/fetchApi";
import CardPage from "@/ui/Subject";
import { PartialSubject, Subject } from "@/lib/models/subject";

interface VocabularyProps {
  vocab: Subject;
  vocabKanjis?: PartialSubject[];
}

export async function getServerSideProps({
  params,
}: GetServerSidePropsContext) {
  if (params == undefined) {
    return { notFound: true };
  }

  const { vocabularyWord } = params;
  const vocabs = await fetchApi<PartialSubject[]>(
    `/subject?kinds=vocabulary&slugs=${vocabularyWord}`,
    "v1"
  );

  if (vocabs == null || vocabs.length == 0) {
    return { notFound: true };
  }
  console.log(
    `found ${vocabs.length} vocabularies with slug ${vocabularyWord}`
  );

  const [vocab, vocabKanjis] = await Promise.all([
    fetchApi<Subject[]>(`/subject/${vocabs[0].id}`, "v1"),
    fetchApi<PartialSubject[]>(
      `/subject?ids=${vocabs[0].dependencies.join(",")}`,
      "v1"
    ),
  ]);

  return { props: { vocab, vocabKanjis } };
}

export default function Page({ vocab, vocabKanjis }: VocabularyProps) {
  const sections = vocabularySections(vocab, vocabKanjis);

  return (
    <CardPage
      card={{
        kind: "vocabulary",
        level: vocab.level,
        name: vocab.name,
        value: vocab.value,
        reading: vocab.study_data
          .find((sd) => sd.kind == "reading")
          ?.items.find((si) => si.is_primary)?.value,
      }}
      sections={sections}
    />
  );
}
