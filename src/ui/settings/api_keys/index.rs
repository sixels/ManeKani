use std::sync::Arc;

use askama::Template;
use axum::extract::State;
use manekani_db::{model::api_key::ApiKeyModel, Database};
use serde::{Deserialize, Serialize};

use crate::{adapter::api_key::get_user_api_keys, domain::api_key::ApiKey, util::time::HumanTime};

#[derive(Debug)]
pub struct ApiKeyInfo {
    pub id: String,
    pub name: String,
    pub prefix: String,
    pub used_at: Option<HumanTime>,
    pub revoked_at: Option<HumanTime>,
    pub created_at: HumanTime,
    pub updated_at: HumanTime,
}

impl From<ApiKeyModel> for ApiKeyInfo {
    fn from(value: ApiKeyModel) -> Self {
        ApiKeyInfo {
            id: value.id.to_string(),
            name: value.name,
            prefix: value.prefix,
            used_at: value.used_at.map(HumanTime::from),
            revoked_at: value.revoked_at.map(HumanTime::from),
            updated_at: HumanTime::from(value.updated_at),
            created_at: HumanTime::from(value.created_at),
        }
    }
}

#[derive(Template)]
#[template(path = "settings/api_keys.html")]
pub struct ApiKeysTemplate {
    pub key_infos: Vec<ApiKeyInfo>,
}

impl From<ApiKey> for ApiKeyInfo {
    fn from(value: ApiKey) -> Self {
        ApiKeyInfo {
            id: value.id.to_string(),
            name: value.name,
            prefix: value.prefix,
            used_at: value.used_at.map(HumanTime::from),
            revoked_at: value.revoked_at.map(HumanTime::from),
            updated_at: HumanTime::from(value.updated_at),
            created_at: HumanTime::from(value.created_at),
        }
    }
}

pub async fn get(State(db): State<Arc<Database>>) -> ApiKeysTemplate {
    let api_keys = get_user_api_keys(&db, String::from("test_user"))
        .await
        .unwrap();

    let api_keys = api_keys.0.into_iter().map(ApiKeyInfo::from).collect();

    println!("{:?}", api_keys);

    ApiKeysTemplate {
        key_infos: api_keys,
    }
}

#[derive(Debug, Deserialize, Serialize)]
pub struct CreateApiKeyRequest {
    pub name: String,
    pub claims: ApiKeyClaims,
}

#[derive(Debug, Deserialize, Serialize, Clone, Copy)]
pub struct ApiKeyClaims {
    pub deck_write: bool,
    pub deck_delete: bool,

    pub subject_write: bool,
    pub subject_delete: bool,

    pub review_create: bool,

    pub study_data_write: bool,
    pub study_data_delete: bool,
}
