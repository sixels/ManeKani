use manekani_types::repository::{RepoInsertable, RepoQueryable, RepoUpdateable};

use crate::entity::{
    kanji::GetKanji,
    radical::{GetRadical, InsertRadical, Radical, RadicalPartial, UpdateRadical},
};

use super::Error;

pub async fn query<R>(repo: &R, radical: GetRadical) -> Result<Radical, Error>
where
    R: RepoQueryable<GetRadical, Radical>,
{
    Ok(repo.query(radical).await?)
}

pub async fn query_by_kanji<R>(repo: &R, kanji: GetKanji) -> Result<Vec<RadicalPartial>, Error>
where
    R: RepoQueryable<GetKanji, Vec<RadicalPartial>>,
{
    Ok(repo.query(kanji).await?)
}

pub async fn insert<R>(repo: &R, radical: InsertRadical) -> Result<Radical, Error>
where
    R: RepoInsertable<InsertRadical, Radical>,
{
    Ok(repo.insert(radical).await?)
}

pub async fn update<R>(repo: &R, radical: UpdateRadical) -> Result<Radical, Error>
where
    R: RepoUpdateable<UpdateRadical, Radical>,
{
    Ok(repo.update(radical).await?)
}

#[cfg(test)]
mod tests {
    use crate::repository::Repository;

    use super::*;

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

        let created_radical = insert(&repo, radical.clone()).await?;

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

        let _ = insert(&repo, radical.clone()).await?;
        let radical = insert(&repo, radical).await;

        assert!(matches!(radical, Err(Error::Conflict)));

        Ok(())
    }
}
