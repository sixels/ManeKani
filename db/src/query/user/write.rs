use crate::{model::user::UserModel, Database};

pub struct CreateUser {
    pub email: String,
    pub username: Option<String>,
    pub is_verified: Option<bool>,
    pub is_complete: Option<bool>,
}

#[derive(Debug, thiserror::Error)]
pub enum CreateUserError {
    #[error("a user with the email already exists")]
    DuplicateEmail,
    #[error("a user with the username already exists")]
    DuplicateUsername,
    #[error("database error")]
    DatabaseError(#[source] sqlx::Error),
}

impl From<sqlx::Error> for CreateUserError {
    fn from(err: sqlx::Error) -> Self {
        let specific_err = match err {
            sqlx::Error::Database(ref db_err) => match db_err.code() {
                Some(code) => {
                    if code == "23505" {
                        dbg!(db_err.constraint());

                        if db_err.constraint() == Some("users_email_key") {
                            Some(CreateUserError::DuplicateEmail)
                        } else {
                            Some(CreateUserError::DuplicateUsername)
                        }
                    } else {
                        None
                    }
                }
                _ => None,
            },
            _ => None,
        };
        specific_err.unwrap_or(CreateUserError::DatabaseError(err))
    }
}

pub async fn create_user(db: &Database, user: CreateUser) -> Result<UserModel, CreateUserError> {
    let result = sqlx::query!(
        r#"
        INSERT INTO users (email, username, is_verified, is_complete)
        VALUES ($1, $2, $3, $4)
        RETURNING *
        "#,
        user.email,
        user.username,
        user.is_verified,
        user.is_complete,
    )
    .fetch_one(db)
    .await?;

    Ok(UserModel {
        id: result.id,
        username: result.username,
        email: result.email,
        is_verified: result.is_verified,
        is_complete: result.is_complete,
        created_at: result.created_at,
        updated_at: result.updated_at,
        deck_ids: None,
    })
}
