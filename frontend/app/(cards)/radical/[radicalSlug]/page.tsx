import { fetchApi } from '@/lib/api/fetchApi';
import RetroFrame from '@/ui/RetroFrame';
import Heading from 'app/(cards)/Heading';

import Section from 'app/(cards)/Section';
import { Radical, RadicalKanji } from './entities';
import sections from './_sections';

const getRadical = async (radicalName: string): Promise<Radical> => {
  const [radical, radicalKanji] = await Promise.all([
    fetchApi<Omit<Radical, 'kanji'>>(`/v1/radical/${radicalName}`),
    fetchApi<RadicalKanji[]>(`/v1/kanji/from-radical/${radicalName}`),
  ]);

  radical.created_at = new Date(radical.created_at);
  radical.updated_at = new Date(radical.updated_at);
  radical.symbol = new Uint8Array(radical.symbol);

  return Object.assign(radical, {
    kanji: radicalKanji,
  });
};

export default async function Page({
  params,
}: {
  params: { radicalSlug: string };
}) {
  const decoder = new TextDecoder('utf8');

  const radicalName = params.radicalSlug;
  const radical = await getRadical(radicalName);

  const icon = (radical.symbol.length > 5 && { data: radical.symbol }) || {
    str: decoder.decode(radical.symbol),
  };

  const radicalSections = sections(radical).map((section, i) => (
    <Section
      kind="radical"
      name={section.name}
      value={section.value}
      items={section.items}
      key={i}
    />
  ));

  return (
    <div className="view-container w-full bg-gray-100">
      <div className="radical-view container mx-auto space-y-12">
        <Heading
          kind="radical"
          item={{ name: radical.name, level: radical.level, icon }}
        />
        {radicalSections}
      </div>
    </div>
  );
}
