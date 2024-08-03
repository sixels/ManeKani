use uuid::Uuid;

use crate::{model::api_key::ApiKeyModel, Database};

pub struct CreateApiKey {
    pub name: String,
    pub token: String,
    pub prefix: String,
    pub claims: serde_json::Value,
    pub user_id: Uuid,
}

#[derive(Debug, thiserror::Error)]
pub enum CreateApiKeyError {
    #[error("an api key with the same name already exists")]
    DuplicateApiKey,
    #[error("database error")]
    DatabaseError(#[source] sqlx::Error),
}

impl From<sqlx::Error> for CreateApiKeyError {
    fn from(err: sqlx::Error) -> Self {
        let specific_err = match err {
            sqlx::Error::Database(ref db_err) => match db_err.code() {
                Some(code) => {
                    if code == "23505" {
                        Some(CreateApiKeyError::DuplicateApiKey)
                    } else {
                        None
                    }
                }
                _ => None,
            },
            _ => None,
        };
        specific_err.unwrap_or(CreateApiKeyError::DatabaseError(err))
    }
}

pub async fn create_api_key(
    db: &Database,
    data: CreateApiKey,
) -> Result<ApiKeyModel, CreateApiKeyError> {
    let result = sqlx::query_as!(
        ApiKeyModel,
        r#"
        INSERT INTO api_keys (name, token, prefix, claims, created_by_user_id)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING
            id,
            name,
            prefix,
            claims,
            created_at,
            updated_at,
            revoked_at,
            used_at,
            created_by_user_id
        "#,
        data.name,
        data.token,
        data.prefix,
        data.claims,
        &data.user_id,
    )
    .fetch_one(db)
    .await?;

    Ok(result)
}

#[derive(Debug, thiserror::Error)]
pub enum DeleteApiKeyError {
    #[error("api key not found")]
    NotFound,
    #[error("database error")]
    DatabaseError(#[from] sqlx::Error),
}

pub async fn delete_api_key(db: &Database, id: Uuid) -> Result<ApiKeyModel, DeleteApiKeyError> {
    let result = sqlx::query_as!(
        ApiKeyModel,
        r#"
        DELETE FROM api_keys
        WHERE id = $1
        RETURNING
            id,
            name,
            prefix,
            claims,
            created_at,
            updated_at,
            revoked_at,
            used_at,
            created_by_user_id
        "#,
        id,
    )
    .fetch_optional(db)
    .await?;

    result.ok_or(DeleteApiKeyError::NotFound)
}
