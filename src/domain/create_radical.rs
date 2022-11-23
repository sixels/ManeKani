use crate::{
    entities::radical::{InsertRadical, Radical},
    repositories::{radical::RadicalRepository, RepositoryError},
};

pub async fn execute<R>(db: &mut R, radical: &InsertRadical) -> Result<Radical, RepositoryError>
where
    R: RadicalRepository,
{
    db.insert(radical).await
}

#[cfg(test)]
mod tests {
    use super::*;

    use sqlx::PgPool;

    #[sqlx::test]
    async fn it_should_create_a_new_radical(pool: PgPool) -> Result<(), RepositoryError> {
        let mut conn = pool.acquire().await?;

        let radical = InsertRadical::builder()
            .name("barb")
            .level(1)
            .symbol("亅")
            .meaning_mnemonic(
                r#"This radical is shaped like a barb. Like the kind you'd see on barb wire. Imagine one of these getting stuck to your arm or your clothes. Think about how much it would hurt with that little hook on the end sticking into you. Say out loud, "Oh dang, I got a barb stuck in me!""#
            )
            .build();

        let created_radical = execute(&mut conn, &radical).await?;

        assert_eq!(created_radical.name, radical.name);
        assert_eq!(&created_radical.symbol, radical.symbol.as_bytes());
        assert_eq!(created_radical.meaning_mnemonic, radical.meaning_mnemonic);
        assert_eq!(created_radical.user_synonyms, None);
        assert_eq!(created_radical.user_meaning_note, None);

        Ok(())
    }

    #[sqlx::test]
    async fn it_should_collide_with_an_existing_radical(
        pool: PgPool,
    ) -> Result<(), RepositoryError> {
        let mut conn = pool.acquire().await?;

        let radical = InsertRadical::builder()
            .name("barb")
            .level(1)
            .symbol("亅")
            .meaning_mnemonic(
                r#"This radical is shaped like a barb. Like the kind you'd see on barb wire. Imagine one of these getting stuck to your arm or your clothes. Think about how much it would hurt with .to_owned()that little hook on the end sticking into you. Say out loud, "Oh dang, I got a barb stuck in me!""#
            )
            .build();

        let _ = execute(&mut conn, &radical).await?;
        let radical = execute(&mut conn, &radical).await;

        assert!(matches!(radical, Err(RepositoryError::Conflict)));

        Ok(())
    }
}
