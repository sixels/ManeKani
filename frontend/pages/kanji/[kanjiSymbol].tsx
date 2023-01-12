import { GetServerSidePropsContext } from "next";
import { fetchApi } from "@/lib/api/fetchApi";

import {
  Kanji,
  PartialRadicalResponse,
  PartialVocabularyResponse,
} from "@/lib/models/cards";

interface KanjiProps {
  kanji?: Kanji;
  kanjiRadicals?: PartialRadicalResponse;
  kanjiVocabularies?: PartialVocabularyResponse;
}

export async function getServerSideProps({
  params,
}: GetServerSidePropsContext) {
  if (params == undefined) {
    return { notFound: true };
  }

  const { kanjiSymbol } = params;
  const kanji = await fetchApi<Kanji>(`/kanji/${kanjiSymbol}`, "v1");
  if (kanji == null) {
    return { notFound: true };
  }

  const [kanjiRadicals, kanjiVocabularies] = await Promise.all([
    fetchApi<PartialRadicalResponse[]>(`/kanji/${kanjiSymbol}/radicals`, "v1"),
    fetchApi<PartialVocabularyResponse[]>(
      `/kanji/${kanjiSymbol}/vocabularies`,
      "v1"
    ),
  ]);

  return { props: { kanji, kanjiRadicals, kanjiVocabularies } };
}

export default function QueryKanji({
  kanji,
  kanjiRadicals,
  kanjiVocabularies,
}: KanjiProps) {
  return (
    <>
      {kanji}
      <br />
      {kanjiRadicals}
      <br />
      {kanjiVocabularies}
    </>
  );
}
