use aws_sdk_s3::{types::ByteStream, Client as S3Client, Endpoint, Region};
use bytes::Bytes;
use futures_util::{Stream, TryStreamExt};
use manekani_types::repository::{InsertError, QueryError, RepoInsertable, RepoQueryable};

use crate::entity::file::{CreateFile, QueryFile};

#[derive(Clone)]
pub struct S3Repo {
    client: S3Client,
    bucket: String,
}

impl S3Repo {
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

type FileStream = Box<dyn Stream<Item = Result<Bytes, Box<dyn std::error::Error>>>>;
#[async_trait::async_trait]
impl RepoQueryable<QueryFile, (u64, FileStream)> for S3Repo {
    async fn query(&self, file: QueryFile) -> Result<(u64, FileStream), QueryError> {
        let key = file.as_s3_key();

        // TODO: handle errors
        let object = self
            .client
            .get_object()
            .bucket(&self.bucket)
            .key(&key)
            .send()
            .await
            .unwrap();

        let size: u64 = object
            .content_length()
            .try_into()
            .expect("Invalid file size");

        Ok((size, Box::new(object.body.map_err(Into::into))))
    }
}

#[async_trait::async_trait]
impl RepoInsertable<CreateFile, String> for S3Repo {
    async fn insert(&self, file: CreateFile) -> Result<String, InsertError> {
        let key = file.as_s3_key();
        let tmp = file.file;

        let contents = match tmp.read_all().await {
            Ok(contents) => contents,
            Err(e) => return Err(InsertError::Unknown(Box::new(e))),
        };

        // TODO: handle errors
        self.client
            .put_object()
            .bucket(&self.bucket)
            .key(&key)
            .body(ByteStream::from(contents))
            .send()
            .await
            .expect("Create object");

        Ok(key)
    }
}
