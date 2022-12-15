use manekani_service_common::repository::{
    error::UpdateError, InsertError, QueryError, RepoInsertable, RepoQueryable, RepoUpdateable,
};
use sqlx::Connection;

use crate::model::{
    kanji::ReqKanjiUpdate, Kanji, KanjiPartial, ReqKanjiInsert, ReqKanjiQuery, ReqRadicalQuery,
};

use super::Repository;

#[async_trait::async_trait]
impl RepoQueryable<ReqKanjiQuery, Kanji> for Repository {
    /// Query a kanji
    async fn query(&self, kanji: ReqKanjiQuery) -> Result<Kanji, QueryError> {
        let mut conn = self.connection().await;

        let result = sqlx::query_as!(Kanji, "SELECT * FROM kanji WHERE symbol = $1", kanji.symbol)
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
            "INSERT INTO kanji
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

        if let Some(radical_composition) = radical_composition.filter(|rc| !rc.is_empty()) {
            let result = sqlx::query!(
                "INSERT INTO kanji_radicals (kanji_symbol, radical_name)
                    SELECT k.symbol, r.name FROM kanji k
                        INNER JOIN radicals r ON
                            k.id = $1 AND r.name = ANY($2)",
                insert_kanji.id,
                &radical_composition
            )
            .execute(&mut transaction)
            .await?;

            if result.rows_affected() == 0 {
                return Err(InsertError::BadRequest);
            }
        }

        transaction.commit().await?;

        Ok(insert_kanji)
    }
}

#[async_trait::async_trait]
impl RepoUpdateable<ReqKanjiUpdate, Kanji> for Repository {
    // TODO: RETURN A ResKanjiUpdate containing the radical_compostion too
    async fn update(&self, kanji: ReqKanjiUpdate) -> Result<Kanji, UpdateError> {
        let mut conn = self.connection().await;

        let ReqKanjiUpdate {
            id,
            symbol,
            level,
            name,
            alt_names,
            meaning_mnemonic,
            reading,
            reading_mnemonic,
            onyomi,
            kunyomi,
            nanori,
            user_synonyms,
            user_meaning_note,
            user_reading_note,
            radical_composition,
        } = kanji;

        let mut transaction = conn.begin().await?;

        let kanji = sqlx::query_as!(
            Kanji,
            "UPDATE kanji SET
                    symbol = COALESCE(symbol, $1),
                    level = COALESCE(level, $2),
                    name = COALESCE(name, $3),
                    alt_names = COALESCE(alt_names, $4),
                    meaning_mnemonic = COALESCE(meaning_mnemonic, $5),
                    reading = COALESCE(reading, $6),
                    reading_mnemonic = COALESCE(reading_mnemonic, $7),
                    onyomi = COALESCE(onyomi, $8),
                    kunyomi = COALESCE(kunyomi, $9),
                    nanori = COALESCE(nanori, $10),
                    user_synonyms = COALESCE(user_synonyms, $11),
                    user_meaning_note = COALESCE(user_meaning_note, $12),
                    user_reading_note = COALESCE(user_reading_note, $13)
                WHERE id = $14
                RETURNING *",
            symbol,
            level,
            name,
            alt_names.as_deref(),
            meaning_mnemonic,
            reading,
            reading_mnemonic,
            onyomi.as_deref(),
            kunyomi.as_deref(),
            nanori.as_deref(),
            user_synonyms.as_deref(),
            user_meaning_note,
            user_reading_note,
            id,
        )
        .fetch_one(&mut transaction)
        .await?;

        if let Some(radical_composition) = radical_composition {
            // Delete every relations and create them again.
            // It would be better with an ORM 😮‍💨️.
            sqlx::query!(
                "DELETE from kanji_radicals WHERE kanji_symbol = $1 AND radical_name = ANY($2)",
                kanji.symbol,
                &radical_composition,
            )
            .execute(&mut transaction)
            .await?;

            let result = sqlx::query!(
                "INSERT INTO kanji_radicals (kanji_symbol, radical_name)
                    SELECT k.symbol, r.name FROM kanji k
                        INNER JOIN radicals r ON
                            k.id = $1 AND r.name = ANY($2)",
                kanji.id,
                &radical_composition
            )
            .execute(&mut transaction)
            .await?;

            if result.rows_affected() == 0 && !radical_composition.is_empty() {
                return Err(UpdateError::BadRequest);
            }
        }

        transaction.commit().await?;

        Ok(kanji)
    }
}

#[async_trait::async_trait]
impl RepoQueryable<ReqRadicalQuery, Vec<KanjiPartial>> for Repository {
    /// Query a kanji by radical
    async fn query(&self, radical: ReqRadicalQuery) -> Result<Vec<KanjiPartial>, QueryError> {
        let mut conn = self.connection().await;

        let result = sqlx::query_as!(
            KanjiPartial,
            "SELECT k.id,k.symbol,k.name,k.reading,k.level FROM kanji k
                INNER JOIN kanji_radicals kr ON
                    k.symbol = kr.kanji_symbol AND kr.radical_name = $1",
            radical.name
        )
        .fetch_all(&mut conn)
        .await?;

        Ok(result)
    }
}
