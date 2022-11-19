use sqlx::{pool::PoolConnection, Postgres};

use crate::entities::kanji::{InsertKanji, Kanji};

use super::RepositoryError;

#[async_trait::async_trait]
pub trait KanjiRepository {
    async fn insert(&mut self, kanji: &InsertKanji) -> Result<Kanji, RepositoryError>;
}

#[async_trait::async_trait]
impl KanjiRepository for PoolConnection<Postgres> {
    async fn insert(&mut self, kanji: &InsertKanji) -> Result<Kanji, RepositoryError> {
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

        let result = sqlx::query_as!(
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
        .fetch_one(self)
        .await?;

        Ok(result)
    }
}
