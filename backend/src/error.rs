use std::fmt::Display;

use actix_web::{
    http::{header::HeaderValue, StatusCode},
    HttpResponse, ResponseError,
};
use manekani_service_common::repository::{
    error::{Error as ServiceError, UpdateError},
    DeleteError, InsertError, QueryAllError, QueryError,
};
use serde::Serialize;

#[derive(Serialize, Debug)]
pub struct Error {
    #[serde(skip_serializing)]
    kind: ErrorKind,
    pub message: String,
}

#[derive(Debug, Serialize, Clone, Copy)]
pub enum ErrorKind {
    NotFound,
    Conflict,
    BadRequest,
    #[allow(unused)]
    Forbidden,
    Internal,
}

impl Error {
    pub fn kind(&self) -> ErrorKind {
        self.kind
    }
}

impl From<ServiceError> for Error {
    fn from(error: ServiceError) -> Self {
        Self {
            message: error.to_string(),
            kind: ErrorKind::from(error),
        }
    }
}

impl From<ServiceError> for ErrorKind {
    fn from(error: ServiceError) -> Self {
        match error {
            ServiceError::Insert(InsertError::Conflict) => Self::Conflict,

            ServiceError::Delete(DeleteError::NotFound)
            | ServiceError::Query(QueryError::NotFound)
            | ServiceError::Update(UpdateError::NotFound) => Self::NotFound,

            ServiceError::Insert(InsertError::BadRequest)
            | ServiceError::Update(UpdateError::BadRequest) => Self::BadRequest,

            ServiceError::Delete(DeleteError::Unknown(_))
            | ServiceError::Insert(InsertError::Unknown(_))
            | ServiceError::Query(QueryError::Unknown(_))
            | ServiceError::QueryAll(QueryAllError::Unknown(_))
            | ServiceError::Update(UpdateError::Unknown(_)) => Self::Internal,
        }
    }
}

impl From<ErrorKind> for StatusCode {
    fn from(error: ErrorKind) -> Self {
        match error {
            ErrorKind::Internal => StatusCode::INTERNAL_SERVER_ERROR,
            ErrorKind::BadRequest => StatusCode::BAD_REQUEST,
            ErrorKind::Conflict => StatusCode::CONFLICT,
            ErrorKind::NotFound => StatusCode::NOT_FOUND,
            ErrorKind::Forbidden => StatusCode::FORBIDDEN,
        }
    }
}

impl ResponseError for Error {
    fn status_code(&self) -> StatusCode {
        self.kind.into()
    }

    fn error_response(&self) -> HttpResponse<actix_web::body::BoxBody> {
        let mut res = HttpResponse::new(self.status_code());

        let mut buf = actix_web::web::BytesMut::new();
        use std::fmt::Write;
        let _ = write!(&mut buf, "{}", self);

        let mime = HeaderValue::from_static("application/json");
        res.headers_mut()
            .insert(actix_web::http::header::CONTENT_TYPE, mime);

        res.set_body(actix_web::body::BoxBody::new(buf))
    }
}

impl Display for Error {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", serde_json::to_string(self).unwrap())
    }
}
