use sqlx::{pool::PoolConnection, postgres::PgArguments, Arguments, Connection, Postgres};

use crate::entities::kanji::{GetKanji, InsertKanji, Kanji};

use super::RepositoryError;

#[async_trait::async_trait]
pub trait KanjiRepository {
    async fn insert(&mut self, kanji: &InsertKanji) -> Result<Kanji, RepositoryError>;
    async fn get(&mut self, req: &GetKanji) -> Result<Kanji, RepositoryError>;
}

#[async_trait::async_trait]
impl KanjiRepository for PoolConnection<Postgres> {
    async fn insert(&mut self, kanji: &InsertKanji) -> Result<Kanji, RepositoryError> {
        let InsertKanji {
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

        let mut transaction = self.begin().await?;

        let insert_kanji = sqlx::query_as!(
            Kanji,
            "INSERT INTO kanjis
                (name, level, alt_names, symbol, reading, onyomi, kunyomi, nanori, meaning_mnemonic, reading_mnemonic)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *",
            name,
            level,
            alt_names,
            symbol,
            reading,
            onyomi,
            kunyomi,
            nanori,
            meaning_mnemonic,
            reading_mnemonic
        )
        .fetch_one(&mut transaction)
        .await?;

        let mut args = PgArguments::default();
        let mut sql = String::from("INSERT INTO kanjis_radicals (kanji_id, radical_name) SELECT k.id, r.name FROM kanjis k INNER JOIN radicals r ON k.id = $1 AND (r.name = $2");
        args.add(insert_kanji.id);
        args.add(&radical_composition[0]);
        for (n, radical) in radical_composition.iter().enumerate().skip(1) {
            sql.push_str(&format!(" OR r.name = ${}", n + 2));
            args.add(radical);
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

        Ok(insert_kanji)
    }

    async fn get(&mut self, req: &GetKanji) -> Result<Kanji, RepositoryError> {
        let GetKanji { symbol } = req;

        let result = sqlx::query_as!(Kanji, "SELECT * FROM kanjis WHERE symbol = $1", symbol)
            .fetch_one(self)
            .await?;

        Ok(result)
    }
}
