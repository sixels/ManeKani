import { GetServerSidePropsContext } from "next";

import { PartialKanjiResponse, Radical, Vocabulary } from "@/lib/models/cards";
import { radicalSections } from "@/ui/CardPage/Sections/radicalSections";
import { API_URL, fetchApi } from "@/lib/api/fetchApi";
import DefaultLayout from "@/ui/layouts/Default";
import { PageWithLayout } from "@/ui/layouts";
import CardPage from "@/ui/CardPage";

interface RadicalProps {
  radical: Radical;
  radicalKanjis?: PartialKanjiResponse[];
}

export async function getServerSideProps({
  params,
}: GetServerSidePropsContext) {
  if (params == undefined) {
    return { notFound: true };
  }

  const { radicalName } = params;
  const radical = await fetchApi<Radical>(`/radical/${radicalName}`, "v1");

  if (radical == null) {
    return { notFound: true };
  }

  const radicalKanjis = await fetchApi<PartialKanjiResponse[]>(
    `/radical/${radicalName}/kanji`,
    "v1"
  );

  return { props: { radical, radicalKanjis } };
}

const PageLayout = DefaultLayout;
const Page: PageWithLayout<RadicalProps> = ({ radical, radicalKanjis }) => {
  const sections = radicalSections(radical, radicalKanjis);

  return (
    <CardPage
      card={{
        kind: "radical",
        level: radical.level,
        meaning: radical.name,
        value: radical.symbol.includes("/")
          ? { image: `${API_URL}/files/${radical.symbol}` }
          : radical.symbol,
      }}
      sections={sections}
    />
  );
};

Page.getLayout = (page) => {
  return <PageLayout>{page}</PageLayout>;
};

export default Page;
