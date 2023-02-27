import { GetServerSidePropsContext } from "next";

import { PartialKanjiResponse, Radical, Vocabulary } from "@/lib/models/cards";
import { radicalSections } from "@/ui/Subject/sections/radical";
import { API_URL, fetchApi } from "@/lib/api/fetchApi";
import CardPage from "@/ui/Subject";
import { PartialSubject, RemoteContent, Subject } from "@/lib/models/subject";

interface RadicalProps {
  radical: Subject;
  radicalKanjis?: PartialSubject[];
}

export async function getServerSideProps({
  params,
}: GetServerSidePropsContext) {
  if (params == undefined) {
    return { notFound: true };
  }

  const { radicalName } = params;
  const radicals = await fetchApi<PartialSubject[]>(
    `/subject?kinds=radical&slugs=${radicalName}`,
    "v1"
  );

  if (radicals == null || radicals.length == 0) {
    return { notFound: true };
  }
  console.log(`found ${radicals.length} radicals with slug ${radicalName}`);

  const [radical, radicalKanjis] = await Promise.all([
    fetchApi<Subject[]>(`/subject/${radicals[0].id}`, "v1"),
    fetchApi<PartialSubject[]>(
      `/subject?ids=${radicals[0].dependents.join(",")}`,
      "v1"
    ),
  ]);

  return { props: { radical, radicalKanjis } };
}

export default function Page({ radical, radicalKanjis }: RadicalProps) {
  const sections = radicalSections(radical, radicalKanjis);

  return (
    <CardPage
      card={{
        kind: "radical",
        level: radical.level,
        name: radical.name,
        value: radical.value,
        value_image: radical.value_image,
      }}
      sections={sections}
    />
  );
}
