import RetroFrame from '@/ui/RetroFrame';
import { CARD_BG } from './cardColors';

type ComponentProps = {
  kind: 'kanji' | 'radical' | 'vocabulary';
  cards: {
    name: string;
    symbol?: string;
    word?: string;
    reading?: string;
  }[];
};

const Card = ({
  kind,
  card,
}: {
  kind: ComponentProps['kind'];
  card: ComponentProps['cards'][number];
}) => {
  const background = CARD_BG[kind];
  const href = `/${kind}/${card.symbol || card.word}`;

  return (
    <RetroFrame backColor="after:bg-gray-100">
      <a
        className={`card flex h-16 items-center justify-between ${background} bg-gradient-to-b px-2 text-gray-100  `}
        href={href}
      >
        <span className="symbol text-4xl ">
          {(card.symbol && card.symbol) || card.word}
        </span>
        <ul className="info items-end text-right text-sm">
          {card.reading && <li className="reading"> {card.reading} </li>}
          <li className="name text-xs capitalize"> {card.name} </li>
        </ul>
      </a>
    </RetroFrame>
  );
};

export default function CardList({ cards, kind }: ComponentProps) {
  const cardList = cards.map((card, i) => (
    <Card card={card} kind={kind} key={i} />
  ));

  return (
    (cards.length > 0 && (
      <div className={`card-list grid grid-cols-1 gap-5`}>{cardList}</div>
    )) || <></>
  );
}
