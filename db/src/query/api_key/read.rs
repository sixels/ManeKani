use uuid::Uuid;

use crate::{model::api_key::ApiKeyModel, Database};

#[derive(Debug, thiserror::Error)]
pub enum GetApiKeyError {
    #[error("api key not found")]
    NotFound,
    #[error("database error")]
    DatabaseError(#[from] sqlx::Error),
}

pub async fn get_api_key(db: &Database, id: Uuid) -> Result<ApiKeyModel, GetApiKeyError> {
    let result = sqlx::query_as!(
        ApiKeyModel,
        r#"
        SELECT
            id,
            name,
            prefix,
            claims,
            created_at,
            updated_at,
            revoked_at,
            used_at,
            created_by_user_id
        FROM api_keys WHERE id = $1
        "#,
        id,
    )
    .fetch_optional(&db.pool)
    .await?;

    result.ok_or(GetApiKeyError::NotFound)
}

#[derive(Debug, thiserror::Error)]
pub enum GetUserApiKeysError {
    #[error("database error")]
    DatabaseError(#[from] sqlx::Error),
}

pub async fn get_user_api_keys(
    db: &Database,
    user: String,
) -> Result<Vec<ApiKeyModel>, GetUserApiKeysError> {
    let result = sqlx::query_as!(
        ApiKeyModel,
        r#"
        SELECT
            id,
            name,
            prefix,
            claims,
            created_at,
            updated_at,
            revoked_at,
            used_at,
            created_by_user_id
        FROM api_keys WHERE created_by_user_id = $1
        ORDER BY created_at DESC
        "#,
        user,
    )
    .fetch_all(&db.pool)
    .await?;

    Ok(result)
}

#[derive(Debug, thiserror::Error)]
pub enum CountUserApiKeysError {
    #[error("database error")]
    DatabaseError(#[from] sqlx::Error),
}

pub async fn count_user_api_keys(
    db: &Database,
    user_id: &str,
) -> Result<usize, CountUserApiKeysError> {
    let result = sqlx::query!(
        r#"
        SELECT COUNT(*)
        FROM api_keys WHERE created_by_user_id = $1
        "#,
        user_id,
    )
    .fetch_one(&db.pool)
    .await?;

    Ok(result.count.unwrap_or(0) as usize)
}
