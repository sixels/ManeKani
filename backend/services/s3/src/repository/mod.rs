use aws_sdk_s3::{types::ByteStream, Client as S3Client, Endpoint, Region};
use manekani_service_common::repository::{InsertError, QueryError, RepoInsertable, RepoQueryable};

use crate::{
    domain::FileStream,
    model::file::{RequestCreate, RequestQuery},
};

#[derive(Clone)]
pub struct S3Repo {
    client: S3Client,
    bucket: String,
}

impl S3Repo {
    /// Creates a new S3 repository
    ///
    /// # Panics
    ///
    /// Panics if the s3 configuration is not valid
    pub async fn new(bucket_name: String) -> Self {
        let region = Region::from_static("sa-east-1");
        let endpoint = Endpoint::immutable("http://127.0.0.1:9000".try_into().unwrap());
        let config = aws_config::from_env()
            .region(region)
            .endpoint_resolver(endpoint)
            .load()
            .await;

        let client = S3Client::new(&config);
        let bucket = bucket_name;

        Self { client, bucket }
    }
}

#[async_trait::async_trait]
impl RepoQueryable<RequestQuery, FileStream> for S3Repo {
    async fn query(&self, file: RequestQuery) -> Result<FileStream, QueryError> {
        let key = file.as_s3_key();

        let object = self
            .client
            .get_object()
            .bucket(&self.bucket)
            .key(&key)
            .send()
            .await?;

        match object.content_length().try_into() {
            Ok(size) => Ok(FileStream::new(object.body, size)),
            Err(e) => Err(QueryError::Unknown(Box::new(e))),
        }
    }
}

#[async_trait::async_trait]
impl RepoInsertable<RequestCreate, String> for S3Repo {
    async fn insert(&self, file: RequestCreate) -> Result<String, InsertError> {
        let key = file.as_s3_key();
        let tmp = file.file;

        let contents = match tmp.read_all().await {
            Ok(contents) => contents,
            Err(e) => return Err(InsertError::Unknown(Box::new(e))),
        };

        self.client
            .put_object()
            .bucket(&self.bucket)
            .key(&key)
            .body(ByteStream::from(contents))
            .send()
            .await?;

        Ok(key)
    }
}
