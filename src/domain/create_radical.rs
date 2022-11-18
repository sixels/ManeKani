use sqlx::{pool::PoolConnection, Postgres};

use crate::entities::radical::{InsertRadical, Radical};

pub async fn execute(
    db: &mut PoolConnection<Postgres>,
    radical: &InsertRadical,
) -> sqlx::Result<Radical> {
    let InsertRadical {
        name,
        symbol,
        meaning_mnemonic,
    } = radical;

    sqlx::query_as!(
        Radical,
        "INSERT INTO radicals (name, symbol, meaning_mnemonic) VALUES ($1, $2, $3) RETURNING *",
        name,
        symbol,
        meaning_mnemonic
    )
    .fetch_one(db)
    .await
}

#[cfg(test)]
mod tests {
    use super::*;

    use sqlx::PgPool;

    // https://www.ibm.com/docs/en/db2-for-zos/13?topic=codes-sqlstate-values-common-error
    const SQLSTATE_ERROR_UNIQUE_INDEX_VIOLATION: &str = "23505";

    #[sqlx::test]
    async fn it_should_create_a_new_radical(pool: PgPool) -> sqlx::Result<()> {
        let mut conn = pool.acquire().await?;

        let radical = InsertRadical::builder()
            .name("Barb")
            .symbol("亅")
            .meaning_mnemonic(
                r#"This radical is shaped like a barb. Like the kind you'd see on barb wire. Imagine one of these getting stuck to your arm or your clothes. Think about how much it would hurt with that little hook on the end sticking into you. Say out loud, "Oh dang, I got a barb stuck in me!""#
            )
            .build();

        let created_radical = execute(&mut conn, &radical).await?;

        assert_eq!(created_radical.name, radical.name);
        assert_eq!(created_radical.symbol, radical.symbol);
        assert_eq!(created_radical.meaning_mnemonic, radical.meaning_mnemonic);
        assert_eq!(created_radical.user_synonyms, None);
        assert_eq!(created_radical.user_meaning_note, None);

        Ok(())
    }

    #[sqlx::test]
    async fn it_should_collide_with_an_existing_radical(pool: PgPool) -> sqlx::Result<()> {
        let mut conn = pool.acquire().await?;

        let radical = InsertRadical::builder()
            .name("Barb")
            .symbol("亅")
            .meaning_mnemonic(
                r#"This radical is shaped like a barb. Like the kind you'd see on barb wire. Imagine one of these getting stuck to your arm or your clothes. Think about how much it would hurt with .to_owned()that little hook on the end sticking into you. Say out loud, "Oh dang, I got a barb stuck in me!""#
            )
            .build();

        let _ = execute(&mut conn, &radical).await?;
        let radical = execute(&mut conn, &radical).await;

        assert!(matches!(radical, Err(sqlx::Error::Database(e))
            if e.code() == Some(SQLSTATE_ERROR_UNIQUE_INDEX_VIOLATION.into())));

        Ok(())
    }
}
