use std::sync::Arc;

use askama::Template;
use axum::{
    extract::{Path, State},
    http::{HeaderMap, StatusCode},
};
use manekani_db::{query::api_key, Database};
use uuid::Uuid;

use crate::{adapter::api_key::get_api_key, domain::api_key::ApiKey};

use super::create::{Permission, Scopes};

#[derive(Template)]
#[template(path = "settings/api-keys/manage.html")]
pub struct ManageApiKeyTemplate {
    api_key: ApiKey,
    permissions: Vec<Permission>,
}

pub async fn get(
    State(db): State<Arc<Database>>,
    Path(api_key_id): Path<Uuid>,
) -> ManageApiKeyTemplate {
    let api_key = get_api_key(&db, api_key_id).await.unwrap();

    ManageApiKeyTemplate {
        permissions: vec![
            Permission {
                id: "deck",
                name: "Deck",
                scopes: Scopes {
                    write: Some(api_key.0.claims.deck_write),
                    delete: Some(api_key.0.claims.deck_delete),
                },
            },
            Permission {
                id: "subject",
                name: "Subject",
                scopes: Scopes {
                    write: Some(api_key.0.claims.subject_write),
                    delete: Some(api_key.0.claims.subject_delete),
                },
            },
            Permission {
                id: "review",
                name: "Review",
                scopes: Scopes {
                    write: Some(api_key.0.claims.review_create),
                    delete: None,
                },
            },
            Permission {
                id: "study_data",
                name: "Study Data",
                scopes: Scopes {
                    write: Some(api_key.0.claims.study_data_write),
                    delete: Some(api_key.0.claims.study_data_delete),
                },
            },
        ],
        api_key: api_key.0,
    }
}

pub async fn delete(
    State(db): State<Arc<Database>>,
    Path(api_key_id): Path<Uuid>,
) -> (HeaderMap, StatusCode) {
    // TODO: check if user owns the token
    let _ = api_key::write::delete_api_key(&db, api_key_id)
        .await
        .unwrap();

    let mut headers = HeaderMap::new();
    headers.insert("HX-Redirect", "/settings/api-keys".parse().unwrap());

    (headers, StatusCode::OK)
}
