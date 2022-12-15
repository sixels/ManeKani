use manekani_service_common::repository::{
    error::UpdateError, InsertError, QueryError, RepoInsertable, RepoQueryable, RepoUpdateable,
};

use crate::model::{
    Radical, RadicalPartial, ReqKanjiQuery, ReqRadicalInsert, ReqRadicalQuery, ReqRadicalUpdate,
};

use super::Repository;

#[async_trait::async_trait]
impl RepoQueryable<ReqRadicalQuery, Radical> for Repository {
    /// Query a radical
    async fn query(&self, radical: ReqRadicalQuery) -> Result<Radical, QueryError> {
        let mut conn = self.connection().await;
        let ReqRadicalQuery { name } = radical;

        let result = sqlx::query_as!(Radical, "SELECT * FROM radicals WHERE name = $1", name)
            .fetch_one(&mut conn)
            .await?;

        Ok(result)
    }
}

#[async_trait::async_trait]
impl RepoInsertable<ReqRadicalInsert, Radical> for Repository {
    /// Insert a radical
    async fn insert(&self, radical: ReqRadicalInsert) -> Result<Radical, InsertError> {
        let mut conn = self.connection().await;

        let ReqRadicalInsert {
            name,
            level,
            symbol,
            meaning_mnemonic,
        } = radical;

        let result = sqlx::query_as!(
            Radical,
            "INSERT INTO radicals (name, level, symbol, meaning_mnemonic) VALUES ($1, $2, $3, $4) RETURNING *",
            name,
            level,
            symbol,
            meaning_mnemonic
        )
        .fetch_one(&mut conn)
        .await?;

        Ok(result)
    }
}

#[async_trait::async_trait]
impl RepoUpdateable<ReqRadicalUpdate, Radical> for Repository {
    /// Update a radical information
    async fn update(&self, radical: ReqRadicalUpdate) -> Result<Radical, UpdateError> {
        let mut conn = self.connection().await;

        let ReqRadicalUpdate {
            name,
            symbol,
            level,
            meaning_mnemonic,
            user_synonyms,
            user_meaning_note,
        } = radical;

        let user_synonyms = user_synonyms.as_deref();

        let result = sqlx::query_as!(
            Radical,
            "UPDATE radicals SET 
                symbol = COALESCE($1, symbol),
                level = COALESCE($2, level),
                meaning_mnemonic = COALESCE($3, meaning_mnemonic),
                user_synonyms = COALESCE($4, user_synonyms),
                user_meaning_note = COALESCE($5, user_meaning_note)
            WHERE name = $6
            RETURNING *",
            symbol,
            level,
            meaning_mnemonic,
            user_synonyms,
            user_meaning_note,
            name
        )
        .fetch_one(&mut conn)
        .await?;

        Ok(result)
    }
}

#[async_trait::async_trait]
impl RepoQueryable<ReqKanjiQuery, Vec<RadicalPartial>> for Repository {
    /// Query a radical by kanji
    async fn query(&self, kanji: ReqKanjiQuery) -> Result<Vec<RadicalPartial>, QueryError> {
        let mut conn = self.connection().await;

        let result = sqlx::query_as!(
            RadicalPartial,
            "SELECT r.id,r.name,r.symbol,r.level FROM radicals r
                INNER JOIN kanji_radicals kr ON r.name = kr.radical_name
                    AND kr.kanji_symbol = $1",
            kanji.symbol
        )
        .fetch_all(&mut conn)
        .await?;

        Ok(result)
    }
}
