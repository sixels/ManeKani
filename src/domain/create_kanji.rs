use sqlx::{pool::PoolConnection, Postgres};

use crate::entities::kanji::{InsertKanji, Kanji};

pub async fn execute(
    db: &mut PoolConnection<Postgres>,
    kanji: &InsertKanji,
) -> sqlx::Result<Kanji> {
    let InsertKanji {
        name,
        alt_names,
        symbol,
        reading,
        onyomi,
        kunyomi,
        nanori,
        meaning_mnemonic,
        reading_mnemonic,
    } = kanji;

    sqlx::query_as!(
        Kanji,
        "INSERT INTO kanjis
            (name, alt_names, symbol, reading, onyomi, kunyomi, nanori, meaning_mnemonic, reading_mnemonic)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *",
        name,
        alt_names,
        symbol,
        reading,
        onyomi,
        kunyomi,
        nanori,
        meaning_mnemonic,
        reading_mnemonic
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
    async fn it_should_create_a_new_kanji(pool: PgPool) -> sqlx::Result<()> {
        let mut conn = pool.acquire().await?;

        let kanji = InsertKanji::builder()
            .name("Finish")
            .alt_names(vec!["Complete".to_owned(), "End".to_owned()])
            .symbol("了")
            .meaning_mnemonic(
                r#"Think about it this way. There is a famous fishing lure inventor, working on his best work ever. He’s old, and he’s been trying to make the best fishing lure for the last 50 years, and knows this will be his last invention. Then he figures it out. He takes one barb, and connects the end of it to the top of another barb. When he does this, he knows his work is finally finished."#
            )
            .reading("りょう")
            .onyomi(vec!["りょう".to_owned()])
            .reading_mnemonic(
                r#"How does the fish lure maker test his newly finished lure out? The one he just finished? He gets in a row (りょう) boat and rows out into the sea."#
            )
            .build();

        let created_kanji = execute(&mut conn, &kanji).await?;

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
    async fn it_should_collide_with_an_existing_kanji(pool: PgPool) -> sqlx::Result<()> {
        let mut conn = pool.acquire().await?;

        let kanji = InsertKanji::builder()
            .name("Finish")
            .alt_names(vec!["Complete".to_owned(), "End".to_owned()])
            .symbol("了")
            .meaning_mnemonic(
                r#"Think about it this way. There is a famous fishing lure inventor, working on his best work ever. He’s old, and he’s been trying to make the best fishing lure for the last 50 years, and knows this will be his last invention. Then he figures it out. He takes one barb, and connects the end of it to the top of another barb. When he does this, he knows his work is finally finished."#
            )
            .reading("りょう")
            .onyomi(vec!["りょう".to_owned()])
            .reading_mnemonic(
                r#"How does the fish lure maker test his newly finished lure out? The one he just finished? He gets in a row (りょう) boat and rows out into the sea."#
            )
            .build();

        let _ = execute(&mut conn, &kanji).await?;
        let collision = execute(&mut conn, &kanji).await;

        assert!(matches!(collision, Err(sqlx::Error::Database(e))
            if e.code() == Some(SQLSTATE_ERROR_UNIQUE_INDEX_VIOLATION.into())));

        Ok(())
    }
}
