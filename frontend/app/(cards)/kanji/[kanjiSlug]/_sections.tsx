import CardList from 'app/(cards)/CardList';
import { SectionItem } from 'app/(cards)/Section';
import Composition from './Composition';
import { Kanji } from './entities';
import Readings from './Readings';

export default function sections(
  kanji: Kanji,
): { name: string; value: string; items: SectionItem[] }[] {
  return [
    {
      name: 'composition',
      value: 'Radical Composition',
      items: [
        {
          kind: 'raw',
          data: (
            <Composition
              radicals={kanji.radicals.map((radical) =>
                Object.assign({
                  name: radical.name,
                  icon:
                    radical.symbol.length > 5
                      ? { kind: 'image', data: radical.symbol }
                      : {
                          kind: 'text',
                          data: new TextDecoder('utf8')
                            .decode(radical.symbol)
                            .toString(),
                        },
                }),
              )}
            />
          ),
        },
      ],
    },
    {
      name: 'meaning',
      value: 'Meaning',
      items: [
        {
          kind: 'inline',
          name: 'primary',
          value: kanji.name,
        },
        kanji.alt_names.length > 0
          ? {
              kind: 'inline',
              name: 'alternative',
              value: kanji.alt_names.join(', '),
            }
          : null,
        {
          kind: 'markdown',
          name: 'mnemonic',
          value: kanji.meaning_mnemonic,
        },
      ],
    },
    {
      name: 'reading',
      value: 'Reading',
      items: [
        {
          kind: 'raw',
          data: (
            <Readings
              primary={kanji.reading}
              onyomi={kanji.onyomi}
              kunyomi={kanji.kunyomi}
              nanori={kanji.nanori}
            />
          ),
        },
        {
          kind: 'markdown',
          name: 'mnemonic',
          value: kanji.reading_mnemonic,
        },
      ],
    },
    {
      name: 'vocabularies',
      value: `Found in ${kanji.vocabularies.length} vocabularies`,
      items: [
        {
          kind: 'raw',
          data: <CardList kind="vocabulary" cards={kanji.vocabularies} />,
        },
      ],
    },
  ];
}
