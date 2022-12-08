pub mod host_proxy;
pub mod upload;

pub mod utils {
    use actix_multipart::{Field, Multipart};
    use futures_util::{Stream, StreamExt};
    use manekani_s3::entity::file::{FileWrapper, WrittenFile};
    use tokio::io::AsyncWriteExt;
    use tracing::warn;

    pub async fn extract_payload_files(
        payload: Multipart,
        category: &'static str,
    ) -> impl Stream<Item = WrittenFile> {
        payload.filter_map(|item| {
            let field = item.expect("split_payload err");
            payload_file(field, category)
        })
    }

    async fn payload_file(mut field: Field, category: &'static str) -> Option<WrittenFile> {
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

        let mut file = FileWrapper::create().await.unwrap();

        // stream field contents to file
        while let Some(chunk) = field.next().await {
            let data = chunk.unwrap();
            file.write_all(&data).await.unwrap();
        }

        Some(file.finish(category, field_name).await)
    }
}
