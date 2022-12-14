use aws_sdk_s3::types::ByteStream;
use manekani_service_common::repository::{
    error::Error, InsertError, RepoInsertable, RepoQueryable,
};

use crate::entity::file::{CreateFile, QueryFile};

pub struct FileStream {
    pub stream: ByteStream,
    pub size: u64,
}

impl FileStream {
    pub fn new(stream: ByteStream, size: u64) -> Self {
        Self { stream, size }
    }
}

pub async fn query_file<R: RepoQueryable<QueryFile, FileStream>>(
    repo: &R,
    req: QueryFile,
) -> Result<FileStream, Error> {
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
