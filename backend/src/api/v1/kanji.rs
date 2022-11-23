use std::sync::Arc;

use actix_web::{get, post, web, HttpResponse, Responder};
use tracing::{debug, info};

use crate::{
    api::state::State,
    domain::create_kanji,
    entities::kanji::{InsertKanji, Kanji},
};

#[get("{kanji}")]
pub async fn get(kanji: web::Path<String>, state: web::Data<Arc<State>>) -> impl Responder {
    let kanji = kanji.into_inner();
    info!(event = "getting kanji", kanji = kanji);

    let mut conn = state
        .db
        .acquire()
        .await
        .expect("Could not get a database connection");

    // temporary
    let Ok(kanji) = sqlx::query_as!(Kanji, "SELECT * FROM kanjis WHERE symbol = $1", kanji)
        .fetch_one(&mut conn)
        .await
        else {
            return HttpResponse::InternalServerError().json("sorry");
        };

    HttpResponse::Ok().json(kanji)
}

#[post("")]
pub async fn create(req: web::Json<InsertKanji>, state: web::Data<Arc<State>>) -> impl Responder {
    info!(event = "Creating kanji", kanji_name = &req.name);

    let mut conn = state
        .db
        .acquire()
        .await
        .expect("Could not get a database connection");

    let Ok(created) = create_kanji::execute(&mut conn, &req.0).await else {
        return HttpResponse::InternalServerError().json("sorry");
    };
    debug!(event = "Created kanji", kanji_id = created.id.to_string());
    HttpResponse::Ok().json(created)
}
