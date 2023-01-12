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
  fetchApi("/kanji", "v1", { method: "POST" });

  function renderKanji() {
    for (let kanj of kanji) {
      return (
        <>
          {kanj.name} - {kanj.symbol}
          <hr />
        </>
      );
    }
  }

  return <>{renderKanji()}</>;
}
