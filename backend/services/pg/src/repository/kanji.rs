use manekani_service_common::repository::{InsertError, QueryError, RepoInsertable, RepoQueryable};
use sqlx::Connection;

use crate::entity::{Kanji, KanjiPartial, ReqKanjiInsert, ReqKanjiQuery, ReqRadicalQuery};

use super::Repository;

#[async_trait::async_trait]
impl RepoQueryable<ReqKanjiQuery, Kanji> for Repository {
    /// Query a kanji
    async fn query(&self, kanji: ReqKanjiQuery) -> Result<Kanji, QueryError> {
        let mut conn = self.connection().await;

        let result = sqlx::query_as!(
            Kanji,
            "SELECT * FROM kanjis WHERE symbol = $1",
            kanji.symbol
        )
        .fetch_one(&mut conn)
        .await?;

        Ok(result)
    }
}

#[async_trait::async_trait]
impl RepoInsertable<ReqKanjiInsert, Kanji> for Repository {
    /// Insert a kanji
    async fn insert(&self, kanji: ReqKanjiInsert) -> Result<Kanji, InsertError> {
        let mut conn = self.connection().await;

        let ReqKanjiInsert {
            name,
            level,
            alt_names,
            symbol,
            reading,
            onyomi,
            kunyomi,
            nanori,
            meaning_mnemonic,
            reading_mnemonic,
            radical_composition,
        } = kanji;

        let mut transaction = conn.begin().await?;

        let insert_kanji = sqlx::query_as!(
            Kanji,
            "INSERT INTO kanjis
                (name, level, alt_names, symbol, reading, onyomi, kunyomi, nanori, meaning_mnemonic, reading_mnemonic)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *",
            name,
            level,
            &alt_names,
            symbol,
            reading,
            &onyomi,
            &kunyomi,
            &nanori,
            meaning_mnemonic,
            reading_mnemonic
        )
        .fetch_one(&mut transaction)
        .await?;

        let result = sqlx::query!(
            "INSERT INTO kanjis_radicals (kanji_symbol, radical_name) SELECT k.symbol, r.name FROM kanjis k INNER JOIN radicals r ON k.id = $1 AND r.name = ANY($2)",
            insert_kanji.id,
            &radical_composition
        ).execute(&mut transaction).await?;

        if result.rows_affected() == 0 {
            return Err(InsertError::BadRequest);
        }

        transaction.commit().await?;

        Ok(insert_kanji)
    }
}

#[async_trait::async_trait]
impl RepoQueryable<ReqRadicalQuery, Vec<KanjiPartial>> for Repository {
    /// Query a kanji by radical
    async fn query(&self, radical: ReqRadicalQuery) -> Result<Vec<KanjiPartial>, QueryError> {
        let mut conn = self.connection().await;

        let result = sqlx::query_as!(
            KanjiPartial,
            "SELECT k.id,k.symbol,k.name,k.reading,k.level FROM kanjis k
                INNER JOIN kanjis_radicals kr ON k.symbol = kr.kanji_symbol
                    AND kr.radical_name = $1",
            radical.name
        )
        .fetch_all(&mut conn)
        .await?;

        Ok(result)
    }
}
