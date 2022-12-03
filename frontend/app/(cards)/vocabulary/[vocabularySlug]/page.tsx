import { fetchApi } from '@/lib/api/fetchApi';
import markdown from '@/lib/markdown';
import RetroFrame from '@/ui/RetroFrame';

import CardList from 'app/(cards)/CardList';
import Heading from 'app/(cards)/Heading';
import Section from 'app/(cards)/Section';
import { Vocabulary } from './entities';
import sections from './_sections';

const symbolB64 = (chars: Uint8Array) => {
  return Buffer.from(chars).toString('base64');
};

const getVocabulary = async (vocabularyName: string): Promise<Vocabulary> => {
  const vocabulary = await fetchApi<Vocabulary>(
    `/v1/vocabulary/${vocabularyName}`,
  );

  vocabulary.created_at = new Date(vocabulary.created_at);
  vocabulary.updated_at = new Date(vocabulary.updated_at);
  [vocabulary.meaning_mnemonic, vocabulary.reading_mnemonic] =
    await Promise.all(
      [vocabulary.meaning_mnemonic, vocabulary.reading_mnemonic].map(
        markdown.toHTML,
      ),
    );

  return vocabulary;
};

export default async function Page({
  params,
}: {
  params: { vocabularySlug: string };
}) {
  const vocabularyName = params.vocabularySlug;
  const vocabulary = await getVocabulary(vocabularyName);

  const vocabularySections = sections(vocabulary).map((section, i) => (
    <Section
      kind="vocabulary"
      name={section.name}
      value={section.value}
      items={section.items}
      key={i}
    />
  ));

  return (
    <div className="view-container w-full bg-gray-100">
      <div className="vocabulary-view container mx-auto space-y-12">
        <Heading
          kind="vocabulary"
          item={{
            name: vocabulary.name,
            level: vocabulary.level,
            icon: { str: vocabulary.word },
          }}
        />
        {vocabularySections}
      </div>
    </div>
  );
}
