import { fetchApi } from '@/lib/api/fetchApi';

import { Kanji, KanjiRadical, KanjiVocabulary } from './entities';
import sections from './_sections';
import Section from 'app/(cards)/Section';
import markdown from '@/lib/markdown';
import Heading from 'app/(cards)/Heading';

const getKanji = async (kanjiSymbol: string): Promise<Kanji> => {
  const [kanji, kanjiVocabularies, kanjiRadicals] = await Promise.all([
    fetchApi<Omit<Kanji, 'vocabularies' | 'radicals'>>(
      `/v1/kanji/${kanjiSymbol}`,
    ),
    fetchApi<KanjiVocabulary[]>(`/v1/vocabulary/from-kanji/${kanjiSymbol}`),
    fetchApi<KanjiRadical[]>(`/v1/radical/from-kanji/${kanjiSymbol}`),
  ]);

  kanji.created_at = new Date(kanji.created_at);
  kanji.updated_at = new Date(kanji.updated_at);
  [kanji.meaning_mnemonic, kanji.reading_mnemonic] = await Promise.all(
    [kanji.meaning_mnemonic, kanji.reading_mnemonic].map(markdown.toHTML),
  );

  kanjiRadicals.forEach((radical) => {
    radical.symbol = new Uint8Array(radical.symbol);
  });

  return Object.assign(kanji, {
    vocabularies: kanjiVocabularies,
    radicals: kanjiRadicals,
  });
};

export default async function Page({
  params,
}: {
  params: { kanjiSlug: string };
}) {
  const kanjiSymbol = params.kanjiSlug;
  const kanji = await getKanji(kanjiSymbol);

  const kanjiSections = sections(kanji).map((section, i) => (
    <Section
      kind="kanji"
      name={section.name}
      value={section.value}
      items={section.items}
      key={i}
    />
  ));

  return (
    <div className="view-container w-full bg-gray-100">
      <div className="kanji-view container mx-auto space-y-12">
        <Heading
          kind="kanji"
          item={{
            name: kanji.name,
            level: kanji.level,
            icon: { str: kanji.symbol },
          }}
        />
        {kanjiSections}
      </div>
    </div>
  );
}
