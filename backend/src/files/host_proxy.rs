use std::sync::Arc;

use actix_web::{body::SizedStream, get, web, HttpResponse};
use manekani_service_s3::{domain::query_file, model::file::RequestQuery};

use crate::{api::state::State, error::Error as ApiError};

#[get("images/{category}/{name}")]
pub async fn images(
    slug: web::Path<(String, String)>,
    state: web::Data<Arc<State>>,
) -> Result<HttpResponse, ApiError> {
    let s3 = &state.s3;

    let slug = slug.into_inner();

    let query = RequestQuery {
        category: format!("images/{}", slug.0),
        name: slug.1,
    };
    let fstream = query_file(s3, query).await?;

    Ok(HttpResponse::Ok().body(SizedStream::new(fstream.size, fstream.stream)))
}
