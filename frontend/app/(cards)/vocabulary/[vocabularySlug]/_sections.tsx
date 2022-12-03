import CardList from 'app/(cards)/CardList';
import { SectionItem } from 'app/(cards)/Section';
import { Vocabulary } from './entities';

export default function sections(
  vocabulary: Vocabulary,
): { name: string; value: string; items: SectionItem[] }[] {
  return [
    {
      name: 'meaning',
      value: 'Meaning',
      items: [
        {
          kind: 'inline',
          name: 'primary',
          value: vocabulary.name,
        },
        vocabulary.alt_names.length > 0
          ? {
              kind: 'inline',
              name: 'alternative',
              value: vocabulary.alt_names.join(', '),
            }
          : null,
        {
          kind: 'inline',
          name: 'word type',
          value: vocabulary.word_type.join(', '),
        },
        {
          kind: 'markdown',
          name: 'mnemonic',
          value: vocabulary.meaning_mnemonic,
        },
      ],
    },
    {
      name: 'reading',
      value: 'Reading',
      items: [
        { kind: 'inline', name: 'primary', value: vocabulary.reading },
        {
          kind: 'markdown',
          name: 'mnemonic',
          value: vocabulary.reading_mnemonic,
        },
      ],
    },
  ];
}
