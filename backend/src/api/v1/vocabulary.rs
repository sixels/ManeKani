use std::sync::Arc;

use actix_web::{get, post, web, HttpResponse};
use manekani_pg::{
    domain::vocabulary::{insert, query, query_by_kanji},
    entity::{GetKanji, GetVocabulary, InsertVocabulary},
};
use tracing::{debug, info};

use crate::api::{error::Error as ApiError, state::State};

#[get("{vocabulary}")]
pub async fn get(
    vocabulary_word: web::Path<String>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let word = vocabulary_word.into_inner();
    let vocab = GetVocabulary { word };

    info!("Getting vocabulary '{}'", vocab.word);
    let vocabulary = query(&state.manekani, vocab).await?;

    Ok(HttpResponse::Ok().json(vocabulary))
}

#[post("")]
pub async fn create(
    req: web::Json<InsertVocabulary>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let vocabulary = req.into_inner();

    let created = insert(&state.manekani, vocabulary).await?;

    debug!(
        "Created vocabulary '{}': '{}'",
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

    info!("Searching vocabularies from kanji: {}", kanji.symbol);
    let vocabularies = query_by_kanji(&state.manekani, kanji).await?;

    Ok(HttpResponse::Ok().json(vocabularies))
}
