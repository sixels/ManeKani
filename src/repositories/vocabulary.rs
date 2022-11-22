use sqlx::{pool::PoolConnection, postgres::PgArguments, Acquire, Arguments, Postgres};

use crate::entities::vocabulary::{GetVocabulary, InsertVocabulary, Vocabulary};

use super::RepositoryError;

#[async_trait::async_trait]
pub trait VocabularyRepository {
    async fn insert(&mut self, vocab: &InsertVocabulary) -> Result<Vocabulary, RepositoryError>;
    async fn get(&mut self, req: &GetVocabulary) -> Result<Vocabulary, RepositoryError>;
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
            kanji_composition,
        } = vocab;

        let mut transaction = self.begin().await?;

        let insert_vocabulary = sqlx::query_as!(
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
        .fetch_one(&mut transaction)
        .await?;

        let mut args = PgArguments::default();
        let mut sql = String::from("INSERT INTO vocabularies_kanjis (vocabulary_id, kanji_symbol) SELECT v.id, k.symbol FROM vocabularies v INNER JOIN kanjis k ON v.id = $1 AND (k.symbol = $2");
        args.add(insert_vocabulary.id);
        args.add(&kanji_composition[0]);
        for (n, kanji) in kanji_composition.iter().enumerate().skip(1) {
            sql.push_str(&format!(" OR k.symbol = ${}", n + 2));
            args.add(kanji);
        }
        sql.push(')');

        let result = sqlx::query_with(&sql, args)
            .execute(&mut transaction)
            .await?;
        if result.rows_affected() == 0 {
            return Err(RepositoryError::BadRequest(
                "One or more radicals in `radical_composition` does not exist.".to_owned(),
            ));
        }

        transaction.commit().await?;

        Ok(insert_vocabulary)
    }
    async fn get(&mut self, req: &GetVocabulary) -> Result<Vocabulary, RepositoryError> {
        let GetVocabulary { word } = req;

        let result = sqlx::query_as!(
            Vocabulary,
            "SELECT * FROM vocabularies WHERE word = $1",
            word
        )
        .fetch_one(self)
        .await?;

        Ok(result)
    }
}
