use manekani_service_common::repository::{InsertError, QueryError, RepoInsertable, RepoQueryable};
use sqlx::{postgres::PgArguments, Acquire, Arguments};

use crate::model::{
    ReqKanjiQuery, ReqVocabularyInsert, ReqVocabularyQuery, Vocabulary, VocabularyPartial,
};

use super::Repository;

#[async_trait::async_trait]
impl RepoQueryable<ReqVocabularyQuery, Vocabulary> for Repository {
    /// Query a vocabulary
    async fn query(&self, vocab: ReqVocabularyQuery) -> Result<Vocabulary, QueryError> {
        let mut conn = self.connection().await;

        let ReqVocabularyQuery { word } = vocab;

        let result = sqlx::query_as!(
            Vocabulary,
            "SELECT * FROM vocabularies WHERE word = $1",
            word
        )
        .fetch_one(&mut conn)
        .await?;

        Ok(result)
    }
}

#[async_trait::async_trait]
impl RepoInsertable<ReqVocabularyInsert, Vocabulary> for Repository {
    /// Query a vocabulary
    async fn insert(&self, vocab: ReqVocabularyInsert) -> Result<Vocabulary, InsertError> {
        let mut conn = self.connection().await;

        let ReqVocabularyInsert {
            name,
            level,
            alt_names,
            word,
            word_type,
            reading,
            meaning_mnemonic,
            reading_mnemonic,
            kanji_composition,
        } = vocab;

        let mut transaction = conn.begin().await?;

        let insert_vocabulary = sqlx::query_as!(
            Vocabulary,
            "INSERT INTO vocabularies
                (name, level, alt_names, word, word_type, reading, meaning_mnemonic, reading_mnemonic)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *",
            name,
            level,
            &alt_names,
            word,
            &word_type,
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

        let result = sqlx::query!(
            "INSERT INTO vocabularies_kanjis (vocabulary_id, kanji_symbol) SELECT v.id, k.symbol FROM vocabularies v INNER JOIN kanjis k ON v.id = $1 AND k.symbol = ANY($2)",
            insert_vocabulary.id,
            &kanji_composition
        ).execute(&mut transaction)
            .await?;

        if result.rows_affected() == 0 {
            return Err(InsertError::BadRequest);
        }

        transaction.commit().await?;

        Ok(insert_vocabulary)
    }
}

#[async_trait::async_trait]
impl RepoQueryable<ReqKanjiQuery, Vec<VocabularyPartial>> for Repository {
    /// Query a vocabulary by kanji
    async fn query(&self, kanji: ReqKanjiQuery) -> Result<Vec<VocabularyPartial>, QueryError> {
        let mut conn = self.connection().await;

        let result = sqlx::query_as!(
            VocabularyPartial,
            "SELECT id,name,word,reading,level FROM vocabularies v
                INNER JOIN vocabularies_kanjis vk ON v.id = vk.vocabulary_id
                AND vk.kanji_symbol = $1",
            kanji.symbol
        )
        .fetch_all(&mut conn)
        .await?;

        Ok(result)
    }
}
