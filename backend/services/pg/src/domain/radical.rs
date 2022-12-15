use manekani_service_common::repository::{RepoInsertable, RepoQueryable, RepoUpdateable};

use crate::entity::{
    Radical, RadicalPartial, ReqKanjiQuery, ReqRadicalInsert, ReqRadicalQuery, ReqRadicalUpdate,
};

use super::Error;

#[async_trait::async_trait]
pub trait Repository:
    RepoQueryable<ReqRadicalQuery, Radical>
    + RepoQueryable<ReqKanjiQuery, Vec<RadicalPartial>>
    + RepoInsertable<ReqRadicalInsert, Radical>
    + RepoUpdateable<ReqRadicalUpdate, Radical>
{
    async fn query_radical(&self, radical: ReqRadicalQuery) -> Result<Radical, Error> {
        Ok(self.query(radical).await?)
    }

    async fn query_radical_by_kanji(
        &self,
        kanji: ReqKanjiQuery,
    ) -> Result<Vec<RadicalPartial>, Error> {
        Ok(self.query(kanji).await?)
    }

    async fn insert_radical(&self, radical: ReqRadicalInsert) -> Result<Radical, Error> {
        Ok(self.insert(radical).await?)
    }

    async fn update_radical(&self, radical: ReqRadicalUpdate) -> Result<Radical, Error> {
        Ok(self.update(radical).await?)
    }
}

impl<T> Repository for T where
    T: RepoQueryable<ReqRadicalQuery, Radical>
        + RepoQueryable<ReqKanjiQuery, Vec<RadicalPartial>>
        + RepoInsertable<ReqRadicalInsert, Radical>
        + RepoUpdateable<ReqRadicalUpdate, Radical>
{
}

#[cfg(test)]
mod tests {
    use crate::{domain::RadicalRepository, entity::ReqRadicalInsert, repository::Repository};

    // use super::*;

    use manekani_service_common::repository::{error::Error, InsertError};
    use sqlx::PgPool;

    #[sqlx::test]
    async fn it_should_create_a_new_radical(pool: PgPool) -> Result<(), Error> {
        let repo = Repository::new(pool);

        let radical = ReqRadicalInsert::builder()
            .name("barb")
            .level(1)
            .symbol("亅")
            .meaning_mnemonic(
                r#"This radical is shaped like a barb. Like the kind you'd see on barb wire. Imagine one of these getting stuck to your arm or your clothes. Think about how much it would hurt with that little hook on the end sticking into you. Say out loud, "Oh dang, I got a barb stuck in me!""#
            )
            .build();

        let created_radical = repo.insert_radical(radical.clone()).await?;

        assert_eq!(created_radical.name, radical.name);
        assert_eq!(created_radical.symbol, radical.symbol);
        assert_eq!(created_radical.meaning_mnemonic, radical.meaning_mnemonic);
        assert_eq!(created_radical.user_synonyms, None);
        assert_eq!(created_radical.user_meaning_note, None);

        Ok(())
    }

    #[sqlx::test]
    async fn it_should_collide_with_an_existing_radical(pool: PgPool) -> Result<(), Error> {
        let repo = Repository::new(pool);

        let radical = ReqRadicalInsert::builder()
            .name("barb")
            .level(1)
            .symbol("亅")
            .meaning_mnemonic(
                r#"This radical is shaped like a barb. Like the kind you'd see on barb wire. Imagine one of these getting stuck to your arm or your clothes. Think about how much it would hurt with .to_owned()that little hook on the end sticking into you. Say out loud, "Oh dang, I got a barb stuck in me!""#
            )
            .build();

        repo.insert_radical(radical.clone()).await?;
        let radical = repo.insert_radical(radical).await;

        assert!(matches!(radical, Err(Error::Insert(InsertError::Conflict))));

        Ok(())
    }
}
