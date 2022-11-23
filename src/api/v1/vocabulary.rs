use std::sync::Arc;

use actix_web::{get, post, web, HttpResponse, Responder};
use tracing::{debug, info};

use crate::{
    api::state::State,
    domain::{create_vocabulary, get_vocabulary},
    entities::vocabulary::{GetVocabulary, InsertVocabulary},
};

#[get("{vocabulary}")]
pub async fn get(vocabulary: web::Path<String>, state: web::Data<Arc<State>>) -> impl Responder {
    let word = vocabulary.into_inner();
    info!("Getting vocabulary '{word}'");

    let mut conn = state
        .db
        .acquire()
        .await
        .expect("Could not get a database connection");

    let Ok(vocabulary) = get_vocabulary::execute(&mut conn, &GetVocabulary { word })
        .await
        else {
            return HttpResponse::InternalServerError().json("sorry");
        };

    HttpResponse::Ok().json(vocabulary)
}

#[post("")]
pub async fn create(
    req: web::Json<InsertVocabulary>,
    state: web::Data<Arc<State>>,
) -> impl Responder {
    info!(event = "Creating vocabulary", vocabulary_name = &req.name);

    let mut conn = state
        .db
        .acquire()
        .await
        .expect("Could not get a database connection");

    let Ok(created) = create_vocabulary::execute(&mut conn, &req.0).await else {
        return HttpResponse::InternalServerError().json("sorry");
    };
    debug!(
        event = "Created vocabulary",
        vocabulary_id = created.id.to_string()
    );
    HttpResponse::Ok().json(created)
}
