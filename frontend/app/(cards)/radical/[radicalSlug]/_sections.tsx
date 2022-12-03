import CardList from 'app/(cards)/CardList';
import { SectionItem } from 'app/(cards)/Section';
import { Radical } from './entities';

export default function sections(
  radical: Radical,
): { name: string; value: string; items: SectionItem[] }[] {
  return [
    {
      name: 'meaning',
      value: 'Meaning',
      items: [
        {
          kind: 'inline',
          name: 'primary',
          value: radical.name,
        },
        {
          kind: 'markdown',
          name: 'mnemonic',
          value: radical.meaning_mnemonic,
        },
      ],
    },
    {
      name: 'kanji',
      value: `Found in ${radical.kanji.length} kanji`,
      items: [
        {
          kind: 'raw',
          data: <CardList kind="kanji" cards={radical.kanji} />,
        },
      ],
    },
  ];
}
