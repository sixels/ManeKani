use actix_web::{delete, get, web, HttpResponse};
use askama::Template;

use manekani_db::query::api_key;
use uuid::Uuid;

use crate::{adapter::api_key::get_api_key, domain::api_key::ApiKey, state::AppState};

use super::create::{Permission, Scopes};

#[derive(Template)]
#[template(path = "settings/api-keys/manage.html")]
pub struct ManageApiKeyTemplate {
    api_key: ApiKey,
    permissions: Vec<Permission>,
}

#[get("/manage")]
pub async fn get(state: web::Data<AppState>, api_key_id: web::Path<Uuid>) -> ManageApiKeyTemplate {
    let api_key = get_api_key(&state.db, *api_key_id).await.unwrap();

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

#[delete("/manage")]
pub async fn delete(state: web::Data<AppState>, api_key_id: web::Path<Uuid>) -> HttpResponse {
    // TODO: check if user owns the token
    let _ = api_key::write::delete_api_key(&state.db, *api_key_id)
        .await
        .unwrap();

    HttpResponse::Ok()
        .insert_header(("HX-Redirect", "/settings/api-keys"))
        .finish()
}
