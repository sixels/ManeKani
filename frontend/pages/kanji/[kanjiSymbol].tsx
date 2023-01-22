import { GetServerSidePropsContext } from "next";

import { fetchApi } from "@/lib/api/fetchApi";
import {
  Kanji,
  PartialRadicalResponse,
  PartialVocabularyResponse,
} from "@/lib/models/cards";
import { kanjiSections } from "@/ui/Sections/kanjiSections";
import { PageWithLayout } from "@/ui/layouts";
import DefaultLayout from "@/ui/layouts/Default";
import CardPage from "@/ui/CardPage";

interface KanjiProps {
  kanji: Kanji;
  kanjiRadicals?: PartialRadicalResponse[];
  kanjiVocabularies?: PartialVocabularyResponse[];
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

const PageLayout = DefaultLayout;
const Page: PageWithLayout<KanjiProps> = ({
  kanji,
  kanjiRadicals,
  kanjiVocabularies,
}) => {
  const sections = kanjiSections(kanji, kanjiVocabularies, kanjiRadicals);

  return (
    <CardPage
      card={{
        kind: "kanji",
        level: kanji.level,
        meaning: kanji.name,
        value: kanji.symbol,
        reading: kanji.reading,
      }}
      sections={sections}
    />
  );
};

Page.getLayout = (page) => {
  return <PageLayout>{page}</PageLayout>;
};

export default Page;
