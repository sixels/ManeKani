use std::error::Error;

#[derive(Debug)]
pub enum QueryError {
    NotFound,
    Unknown(Box<dyn Error>),
}

#[derive(Debug)]
pub enum QueryAllError {
    Unknown(Box<dyn Error>),
}

#[derive(Debug)]
pub enum InsertError {
    Conflict,
    BadRequest,
    Unknown(Box<dyn Error>),
}

#[derive(Debug)]
pub enum DeleteError {
    NotFound,
    Unknown(Box<dyn Error>),
}

#[derive(Debug)]
pub enum UpdateError {
    BadRequest,
    NotFound,
    Unknown(Box<dyn Error>),
}

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
