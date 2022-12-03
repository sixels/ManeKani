import RetroFrame from '@/ui/RetroFrame';

import { CardKind, CARD_BG } from './cardColors';

const bytesToB64 = (bytes: Uint8Array) => {
  return Buffer.from(bytes).toString('base64');
};

type ComponentProps = {
  kind: CardKind;
  item: {
    name: string;
    level: number;
    icon: { data: Uint8Array } | { str: string };
  };
};
export default function Heading({ kind, item }: ComponentProps) {
  const headingIcon =
    'data' in item.icon ? (
      <div className="h-12 w-12" style={{ imageRendering: 'crisp-edges' }}>
        <img src={`data:image/png;base64, ${bytesToB64(item.icon.data)}`} />
      </div>
    ) : (
      <span lang="ja" className="text-5xl text-white ">
        {item.icon.str}
      </span>
    );

  return (
    <>
      <section className="heading-container flex w-full items-center py-14 px-4">
        <div className="heading-wrapper flex"></div>
        <div className="heading h-20 w-max min-w-[80px]">
          <RetroFrame backColor="after:bg-gray-100">
            <div
              className={`grid h-full w-full place-items-center ${CARD_BG[kind]} px-4`}
            >
              {headingIcon}
            </div>
          </RetroFrame>
        </div>
        <div className="info ml-5 flex flex-col justify-center">
          <div className="item-name radical-name font-grotesk text-4xl font-bold capitalize">
            <span> {item.name}</span>
          </div>
          <div className="item-level radical-level font-grotesk font-bold uppercase text-gray-800">
            <span> Level {item.level}</span>
          </div>
        </div>
      </section>
    </>
  );
}
