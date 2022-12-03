type ComponentProps = {
  primary: string;
  onyomi: string[];
  kunyomi: string[];
  nanori: string[];
};

export default function Readings({
  primary,
  onyomi,
  kunyomi,
  nanori,
}: ComponentProps) {
  const isPrimary = (readings: string[]) => readings.includes(primary);

  return (
    <div className="flex px-6 py-2">
      <Reading title="onyomi" readings={onyomi} isPrimary={isPrimary} />
      <Reading title="kunyomi" readings={kunyomi} isPrimary={isPrimary} />
      <Reading title="nanori" readings={nanori} isPrimary={isPrimary} />
    </div>
  );
}

function Reading({
  title,
  readings,
  isPrimary,
}: {
  title: string;
  readings: string[];
  isPrimary: (readings: string[]) => boolean;
}) {
  const textDim = 'text-gray-400';

  const textColor = (isPrimary(readings) && ' ') || textDim;

  const readingReadings = readings.map((reading, i) => {
    return (
      <span className={`${textColor}`} key={i}>
        {i > 0 && ', '}
        {reading}
      </span>
    );
  });

  return (
    <div className="reading w-1/3 py-3 text-base">
      <span className="reading-title  font-grotesk font-bold uppercase leading-relaxed">
        {title}
      </span>
      <div className="reading-readings mt-2 capitalize">
        {(readingReadings.length > 0 && readingReadings) || (
          <span className={textDim}>None</span>
        )}
      </div>
    </div>
  );
}
