use std::sync::Arc;

use actix_web::{get, post, web, HttpResponse, Responder};
use tracing::{debug, info};

use crate::{
    api::state::State,
    domain::{create_radical, get_radical},
    entities::radical::{GetRadical, InsertRadical},
};

#[get("{radical}")]
pub async fn get(radical: web::Path<String>, state: web::Data<Arc<State>>) -> impl Responder {
    let symbol = radical.into_inner();
    info!("Getting radical '{symbol}'");

    let mut conn = state
        .db
        .acquire()
        .await
        .expect("Could not get a database connection");

    let radical = get_radical::execute(&mut conn, &GetRadical { symbol })
        .await
        .unwrap();

    HttpResponse::Ok().json(radical)
}

#[post("")]
pub async fn create(req: web::Json<InsertRadical>, state: web::Data<Arc<State>>) -> impl Responder {
    info!(event = "Creating radical", radical_name = &req.name);

    let mut conn = state
        .db
        .acquire()
        .await
        .expect("Could not get a database connection");

    let created = create_radical::execute(&mut conn, &req.0).await.unwrap();
    debug!(
        event = "Created radical",
        radical_id = created.id.to_string()
    );
    HttpResponse::Ok().json(created)
}
