// use crate::entitity::{Deletable, Insertable, Queryable};

#[async_trait::async_trait]
pub trait RepoQueryable<Entity, Output>: Send + Sync {
    async fn query(&self, item: Entity) -> Result<Output, QueryError>;
}

#[async_trait::async_trait]
pub trait RepoQueryableAll<Output>: Send + Sync {
    async fn query_all(&self) -> Result<Vec<Output>, QueryAllError>;
}

#[async_trait::async_trait]
pub trait RepoInsertable<Entity, Output>: Send + Sync {
    async fn insert(&self, item: Entity) -> Result<Output, InsertError>;
}

#[async_trait::async_trait]
pub trait RepoDeletable<Entity, Output>: Send + Sync {
    async fn delete(&self, item: Entity) -> Result<Output, DeleteError>;
}

pub enum QueryError {
    NotFound,
    Unknown,
}

pub enum QueryAllError {
    Unknown,
}

pub enum InsertError {
    Conflict,
    BadRequest,
    Unknown,
}

pub enum DeleteError {
    NotFound,
    Unknown,
}

impl From<sqlx::Error> for QueryError {
    fn from(error: sqlx::Error) -> Self {
        match error {
            sqlx::Error::RowNotFound => Self::NotFound,
            _ => Self::Unknown,
        }
    }
}

impl From<sqlx::Error> for QueryAllError {
    fn from(_: sqlx::Error) -> Self {
        Self::Unknown
    }
}

impl From<sqlx::Error> for InsertError {
    fn from(error: sqlx::Error) -> Self {
        match error {
            sqlx::Error::Database(d) if d.code() == Some("23505".into()) => Self::Conflict,
            _ => Self::Unknown,
        }
    }
}

impl From<sqlx::Error> for DeleteError {
    fn from(error: sqlx::Error) -> Self {
        match error {
            sqlx::Error::RowNotFound => Self::NotFound,
            _ => Self::Unknown,
        }
    }
}
