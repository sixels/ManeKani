use actix_web::http::StatusCode;
use manekani_service_s3::{domain::create_file, entity::file::WrittenFile, S3Repo};
use serde::Serialize;

use crate::api::error::Error;

pub async fn upload_file(s3: &S3Repo, file: WrittenFile) -> UploadStatus {
    let name = file.name().to_owned();

    match create_file(s3, file.into()).await {
        Ok(key) => UploadStatus::Created(UploadCreated {
            code: StatusCode::CREATED.as_u16(),
            name,
            key,
        }),
        Err(e) => {
            let error = Error::from(e);
            UploadStatus::Error(UploadError {
                code: StatusCode::from(error.kind()).as_u16(),
                name,
                error,
            })
        }
    }
}

#[derive(Serialize)]
pub enum UploadStatus {
    Created(UploadCreated),
    Error(UploadError),
}
#[derive(Serialize)]
pub struct UploadCreated {
    pub code: u16,
    pub name: String,
    pub key: String,
}
#[derive(Serialize)]
pub struct UploadError {
    pub code: u16,
    pub name: String,
    pub error: Error,
}

impl From<Result<UploadCreated, UploadError>> for UploadStatus {
    fn from(result: Result<UploadCreated, UploadError>) -> Self {
        match result {
            Ok(c) => Self::Created(c),
            Err(e) => Self::Error(e),
        }
    }
}
