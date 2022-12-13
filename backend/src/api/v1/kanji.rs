use std::sync::Arc;

use actix_web::{get, post, web, HttpResponse};
use manekani_pg::{
    domain::kanji::KanjiRepository,
    entity::{GetKanji, GetRadical, InsertKanji},
};
use tracing::{debug, info};

use crate::api::{error::Error as ApiError, state::State};

#[get("{kanji}")]
pub async fn get(
    kanji_symbol: web::Path<String>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let symbol = kanji_symbol.into_inner();
    let kanji = GetKanji { symbol };

    info!("Querying kanji: {}", kanji.symbol);
    let kanji = state.manekani.query_kanji(kanji).await?;

    Ok(HttpResponse::Ok().json(kanji))
}

#[post("")]
pub async fn create(
    req: web::Json<InsertKanji>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let kanji = req.into_inner();

    let created = state.manekani.insert_kanji(kanji).await?;

    debug!(
        "Created kanji '{}': {}",
        created.symbol,
        created.id.to_string()
    );
    Ok(HttpResponse::Ok().json(created))
}

#[get("from-radical/{radical}")]
pub async fn from_radical(
    radical_name: web::Path<String>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let name = radical_name.into_inner();
    let radical = GetRadical { name };

    info!("Searching kanjis from radical: {}", radical.name);
    let kanjis = state.manekani.query_kanji_by_radical(radical).await?;

    Ok(HttpResponse::Ok().json(kanjis))
}
