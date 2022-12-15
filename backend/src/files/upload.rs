use actix_web::http::StatusCode;
use manekani_service_s3::{domain::create_file, entity::file::Written, S3Repo};
use serde::Serialize;

use crate::error::Error as ApiError;

/// Uploads a file to the s3 repository
pub async fn file(s3: &S3Repo, file: Written) -> Status {
    let name = file.name().to_owned();

    match create_file(s3, file.into()).await {
        Ok(key) => Status::Created(Created {
            code: StatusCode::CREATED.as_u16(),
            field: name,
            key,
        }),
        Err(e) => {
            let error = ApiError::from(e);
            Status::Error(Error {
                code: StatusCode::from(error.kind).as_u16(),
                field: name,
                error,
            })
        }
    }
}

#[derive(Serialize)]
pub enum Status {
    Created(Created),
    Error(Error),
}

#[derive(Serialize)]
pub struct Created {
    pub code: u16,
    pub field: String,
    pub key: String,
}

#[derive(Serialize)]
pub struct Error {
    pub code: u16,
    pub field: String,
    pub error: ApiError,
}

impl From<Result<Created, Error>> for Status {
    fn from(result: Result<Created, Error>) -> Self {
        match result {
            Ok(c) => Self::Created(c),
            Err(e) => Self::Error(e),
        }
    }
}
