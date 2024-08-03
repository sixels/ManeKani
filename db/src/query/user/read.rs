use uuid::Uuid;

use crate::{model::user::UserModel, Database};

pub async fn get_user(db: &Database, user_id: Uuid) -> Result<UserModel, sqlx::Error> {
    let user = sqlx::query_as!(
        UserModel,
        r#"
        SELECT
            users.id,
            users.email,
            users.username,
            users.is_verified,
            users.is_complete,
            users.created_at,
            users.updated_at,
            ARRAY_REMOVE(ARRAY_AGG(decks.id), NULL) deck_ids
        FROM users
            LEFT JOIN decks ON users.id = decks.created_by_user_id
        WHERE users.id = $1
        GROUP BY users.id
        "#,
        user_id
    )
    .fetch_one(db)
    .await?;

    Ok(user)
}
