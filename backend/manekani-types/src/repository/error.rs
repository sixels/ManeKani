#[derive(Debug)]
pub enum QueryError {
    NotFound,
    Unknown,
}

#[derive(Debug)]
pub enum QueryAllError {
    Unknown,
}

#[derive(Debug)]
pub enum InsertError {
    Conflict,
    BadRequest,
    Unknown,
}

#[derive(Debug)]
pub enum DeleteError {
    NotFound,
    Unknown,
}

impl From<sqlx::Error> for QueryError {
    fn from(error: sqlx::Error) -> Self {
        match error {
            sqlx::Error::RowNotFound => Self::NotFound,
            _ => Self::Unknown,
        }
    }
}

impl From<sqlx::Error> for QueryAllError {
    fn from(_: sqlx::Error) -> Self {
        Self::Unknown
    }
}

impl From<sqlx::Error> for InsertError {
    fn from(error: sqlx::Error) -> Self {
        match error {
            sqlx::Error::Database(d) if d.code() == Some("23505".into()) => Self::Conflict,
            _ => Self::Unknown,
        }
    }
}

impl From<sqlx::Error> for DeleteError {
    fn from(error: sqlx::Error) -> Self {
        match error {
            sqlx::Error::RowNotFound => Self::NotFound,
            _ => Self::Unknown,
        }
    }
}
