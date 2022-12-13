use std::sync::Arc;

use actix_web::{body::SizedStream, get, web, HttpResponse};
use manekani_service_s3::{domain::query_file, entity::file::QueryFile};

use crate::{error::Error as ApiError, api::state::State};

#[get("images/{category}/{name}")]
pub async fn images(
    slug: web::Path<(String, String)>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let s3 = &state.s3;

    let slug = slug.into_inner();

    let query = QueryFile {
        category: format!("images/{}", slug.0),
        name: slug.1,
    };
    let (size, stream) = query_file(s3, query).await?;

    Ok(HttpResponse::Ok().body(SizedStream::new(size, stream)))
}
