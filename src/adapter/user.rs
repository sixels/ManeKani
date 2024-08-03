pub mod error;

use error::CreateUserError;
use manekani_db::{model::user::UserModel, query::user, Database};

use crate::domain::user::User;

#[derive(Debug, Default)]
pub struct CreateUserRequest {
    pub email: String,
    pub username: Option<String>,
    pub is_complete: Option<bool>,
    pub is_verified: Option<bool>,
}

pub struct CreateUserResponse(pub User);

pub async fn create_user(
    db: &Database,
    data: CreateUserRequest,
) -> Result<CreateUserResponse, CreateUserError> {
    let result = user::write::create_user(
        db,
        user::write::CreateUser {
            email: data.email,
            username: data.username,
            is_complete: Some(false),
            is_verified: Some(false),
        },
    )
    .await?;

    Ok(CreateUserResponse(result.into()))
}

impl From<UserModel> for User {
    fn from(user: UserModel) -> Self {
        User {
            id: user.id,
            email: user.email,
            username: user.username,
            is_verified: user.is_verified,
            is_complete: user.is_complete,
            created_at: user.created_at,
            updated_at: user.updated_at,
        }
    }
}
