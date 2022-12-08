use manekani_types::repository::{InsertError, QueryError};

#[derive(Debug)]
pub enum Error {
    Conflict,
    NotFound,
    BadRequest,
    Internal,
}

impl From<QueryError> for Error {
    fn from(error: QueryError) -> Self {
        match error {
            QueryError::NotFound => Self::NotFound,
            QueryError::Unknown => Self::Internal,
        }
    }
}

impl From<InsertError> for Error {
    fn from(error: InsertError) -> Self {
        match error {
            InsertError::Conflict => Self::Conflict,
            InsertError::BadRequest => Self::BadRequest,
            InsertError::Unknown => Self::Internal,
        }
    }
}
