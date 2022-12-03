use manekani_types::repository::{InsertError, QueryError, RepoInsertable, RepoQueryable};

use crate::entity::{
    kanji::GetKanji,
    radical::{GetRadical, InsertRadical, Radical, RadicalPartial},
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

        let symbol = if let Some('=') = symbol.chars().last() {
            // base64::decode(symbol).unwrap()
            todo!("We will no longer store a base64 string for images. A key-value database will be used to handle this")
        } else {
            symbol.clone().into_bytes()
        };

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
