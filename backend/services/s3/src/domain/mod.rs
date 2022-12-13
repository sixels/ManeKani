use std::pin::Pin;

use bytes::Bytes;
use futures_util::Stream;
use manekani_service_common::repository::{
    error::Error, InsertError, RepoInsertable, RepoQueryable,
};

use crate::entity::file::{CreateFile, QueryFile};

pub type FileStream = Pin<Box<dyn Stream<Item = Result<Bytes, Box<dyn std::error::Error>>>>>;

pub async fn query_file<R: RepoQueryable<QueryFile, (u64, FileStream)>>(
    repo: &R,
    req: QueryFile,
) -> Result<(u64, FileStream), Error> {
    Ok(repo.query(req).await?)
}

pub async fn create_file<R: RepoInsertable<CreateFile, String>>(
    repo: &R,
    req: CreateFile,
) -> Result<String, Error> {
    let file_name = req.file.name();
    if file_name.is_empty() || file_name.ends_with('/') {
        return Err(InsertError::BadRequest.into());
    }
    Ok(repo.insert(req).await?)
}
