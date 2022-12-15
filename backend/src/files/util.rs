use actix_multipart::{Field, Multipart};
use futures_util::{Stream, StreamExt};
use manekani_service_s3::entity::file::{Wrapper, Written};
use tokio::io::AsyncWriteExt;
use tracing::warn;

/// Extract all files from the given payload
#[allow(clippy::unused_async)]
pub async fn extract_payload_files(
    payload: Multipart,
    category: &'static str,
) -> impl Stream<Item = Written> {
    payload.filter_map(|item| async {
        let Ok(field) = item else {
            return None
        };
        payload_file(field, category).await
    })
}

async fn payload_file(mut field: Field, category: &'static str) -> Option<Written> {
    let cd = field.content_disposition();

    let Some(field_name) = cd.get_name() else {
            warn!("Unnamed field");
            return None;
        };

    let field_name = field_name.to_owned();

    let Some(_filename) = cd.get_filename() else {
        warn!("field {:?} is not a file", field_name);
        return None;
    };

    let mut file = Wrapper::create().await.unwrap();

    // stream field contents to file
    while let Some(chunk) = field.next().await {
        let data = chunk.unwrap();
        file.write_all(&data).await.unwrap();
    }

    Some(file.finish(category, field_name).await)
}
