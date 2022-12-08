pub mod error;

pub use error::Error;
use manekani_types::repository::{RepoInsertable, RepoQueryable};

use crate::entity::file::{CreateFile, QueryFile};

pub async fn query_file<R: RepoQueryable<QueryFile, String>>(
    repo: &R,
    req: QueryFile,
) -> Result<String, Error> {
    Ok(repo.query(req).await?)
}

pub async fn create_file<R: RepoInsertable<CreateFile, String>>(
    repo: &R,
    req: CreateFile,
) -> Result<String, Error> {
    let file_name = req.file.name();
    if file_name.is_empty() || file_name.ends_with('/') {
        return Err(Error::BadRequest);
    }
    Ok(repo.insert(req).await?)
}
