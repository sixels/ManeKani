use sqlx::{pool::PoolConnection, Postgres};

use crate::entities::vocabulary::{InsertVocabulary, Vocabulary};

use super::RepositoryError;

#[async_trait::async_trait]
pub trait VocabularyRepository {
    async fn insert(&mut self, vocab: &InsertVocabulary) -> Result<Vocabulary, RepositoryError>;
}

#[async_trait::async_trait]
impl VocabularyRepository for PoolConnection<Postgres> {
    async fn insert(&mut self, vocab: &InsertVocabulary) -> Result<Vocabulary, RepositoryError> {
        let InsertVocabulary {
            name,
            alt_names,
            word,
            word_type,
            reading,
            meaning_mnemonic,
            reading_mnemonic,
        } = vocab;

        let result = sqlx::query_as!(
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
        .fetch_one(self)
        .await?;

        Ok(result)
    }
}
