use std::fmt::Debug;

use actix_web::{get, web};
use askama::Template;
use manekani_db::{model::api_key::ApiKeyModel, Database};

use crate::{
    adapter::api_key::{error::GetApiKeysError, get_user_api_keys},
    domain::api_key::ApiKey,
    misc::time::HumanTime,
    routing::auth::extractor::CurrentUser,
};

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

#[get("")]
pub async fn get(
    user: CurrentUser,
    db: web::Data<Database>,
) -> Result<ApiKeysTemplate, actix_web::Error> {
    let api_keys = match get_user_api_keys(&db, user.id).await {
        Ok(api_keys) => api_keys,
        Err(err) => match err {
            GetApiKeysError::InternalError => {
                tracing::error!("failed to get user's api keys: {:?}", err);
                return Err(actix_web::error::ErrorInternalServerError(
                    "failed to get user's api key",
                ));
            }
        },
    };

    let api_keys = api_keys.0.into_iter().map(ApiKeyInfo::from).collect();

    println!("{:?}", api_keys);

    Ok(ApiKeysTemplate {
        key_infos: api_keys,
    })
}
