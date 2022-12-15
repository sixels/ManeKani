use std::sync::Arc;

use actix_web::{get, post, web, HttpResponse};
use manekani_service_pg::{
    domain::kanji::Repository,
    model::{ReqKanjiInsert, ReqKanjiQuery, ReqRadicalQuery},
};
use tracing::{debug, info};

use crate::{api::state::State, error::Error as ApiError};

#[get("{kanji}")]
pub async fn get(
    kanji_symbol: web::Path<String>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let symbol = kanji_symbol.into_inner();
    let kanji = ReqKanjiQuery { symbol };

    info!("Querying kanji: {}", kanji.symbol);
    let kanji = state.manekani.query_kanji(kanji).await?;

    Ok(HttpResponse::Ok().json(kanji))
}

#[post("")]
pub async fn create(
    req: web::Json<ReqKanjiInsert>,
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
    let radical = ReqRadicalQuery { name };

    info!("Searching kanji from radical: {}", radical.name);
    let kanji = state.manekani.query_kanji_by_radical(radical).await?;

    Ok(HttpResponse::Ok().json(kanji))
}
