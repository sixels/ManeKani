use sqlx::error::DatabaseError;

pub mod kanji;
pub mod radical;
pub mod vocabulary;

#[derive(thiserror::Error, Debug)]
pub enum RepositoryError {
    #[error("request conflicts with an already existing data")]
    Conflict,
    #[error("invalid request format: {0}")]
    BadRequest(String),
    #[error(transparent)]
    Database(#[from] Box<dyn DatabaseError>),
    #[error(transparent)]
    Unknown(#[from] Box<dyn std::error::Error>),
}

impl From<sqlx::Error> for RepositoryError {
    fn from(sqlx_error: sqlx::Error) -> Self {
        match sqlx_error {
            sqlx::Error::Database(d) if d.code() == Some("23505".into()) => {
                RepositoryError::Conflict
            }
            sqlx::Error::Database(db_error) => RepositoryError::Database(db_error),
            _ => RepositoryError::Unknown(Box::new(sqlx_error)),
        }
    }
}
