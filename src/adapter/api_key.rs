pub mod error;
mod token;

use manekani_db::{model::api_key::ApiKeyModel, query::api_key, Database};
use uuid::Uuid;

use crate::domain::api_key::{ApiKey, ApiKeyClaims};

use self::{
    error::{CreateApiKeyError, GetApiKeyError, GetApiKeysError},
    token::generate_random_token,
};

impl From<ApiKeyModel> for ApiKey {
    fn from(value: ApiKeyModel) -> Self {
        ApiKey {
            id: value.id.to_string(),
            name: value.name,
            prefix: value.prefix,
            // maybe add an error check if the claims change in the future; or add aliases to serde
            claims: serde_json::from_value(value.claims).unwrap(),
            used_at: value.used_at,
            revoked_at: value.revoked_at,
            created_at: value.created_at,
            updated_at: value.updated_at,
        }
    }
}

pub struct GetApiKeyResponse(pub ApiKey);

pub async fn get_api_key(db: &Database, id: Uuid) -> Result<GetApiKeyResponse, GetApiKeyError> {
    let result = api_key::read::get_api_key(db, id).await?;
    Ok(GetApiKeyResponse(result.into()))
}

pub struct GetApiKeysResponse(pub Vec<ApiKey>);

pub async fn get_user_api_keys(
    db: &Database,
    user_id: Uuid,
) -> Result<GetApiKeysResponse, GetApiKeysError> {
    let result = api_key::read::get_user_api_keys(db, user_id).await?;
    Ok(GetApiKeysResponse(
        result.into_iter().map(ApiKey::from).collect(),
    ))
}

pub struct CreateApiKeyRequest {
    pub name: String,
    pub claims: ApiKeyClaims,
}

pub struct CreateApiKeyResponse {
    pub generated_key: String,
    pub api_key: ApiKey,
}

pub async fn create_api_key(
    db: &Database,
    user_id: Uuid,
    data: CreateApiKeyRequest,
) -> Result<CreateApiKeyResponse, CreateApiKeyError> {
    if api_key::read::count_user_api_keys(db, user_id).await? >= 10 {
        return Err(CreateApiKeyError::LimitExceeded(10));
    }

    let token = generate_random_token();
    let json_claims =
        serde_json::to_value(&data.claims).map_err(|_| CreateApiKeyError::InternalError)?;

    let result = api_key::write::create_api_key(
        &db,
        api_key::write::CreateApiKey {
            name: data.name,
            prefix: token.prefix,
            token: token.hash,
            claims: json_claims,
            user_id,
        },
    )
    .await?;

    Ok(CreateApiKeyResponse {
        generated_key: token.key,
        api_key: result.into(),
    })
}
