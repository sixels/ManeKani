use manekani_types::repository::{
    error::UpdateError, InsertError, QueryError, RepoInsertable, RepoQueryable, RepoUpdateable,
};

use crate::entity::{
    kanji::GetKanji,
    radical::{GetRadical, InsertRadical, Radical, RadicalPartial, UpdateRadical},
};

use super::Repository;

#[async_trait::async_trait]
impl RepoQueryable<GetRadical, Radical> for Repository {
    /// Query a radical
    async fn query(&self, radical: GetRadical) -> Result<Radical, QueryError> {
        let mut conn = self.connection().await;
        let GetRadical { name } = radical;

        let result = sqlx::query_as!(Radical, "SELECT * FROM radicals WHERE name = $1", name)
            .fetch_one(&mut conn)
            .await?;

        Ok(result)
    }
}

#[async_trait::async_trait]
impl RepoInsertable<InsertRadical, Radical> for Repository {
    /// Insert a radical
    async fn insert(&self, radical: InsertRadical) -> Result<Radical, InsertError> {
        let mut conn = self.connection().await;

        let InsertRadical {
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
impl RepoUpdateable<UpdateRadical, Radical> for Repository {
    /// Update a radical information
    async fn update(&self, radical: UpdateRadical) -> Result<Radical, UpdateError> {
        let mut conn = self.connection().await;

        let UpdateRadical {
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
impl RepoQueryable<GetKanji, Vec<RadicalPartial>> for Repository {
    /// Query a radical by kanji
    async fn query(&self, kanji: GetKanji) -> Result<Vec<RadicalPartial>, QueryError> {
        let mut conn = self.connection().await;

        let result = sqlx::query_as!(
            RadicalPartial,
            "SELECT r.id,r.name,r.symbol,r.level FROM radicals r
                INNER JOIN kanjis_radicals kr ON r.name = kr.radical_name
                    AND kr.kanji_symbol = $1",
            kanji.symbol
        )
        .fetch_all(&mut conn)
        .await?;

        Ok(result)
    }
}
