use std::fmt::Display;

use actix_web::{
    http::{header::HeaderValue, StatusCode},
    HttpResponse, ResponseError,
};
use manekani_pg::domain::Error as ManekaniDomainError;
use serde::Serialize;

#[derive(Debug, Serialize)]
pub enum Error {
    NotFound,
    Conflict,
    BadRequest,
    #[allow(unused)]
    Forbidden,
    Internal {
        message: String,
    },
}

impl From<ManekaniDomainError> for Error {
    fn from(error: ManekaniDomainError) -> Self {
        match error {
            ManekaniDomainError::Conflict => Self::Conflict,
            ManekaniDomainError::NotFound => Self::NotFound,
            ManekaniDomainError::BadRequest => Self::BadRequest,
            // TODO: use a message for domain internal error too
            ManekaniDomainError::Internal => Self::Internal {
                message: String::from("Something went wrong"),
            },
        }
    }
}

impl ResponseError for Error {
    fn status_code(&self) -> StatusCode {
        match self {
            Self::Internal { .. } => StatusCode::INTERNAL_SERVER_ERROR,
            Self::BadRequest => StatusCode::BAD_REQUEST,
            Self::Conflict => StatusCode::CONFLICT,
            Self::NotFound => StatusCode::NOT_FOUND,
            Self::Forbidden => StatusCode::FORBIDDEN,
        }
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
