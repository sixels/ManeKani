use std::fmt::Display;

use actix_web::{
    http::{header::HeaderValue, StatusCode},
    HttpResponse, ResponseError,
};
use manekani_pg::domain::Error as ManekaniDomainError;
use manekani_s3::domain::Error as S3DomainError;
use serde::Serialize;

#[derive(Serialize, Debug)]
pub struct Error {
    #[serde(skip_serializing)]
    kind: ErrorKind,
    pub message: String,
}

impl Error {
    pub fn kind(&self) -> &ErrorKind {
        &self.kind
    }

    pub fn conflict() -> Self {
        Self {
            kind: ErrorKind::Conflict,
            message: String::from("Request conflicts with an already existing item"),
        }
    }

    pub fn not_found() -> Self {
        Self {
            kind: ErrorKind::NotFound,
            message: String::from("Request item not found"),
        }
    }

    pub fn bad_request() -> Self {
        Self {
            kind: ErrorKind::BadRequest,
            message: String::from("Request item is invalid"),
        }
    }

    pub fn internal<M: Into<String>>(message: M) -> Self {
        Self {
            kind: ErrorKind::Internal,
            message: message.into(),
        }
    }
}

#[derive(Debug, Serialize)]
pub enum ErrorKind {
    NotFound,
    Conflict,
    BadRequest,
    #[allow(unused)]
    Forbidden,
    Internal,
}

impl From<ManekaniDomainError> for Error {
    fn from(error: ManekaniDomainError) -> Self {
        match error {
            ManekaniDomainError::Conflict => Self::conflict(),
            ManekaniDomainError::NotFound => Self::not_found(),
            ManekaniDomainError::BadRequest => Self::bad_request(),
            // TODO: use a message for domain internal error too
            ManekaniDomainError::Internal(e) => {
                Self::internal(format!("Something went wrong: {e:?}"))
            }
        }
    }
}

impl From<S3DomainError> for Error {
    fn from(error: S3DomainError) -> Self {
        match error {
            S3DomainError::Conflict => Self::conflict(),
            S3DomainError::NotFound => Self::not_found(),
            S3DomainError::BadRequest => Self::bad_request(),
            // TODO: use a message for domain internal error too
            S3DomainError::Internal => Self::internal("Something went wrong"),
        }
    }
}

impl From<&ErrorKind> for StatusCode {
    fn from(error: &ErrorKind) -> Self {
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
        StatusCode::from(self.kind())
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
