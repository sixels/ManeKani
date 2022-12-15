use manekani_service_common::repository::{RepoInsertable, RepoQueryable, RepoUpdateable};

use crate::model::{
    kanji::ReqKanjiUpdate, Kanji, KanjiPartial, ReqKanjiInsert, ReqKanjiQuery, ReqRadicalQuery,
};

use super::Error;

#[async_trait::async_trait]
pub trait Repository:
    RepoQueryable<ReqKanjiQuery, Kanji>
    + RepoQueryable<ReqRadicalQuery, Vec<KanjiPartial>>
    + RepoInsertable<ReqKanjiInsert, Kanji>
    + RepoUpdateable<ReqKanjiUpdate, Kanji>
{
    async fn query_kanji(&self, kanji: ReqKanjiQuery) -> Result<Kanji, Error> {
        Ok(self.query(kanji).await?)
    }

    async fn query_kanji_by_radical(
        &self,
        radical: ReqRadicalQuery,
    ) -> Result<Vec<KanjiPartial>, Error> {
        Ok(self.query(radical).await?)
    }

    async fn insert_kanji(&self, kanji: ReqKanjiInsert) -> Result<Kanji, Error> {
        Ok(self.insert(kanji).await?)
    }

    async fn update_kanji(&self, kanji: ReqKanjiUpdate) -> Result<Kanji, Error> {
        Ok(self.update(kanji).await?)
    }
}

impl<T> Repository for T where
    T: RepoQueryable<ReqKanjiQuery, Kanji>
        + RepoQueryable<ReqRadicalQuery, Vec<KanjiPartial>>
        + RepoInsertable<ReqKanjiInsert, Kanji>
        + RepoUpdateable<ReqKanjiUpdate, Kanji>
{
}

#[cfg(test)]
mod tests {

    use crate::{
        domain::{KanjiRepository, RadicalRepository},
        model::{radical::barb, ReqKanjiInsert},
        repository::Repository,
    };

    use manekani_service_common::repository::{error::Error, InsertError};
    use sqlx::PgPool;

    #[sqlx::test]
    async fn it_should_create_a_new_kanji(pool: PgPool) -> Result<(), Error> {
        let repo = Repository::new(pool);

        let barb = repo.insert_radical(barb()).await?;
        let kanji = ReqKanjiInsert::builder()
            .name("Finish")
            .alt_names(vec!["Complete".to_owned(), "End".to_owned()])
            .symbol("了")
            .level(2)
            .meaning_mnemonic(
                r#"Think about it this way. There is a famous fishing lure inventor, working on his best work ever. He’s old, and he’s been trying to make the best fishing lure for the last 50 years, and knows this will be his last invention. Then he figures it out. He takes one barb, and connects the end of it to the top of another barb. When he does this, he knows his work is finally finished."#
            )
            .reading("りょう")
            .onyomi(vec!["りょう".to_owned()])
            .reading_mnemonic(
                r#"How does the fish lure maker test his newly finished lure out? The one he just finished? He gets in a row (りょう) boat and rows out into the sea."#
            )
            .radical_composition(vec![barb.name])
            .build();

        let created_kanji = repo.insert_kanji(kanji.clone()).await?;

        assert_eq!(created_kanji.name, kanji.name);
        assert_eq!(created_kanji.alt_names, kanji.alt_names);
        assert_eq!(created_kanji.symbol, kanji.symbol);
        assert_eq!(created_kanji.meaning_mnemonic, kanji.meaning_mnemonic);
        assert_eq!(created_kanji.reading, kanji.reading);
        assert_eq!(created_kanji.onyomi, kanji.onyomi);
        assert_eq!(created_kanji.kunyomi, Vec::<String>::new());
        assert_eq!(created_kanji.nanori, Vec::<String>::new());
        assert_eq!(created_kanji.reading_mnemonic, kanji.reading_mnemonic);

        Ok(())
    }

    #[sqlx::test]
    async fn it_should_collide_with_an_existing_kanji(pool: PgPool) -> Result<(), Error> {
        let repo = Repository::new(pool);

        let barb = repo.insert_radical(barb()).await?;
        let kanji = ReqKanjiInsert::builder()
            .name("finish")
            .alt_names(vec!["Complete".to_owned(), "End".to_owned()])
            .symbol("了")
            .level(2)
            .meaning_mnemonic(
                r#"Think about it this way. There is a famous fishing lure inventor, working on his best work ever. He’s old, and he’s been trying to make the best fishing lure for the last 50 years, and knows this will be his last invention. Then he figures it out. He takes one barb, and connects the end of it to the top of another barb. When he does this, he knows his work is finally finished."#
            )
            .reading("りょう")
            .onyomi(vec!["りょう".to_owned()])
            .reading_mnemonic(
                r#"How does the fish lure maker test his newly finished lure out? The one he just finished? He gets in a row (りょう) boat and rows out into the sea."#
            )
            .radical_composition(vec![barb.name])
            .build();

        repo.insert_kanji(kanji.clone()).await?;
        let kanji = repo.insert_kanji(kanji).await;

        assert!(matches!(kanji, Err(Error::Insert(InsertError::Conflict))));

        Ok(())
    }
}
