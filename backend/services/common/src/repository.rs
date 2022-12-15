pub mod error;

pub use error::{DeleteError, InsertError, QueryAllError, QueryError};

use self::error::UpdateError;

#[async_trait::async_trait]
pub trait RepoQueryable<Item, Output>: Send + Sync {
    async fn query(&self, item: Item) -> Result<Output, QueryError>;
}

#[async_trait::async_trait]
pub trait RepoQueryableAll<Output>: Send + Sync {
    async fn query_all(&self) -> Result<Vec<Output>, QueryAllError>;
}

#[async_trait::async_trait]
pub trait RepoInsertable<Item, Output>: Send + Sync {
    async fn insert(&self, item: Item) -> Result<Output, InsertError>;
}

#[async_trait::async_trait]
pub trait RepoDeletable<Item, Output>: Send + Sync {
    async fn delete(&self, item: Item) -> Result<Output, DeleteError>;
}

#[async_trait::async_trait]
pub trait RepoUpdateable<Item, Output>: Send + Sync {
    async fn update(&self, item: Item) -> Result<Output, UpdateError>;
}
