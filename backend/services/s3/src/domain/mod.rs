use aws_sdk_s3::types::ByteStream;
use manekani_service_common::repository::{
    error::Error, InsertError, RepoInsertable, RepoQueryable,
};

use crate::entity::file::{RequestCreate, RequestQuery};

pub struct FileStream {
    pub stream: ByteStream,
    pub size: u64,
}

impl FileStream {
    pub fn new(stream: ByteStream, size: u64) -> Self {
        Self { stream, size }
    }
}

/// Query a file in the given repository
///
/// # Errors
///
/// This function will return an error if the query was not successful
pub async fn query_file<R: RepoQueryable<RequestQuery, FileStream>>(
    repo: &R,
    req: RequestQuery,
) -> Result<FileStream, Error> {
    Ok(repo.query(req).await?)
}

/// Create a file in the given repository
///
/// # Errors
///
/// This function will return an error if the insertion was not successful
pub async fn create_file<R: RepoInsertable<RequestCreate, String>>(
    repo: &R,
    req: RequestCreate,
) -> Result<String, Error> {
    let file_name = req.file.name();
    if file_name.is_empty() || file_name.ends_with('/') {
        return Err(InsertError::BadRequest.into());
    }
    Ok(repo.insert(req).await?)
}
