use std::error::Error as StdError;

use thiserror::Error as ThisError;

#[derive(ThisError, Debug)]
pub enum Error {
    #[error("query error: {0}")]
    Query(#[from] QueryError),
    #[error("query_all error: {0}")]
    QueryAll(#[from] QueryAllError),
    #[error("insert error: {0}")]
    Insert(#[from] InsertError),
    #[error("delete error: {0}")]
    Delete(#[from] DeleteError),
    #[error("update error: {0}")]
    Update(#[from] UpdateError),
}

#[derive(ThisError, Debug)]
pub enum QueryError {
    #[error("could not find any item")]
    NotFound,
    #[error(transparent)]
    Unknown(#[from] Box<dyn StdError>),
}

#[derive(ThisError, Debug)]
pub enum QueryAllError {
    #[error(transparent)]
    Unknown(#[from] Box<dyn StdError>),
}

#[derive(ThisError, Debug)]
pub enum InsertError {
    #[error("the item conflicts with an already existing one")]
    Conflict,
    #[error("the item have an invalid data")]
    BadRequest,
    #[error(transparent)]
    Unknown(#[from] Box<dyn StdError>),
}

#[derive(ThisError, Debug)]
pub enum DeleteError {
    #[error("the item could not exist")]
    NotFound,
    #[error(transparent)]
    Unknown(#[from] Box<dyn StdError>),
}

#[derive(ThisError, Debug)]
pub enum UpdateError {
    #[error("the item have an invalid data")]
    BadRequest,
    #[error("the item could not be found")]
    NotFound,
    #[error(transparent)]
    Unknown(#[from] Box<dyn StdError>),
}

#[cfg(feature = "sqlx-errors")]
mod sqlx_errors {
    use super::{DeleteError, InsertError, QueryAllError, QueryError, UpdateError};

    impl From<sqlx::Error> for QueryError {
        fn from(error: sqlx::Error) -> Self {
            match error {
                sqlx::Error::RowNotFound => Self::NotFound,
                e => Self::Unknown(Box::new(e)),
            }
        }
    }

    impl From<sqlx::Error> for QueryAllError {
        fn from(e: sqlx::Error) -> Self {
            Self::Unknown(Box::new(e))
        }
    }

    impl From<sqlx::Error> for InsertError {
        fn from(error: sqlx::Error) -> Self {
            match error {
                sqlx::Error::Database(d) if d.code() == Some("23505".into()) => Self::Conflict,
                e => Self::Unknown(Box::new(e)),
            }
        }
    }

    impl From<sqlx::Error> for DeleteError {
        fn from(error: sqlx::Error) -> Self {
            match error {
                sqlx::Error::RowNotFound => Self::NotFound,
                e => Self::Unknown(Box::new(e)),
            }
        }
    }

    impl From<sqlx::Error> for UpdateError {
        fn from(error: sqlx::Error) -> Self {
            match error {
                sqlx::Error::RowNotFound => Self::NotFound,
                e => Self::Unknown(Box::new(e)),
            }
        }
    }
}

#[cfg(feature = "s3-errors")]
mod s3_errors {
    use super::{InsertError, QueryError};

    use aws_sdk_s3::{
        error::{GetObjectError, GetObjectErrorKind, PutObjectError},
        types::SdkError,
    };

    impl From<SdkError<GetObjectError>> for QueryError {
        fn from(error: SdkError<GetObjectError>) -> Self {
            if let SdkError::ServiceError { err, .. } = &error {
                if let GetObjectErrorKind::NoSuchKey(_)
                | GetObjectErrorKind::InvalidObjectState(_) = err.kind
                {
                    return Self::NotFound;
                }
            };
            Self::Unknown(Box::new(error))
        }
    }

    impl From<SdkError<PutObjectError>> for InsertError {
        fn from(error: SdkError<PutObjectError>) -> Self {
            // if let SdkError::ServiceError { err, .. } = &error {
            //     if let PutObjectErrorKind::Unhandled(e) = err.kind {
            //         return Self(InsertError::Unknown(e));
            //     }
            // };
            Self::Unknown(Box::new(error))
        }
    }
}
