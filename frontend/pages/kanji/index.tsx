import { fetchApi } from "@/lib/api/fetchApi";
import { isApiError } from "@/lib/api/utils";
import { PartialKanjiResponse } from "@/lib/models/cards";
import { GetServerSidePropsContext } from "next";

interface IndexProps {
  kanji: PartialKanjiResponse[];
}

export async function getServerSideProps(ctx: GetServerSidePropsContext) {
  const kanji = await fetchApi<PartialKanjiResponse[]>(`/kanji`, "v1");
  if (kanji == null) {
    return { notFound: true };
  }

  return { props: { kanji } };
}

export default function Index({ kanji }: IndexProps) {
  fetchApi("/kanji", "v1");

  return (
    <>
      {kanji.map((k) => (
        <>
          {k.name} - {k.symbol}
          <hr />
        </>
      ))}
    </>
  );
}
