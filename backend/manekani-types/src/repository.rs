pub mod error;

pub use error::{DeleteError, InsertError, QueryAllError, QueryError};

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
