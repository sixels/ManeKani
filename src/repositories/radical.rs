use sqlx::{pool::PoolConnection, Postgres};

use crate::entities::radical::{GetRadical, InsertRadical, Radical};

use super::RepositoryError;

#[async_trait::async_trait]
pub trait RadicalRepository {
    async fn insert(&mut self, radical: &InsertRadical) -> Result<Radical, RepositoryError>;
    async fn get(&mut self, req: &GetRadical) -> Result<Radical, RepositoryError>;
}

#[async_trait::async_trait]
impl RadicalRepository for PoolConnection<Postgres> {
    async fn insert(&mut self, radical: &InsertRadical) -> Result<Radical, RepositoryError> {
        let InsertRadical {
            name,
            level,
            symbol,
            meaning_mnemonic,
        } = radical;

        let symbol = if let Some('=') = symbol.chars().last() {
            base64::decode(symbol).unwrap()
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
        .fetch_one(self)
        .await?;

        Ok(result)
    }
    async fn get(&mut self, req: &GetRadical) -> Result<Radical, RepositoryError> {
        let GetRadical { name } = req;

        let result = sqlx::query_as!(Radical, "SELECT * FROM radicals WHERE name = $1", name)
            .fetch_one(self)
            .await?;

        Ok(result)
    }
}
