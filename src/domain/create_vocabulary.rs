use sqlx::{pool::PoolConnection, Postgres};

use crate::entities::vocabulary::{InsertVocabulary, Vocabulary};

pub async fn execute(
    db: &mut PoolConnection<Postgres>,
    vocab: &InsertVocabulary,
) -> sqlx::Result<Vocabulary> {
    let InsertVocabulary {
        name,
        alt_names,
        word,
        word_type,
        reading,
        meaning_mnemonic,
        reading_mnemonic,
    } = vocab;

    sqlx::query_as!(
        Vocabulary,
        "INSERT INTO vocabularies
            (name, alt_names, word, word_type, reading, meaning_mnemonic, reading_mnemonic)
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *",
        name,
        alt_names,
        word,
        word_type,
        reading,
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
    async fn it_should_create_a_new_vocabulary(pool: PgPool) -> sqlx::Result<()> {
        let mut conn = pool.acquire().await?;

        let radical = InsertVocabulary::builder()
            .name("Suspension")
            .alt_names(vec!["Cancellation".to_owned(), "Discontinuation".to_owned()])
            .word("中止")
            .word_type(vec!["noun".to_owned(), "する verb".to_owned()])
            .reading("ちゅうし")
            .meaning_mnemonic(
                r#"When something is in the middle of an action but it's stopped, this means it's put into suspension. Imagine getting frozen while you're running. It's kind of like that."#
            )
            .reading_mnemonic(
                r#"This is a jukugo word, which usually means on'yomi readings from the kanji. If you know the readings of your kanji you'll know how to read this as well."#
            )
            .build();

        let created_radical = execute(&mut conn, &radical).await?;

        assert_eq!(created_radical.name, radical.name);
        assert_eq!(created_radical.word, radical.word);
        assert_eq!(created_radical.word_type, radical.word_type);
        assert_eq!(created_radical.reading, radical.reading);
        assert_eq!(created_radical.meaning_mnemonic, radical.meaning_mnemonic);
        assert_eq!(created_radical.reading_mnemonic, radical.reading_mnemonic);
        assert_eq!(created_radical.user_synonyms, None);
        assert_eq!(created_radical.user_meaning_note, None);

        Ok(())
    }

    #[sqlx::test]
    async fn it_should_collide_with_an_existing_vocabulary(pool: PgPool) -> sqlx::Result<()> {
        let mut conn = pool.acquire().await?;

        let radical = InsertVocabulary::builder()
            .name("Suspension")
            .alt_names(vec!["Cancellation".to_owned(), "Discontinuation".to_owned()])
            .word("中止")
            .word_type(vec!["noun".to_owned(), "する verb".to_owned()])
            .reading("ちゅうし")
            .meaning_mnemonic(
                r#"When something is in the middle of an action but it's stopped, this means it's put into suspension. Imagine getting frozen while you're running. It's kind of like that."#
            )
            .reading_mnemonic(
                r#"This is a jukugo word, which usually means on'yomi readings from the kanji. If you know the readings of your kanji you'll know how to read this as well."#
            )
            .build();

        let _ = execute(&mut conn, &radical).await?;
        let collision = execute(&mut conn, &radical).await;

        assert!(matches!(collision, Err(sqlx::Error::Database(e))
            if e.code() == Some(SQLSTATE_ERROR_UNIQUE_INDEX_VIOLATION.into())));

        Ok(())
    }
}
