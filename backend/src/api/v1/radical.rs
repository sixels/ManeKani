use std::sync::Arc;

use actix_multipart::Multipart;
use actix_web::{get, post, web, HttpResponse};
use futures_util::StreamExt;
use manekani_service_pg::{
    domain::radical::RadicalRepository,
    entity::{radical::UpdateRadical, GetKanji, GetRadical, InsertRadical},
};
use tracing::{debug, info};

use crate::{
    api::state::State,
    error::Error as ApiError,
    files::{
        upload::{upload_file, UploadStatus},
        utils::extract_payload_files,
    },
};

#[get("{radical}")]
pub async fn get(
    radical_name: web::Path<String>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let name = radical_name.into_inner();
    let radical = GetRadical { name };

    info!("Querying radical: {}", radical.name);
    let radical = state.manekani.query_radical(radical).await?;

    Ok(HttpResponse::Ok().json(radical))
}

#[post("")]
pub async fn create(
    req: web::Json<InsertRadical>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let radical = req.into_inner();

    let created = state.manekani.insert_radical(radical).await?;

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
    let radicals = state.manekani.query_radical_by_kanji(kanji).await?;

    Ok(HttpResponse::Ok().json(radicals))
}

#[post("symbol")]
pub async fn upload_radical_symbol(
    payload: Multipart,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    info!("Uploading radicals symbol");

    let s3 = &state.s3;
    let manekani = &state.manekani;

    let uploads = extract_payload_files(payload, "images/radical")
        .await
        .map(|file| async {
            let status = upload_file(s3, file).await;

            if let UploadStatus::Created(file) = &status {
                let key = &file.key;
                let name = &file.name;

                let update_radical = UpdateRadical {
                    name: name.clone(),
                    symbol: Some(key.clone()),
                    ..UpdateRadical::default()
                };

                manekani.update_radical(update_radical).await?;
            };

            Result::<UploadStatus, ApiError>::Ok(status)
        })
        .buffer_unordered(5);

    let status = uploads.collect::<Vec<_>>().await;

    Ok(HttpResponse::MultiStatus().json(status))
}
