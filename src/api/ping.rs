use actix_web::{get, HttpResponse, Responder};
use tracing::info;

#[get("/ping")]
pub async fn ping() -> impl Responder {
    info!("ping");
    HttpResponse::Ok().body("pong")
}