use crate::{model::user::UserModel, Database};

pub struct CreateUser {
    pub id: String,
    pub username: Option<String>,
    pub is_verified: Option<bool>,
    pub is_complete: Option<bool>,
}

pub async fn create_user(db: &Database, user: CreateUser) -> Result<UserModel, sqlx::Error> {
    let result = sqlx::query!(
        r#"
        INSERT INTO users (id, username, is_verified, is_complete)
        VALUES ($1, $2, $3, $4)
        RETURNING *
        "#,
        user.id,
        user.username,
        user.is_verified,
        user.is_complete,
    )
    .fetch_one(db)
    .await?;

    Ok(UserModel {
        id: result.id,
        username: result.username,
        is_verified: result.is_verified,
        is_complete: result.is_complete,
        created_at: result.created_at,
        updated_at: result.updated_at,
        deck_ids: None,
    })
}
