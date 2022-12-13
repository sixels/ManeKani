use manekani_service_common::repository::{RepoInsertable, RepoQueryable, RepoUpdateable};

use crate::entity::{
    kanji::GetKanji,
    radical::{GetRadical, InsertRadical, Radical, RadicalPartial, UpdateRadical},
};

use super::Error;

#[async_trait::async_trait]
pub trait RadicalRepository:
    RepoQueryable<GetRadical, Radical>
    + RepoQueryable<GetKanji, Vec<RadicalPartial>>
    + RepoInsertable<InsertRadical, Radical>
    + RepoUpdateable<UpdateRadical, Radical>
{
    async fn query_radical(&self, radical: GetRadical) -> Result<Radical, Error> {
        Ok(self.query(radical).await?)
    }

    async fn query_radical_by_kanji(&self, kanji: GetKanji) -> Result<Vec<RadicalPartial>, Error> {
        Ok(self.query(kanji).await?)
    }

    async fn insert_radical(&self, radical: InsertRadical) -> Result<Radical, Error> {
        Ok(self.insert(radical).await?)
    }

    async fn update_radical(&self, radical: UpdateRadical) -> Result<Radical, Error> {
        Ok(self.update(radical).await?)
    }
}

impl<T> RadicalRepository for T where
    T: RepoQueryable<GetRadical, Radical>
        + RepoQueryable<GetKanji, Vec<RadicalPartial>>
        + RepoInsertable<InsertRadical, Radical>
        + RepoUpdateable<UpdateRadical, Radical>
{
}

#[cfg(test)]
mod tests {
    use crate::repository::Repository;

    use super::*;

    use manekani_service_common::repository::InsertError;
    use sqlx::PgPool;

    #[sqlx::test]
    async fn it_should_create_a_new_radical(pool: PgPool) -> Result<(), Error> {
        let repo = Repository::new(pool);

        let radical = InsertRadical::builder()
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

        let radical = InsertRadical::builder()
            .name("barb")
            .level(1)
            .symbol("亅")
            .meaning_mnemonic(
                r#"This radical is shaped like a barb. Like the kind you'd see on barb wire. Imagine one of these getting stuck to your arm or your clothes. Think about how much it would hurt with .to_owned()that little hook on the end sticking into you. Say out loud, "Oh dang, I got a barb stuck in me!""#
            )
            .build();

        let _ = repo.insert_radical(radical.clone()).await?;
        let radical = repo.insert_radical(radical).await;

        assert!(matches!(radical, Err(Error::Insert(InsertError::Conflict))));

        Ok(())
    }
}
