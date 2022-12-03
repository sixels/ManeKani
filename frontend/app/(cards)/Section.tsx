import RetroFrame from '@/ui/RetroFrame';
import React from 'react';

export type SectionItem =
  | { kind: 'raw'; data: React.ReactNode }
  | { kind: 'inline'; name: string; value: string }
  | { kind: 'markdown'; name: string; value: string }
  | null;

type SectionProps = {
  kind: 'kanji' | 'radical' | 'vocabulary';
  name: string;
  value: string;
  items: SectionItem[];
};

const CARD_BG_ACCENT = Object.freeze({
  kanji: 'after:bg-card-kanji',
  radical: 'after:bg-card-radical',
  vocabulary: 'after:bg-card-vocabulary',
});

const Item = (item: SectionItem): React.ReactNode => {
  if (item?.kind == 'raw') {
    return item.data;
  }
  if (item?.kind == 'inline') {
    return (
      <div className="item flex gap-4 uppercase">
        <span className="item-key font-grotesk font-semibold uppercase">
          {item.name}
        </span>
        <span className="item-value capitalize"> {item.value} </span>
      </div>
    );
  }
  if (item?.kind == 'markdown') {
    return (
      <div className="item mt-2">
        <span className="item-title font-grotesk font-semibold uppercase">
          {item.name}
        </span>
        <div
          className="item-value relative mt-2 font-sans leading-relaxed"
          dangerouslySetInnerHTML={{ __html: item.value }}
        ></div>
      </div>
    );
  }

  return <></>;
};

export default function Section({ name, value, kind, items }: SectionProps) {
  const accent = CARD_BG_ACCENT[kind];

  const sectionItems = items.map(Item);

  return (
    <section
      id={name}
      className="meaning-container relative w-full rounded-2xl py-12 shadow-[6px_6px_0] shadow-black ring-[3px] ring-inset ring-black"
    >
      <div className="meaning-title absolute -top-[15px] left-20 w-max">
        <RetroFrame backColor={accent}>
          <span className="h-full w-full bg-gray-100 py-1 px-3 font-grotesk text-lg capitalize">
            {value}
          </span>
        </RetroFrame>
      </div>
      <div className="kanji-meaning px-6">
        <div className="items grid gap-4 text-sm">{sectionItems}</div>
      </div>
    </section>
  );
}
