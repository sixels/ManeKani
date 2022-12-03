import RetroFrame from '@/ui/RetroFrame';

type ComponentProps = {
  radicals: {
    name: string;
    icon: { kind: 'image'; data: Uint8Array } | { kind: 'text'; data: string };
  }[];
};

export default function Composition({ radicals }: ComponentProps) {
  const radicalComposition = radicals.map((radical, i) => {
    return (
      <>
        <a href={`/radical/${radical.name}`} className="flex items-center">
          <div className="h-16  min-w-[80px] " key={i}>
            <RetroFrame backColor="after:bg-gray-100">
              <div
                className={`grid h-full grid-flow-col place-items-center gap-4 whitespace-nowrap bg-card-radical px-4 text-white`}
              >
                {radical.icon.kind == 'image' ? (
                  <div
                    className="h-12 w-12"
                    style={{ imageRendering: 'crisp-edges' }}
                  >
                    <img
                      src={`data:image/png;base64, ${bytesToB64(
                        radical.icon.data,
                      )}`}
                    />
                  </div>
                ) : (
                  <span lang="ja" className="text-5xl">
                    {radical.icon.data}
                  </span>
                )}
                <span className="font-grotesk text-2xl capitalize">
                  {radical.name}
                </span>
              </div>
            </RetroFrame>
          </div>
        </a>
        {i < radicals.length - 1 && (
          <span className="my-auto ml-2 font-grotesk text-3xl font-semibold ">
            +
          </span>
        )}
      </>
    );
  });

  return <div className="flex gap-6">{radicalComposition}</div>;
}

const bytesToB64 = (bytes: Uint8Array) => {
  return Buffer.from(bytes).toString('base64');
};
