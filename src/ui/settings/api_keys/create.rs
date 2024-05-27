use std::sync::Arc;

use askama::Template;
use axum::extract::{Form, State};
use lazy_static::lazy_static;
use manekani_db::Database;
use serde::{Deserialize, Deserializer, Serialize};

use crate::{
    adapter::api_key::{create_api_key, CreateApiKeyRequest},
    domain::api_key::ApiKeyClaims,
};

#[derive(Template)]
#[template(path = "settings/api-keys/create.html")]
pub struct CreateApiKeyTemplate {
    pub permissions: Vec<Permission>,
    pub generated_api_key: Option<String>,
}

#[derive(Debug, Clone, Copy)]
pub struct Permission {
    pub id: &'static str,
    pub name: &'static str,
    pub scopes: Scopes,
}

#[derive(Debug, Clone, Copy)]
pub struct Scopes {
    pub write: Option<bool>,
    pub delete: Option<bool>,
}

lazy_static! {
    pub static ref PERMISSIONS: Vec<Permission> = vec![
        Permission {
            id: "deck",
            name: "Deck",
            scopes: Scopes {
                write: Some(false),
                delete: Some(false)
            }
        },
        Permission {
            id: "subject",
            name: "Subject",
            scopes: Scopes {
                write: Some(false),
                delete: Some(false)
            }
        },
        Permission {
            id: "review",
            name: "Review",
            scopes: Scopes {
                write: Some(false),
                delete: None,
            }
        },
        Permission {
            id: "study_data",
            name: "Study Data",
            scopes: Scopes {
                write: Some(false),
                delete: Some(false)
            }
        }
    ];
}

pub async fn get() -> CreateApiKeyTemplate {
    CreateApiKeyTemplate {
        permissions: PERMISSIONS.clone(),
        generated_api_key: None,
    }
}

#[derive(Debug, Deserialize, Serialize)]
pub struct CreateApiKeyForm {
    pub name: String,
    #[serde(flatten)]
    pub claims: ApiKeyClaimsForm,
}

#[derive(Debug, Deserialize, Serialize, Clone, Copy)]
pub struct ApiKeyClaimsForm {
    #[serde(default)]
    #[serde(deserialize_with = "deserialize_checkbox")]
    pub deck_write: bool,
    #[serde(default)]
    #[serde(deserialize_with = "deserialize_checkbox")]
    pub deck_delete: bool,

    #[serde(default)]
    #[serde(deserialize_with = "deserialize_checkbox")]
    pub subject_write: bool,
    #[serde(default)]
    #[serde(deserialize_with = "deserialize_checkbox")]
    pub subject_delete: bool,

    #[serde(default)]
    #[serde(deserialize_with = "deserialize_checkbox")]
    pub review_create: bool,

    #[serde(default)]
    #[serde(deserialize_with = "deserialize_checkbox")]
    pub study_data_write: bool,
    #[serde(default)]
    #[serde(deserialize_with = "deserialize_checkbox")]
    pub study_data_delete: bool,
}

impl From<ApiKeyClaimsForm> for ApiKeyClaims {
    fn from(value: ApiKeyClaimsForm) -> Self {
        ApiKeyClaims {
            deck_write: value.deck_write,
            deck_delete: value.deck_delete,
            subject_write: value.subject_write,
            subject_delete: value.subject_delete,
            review_create: value.review_create,
            study_data_write: value.study_data_write,
            study_data_delete: value.study_data_delete,
        }
    }
}

#[derive(Template)]
#[template(path = "settings/api-keys/CreateKeyForm.html")]
pub struct CreatedApiKeyTemplate {
    pub permissions: Vec<Permission>,
    pub generated_api_key: Option<String>,
}
pub async fn post(
    State(db): State<Arc<Database>>,
    Form(data): Form<CreateApiKeyForm>,
    // TODO: get user from authenticator
) -> CreatedApiKeyTemplate {
    dbg!(&data.claims);
    let result = create_api_key(
        &db,
        "test_user",
        CreateApiKeyRequest {
            name: data.name,
            claims: data.claims.into(),
        },
    )
    .await
    .unwrap();

    CreatedApiKeyTemplate {
        permissions: PERMISSIONS.clone(),
        generated_api_key: Some(result.generated_key),
    }
}

fn deserialize_checkbox<'de, D>(deserializer: D) -> Result<bool, D::Error>
where
    D: Deserializer<'de>,
{
    let s = String::deserialize(deserializer)?;
    Ok(s == "on")
}
