use std::error::Error as StdError;

use manekani_types::repository::{error::UpdateError, InsertError, QueryError};

#[derive(Debug)]
pub enum Error {
    Conflict,
    NotFound,
    BadRequest,
    Internal(Box<dyn StdError>),
}

impl From<QueryError> for Error {
    fn from(error: QueryError) -> Self {
        match error {
            QueryError::NotFound => Self::NotFound,
            QueryError::Unknown(e) => Self::Internal(e),
        }
    }
}

impl From<InsertError> for Error {
    fn from(error: InsertError) -> Self {
        match error {
            InsertError::Conflict => Self::Conflict,
            InsertError::BadRequest => Self::BadRequest,
            InsertError::Unknown(e) => Self::Internal(e),
        }
    }
}

impl From<UpdateError> for Error {
    fn from(error: UpdateError) -> Self {
        match error {
            UpdateError::NotFound => Self::NotFound,
            UpdateError::BadRequest => Self::BadRequest,
            UpdateError::Unknown(e) => Self::Internal(e),
        }
    }
}
