import { GetServerSidePropsContext } from "next";

import { PartialKanjiResponse, Vocabulary } from "@/lib/models/cards";
import { vocabularySections } from "@/ui/Sections/vocabularySections";
import { fetchApi } from "@/lib/api/fetchApi";
import DefaultLayout from "@/ui/layouts/Default";
import { PageWithLayout } from "@/ui/layouts";
import CardPage from "@/ui/CardPage";

interface VocabularyProps {
  vocab: Vocabulary;
  vocabKanjis?: PartialKanjiResponse[];
}

export async function getServerSideProps({
  params,
}: GetServerSidePropsContext) {
  if (params == undefined) {
    return { notFound: true };
  }

  const { vocabularyWord } = params;
  const vocab = await fetchApi<Vocabulary>(
    `/vocabulary/${vocabularyWord}`,
    "v1"
  );

  if (vocab == null) {
    return { notFound: true };
  }

  const vocabKanjis = await fetchApi<PartialKanjiResponse[]>(
    `/vocabulary/${vocabularyWord}/kanji`,
    "v1"
  );

  return { props: { vocab, vocabKanjis } };
}

const PageLayout = DefaultLayout;
const Page: PageWithLayout<VocabularyProps> = ({ vocab, vocabKanjis }) => {
  const sections = vocabularySections(vocab, vocabKanjis);

  return (
    <CardPage
      card={{
        kind: "vocabulary",
        level: vocab.level,
        meaning: vocab.name,
        value: vocab.word,
        reading: vocab.reading,
      }}
      sections={sections}
    />
  );
};

Page.getLayout = (page) => {
  return <PageLayout>{page}</PageLayout>;
};

export default Page;
