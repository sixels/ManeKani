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
            symbol,
            meaning_mnemonic,
        } = radical;

        let result = sqlx::query_as!(
            Radical,
            "INSERT INTO radicals (name, symbol, meaning_mnemonic) VALUES ($1, $2, $3) RETURNING *",
            name,
            symbol,
            meaning_mnemonic
        )
        .fetch_one(self)
        .await?;

        Ok(result)
    }
    async fn get(&mut self, req: &GetRadical) -> Result<Radical, RepositoryError> {
        let GetRadical { symbol } = req;

        let result = sqlx::query_as!(Radical, "SELECT * FROM radicals WHERE symbol = $1", symbol)
            .fetch_one(self)
            .await?;

        Ok(result)
    }
}
