use std::sync::Arc;

use actix_web::{get, post, web, HttpResponse};
use manekani_pg::{
    domain::radical::{insert, query, query_by_kanji},
    entity::{GetKanji, GetRadical, InsertRadical},
};
use tracing::{debug, info};

use crate::api::{error::Error as ApiError, state::State};

#[get("{radical}")]
pub async fn get(
    radical_name: web::Path<String>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let name = radical_name.into_inner();
    let radical = GetRadical { name };

    info!("Querying radical: {}", radical.name);
    let radical = query(&state.manekani, radical).await?;

    Ok(HttpResponse::Ok().json(radical))
}

#[post("")]
pub async fn create(
    req: web::Json<InsertRadical>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let radical = req.into_inner();

    let created = insert(&state.manekani, radical).await?;

    debug!(
        "Created radical '{}': {}",
        created.name,
        created.id.to_string()
    );
    Ok(HttpResponse::Ok().json(created))
}

#[get("from-kanji/{kanji}")]
pub async fn from_kanji(
    kanji_symbol: web::Path<String>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let symbol = kanji_symbol.into_inner();
    let kanji = GetKanji { symbol };

    info!("Searching radicals from kanji: {}", kanji.symbol);
    let radicals = query_by_kanji(&state.manekani, kanji).await?;

    Ok(HttpResponse::Ok().json(radicals))
}
