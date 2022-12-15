use std::{fmt::Debug, sync::Arc};

use actix_web::{
    middleware::{NormalizePath, TrailingSlash},
    web, App, HttpServer,
};
use tracing::info;
use tracing_actix_web::TracingLogger;

use crate::{
    api::{self, state::State},
    files,
};

/// Start the server listening at the given address
///
/// # Errors
///
/// This function will return an error if the address is already being used or
/// an IO error occurred
pub async fn serve<A>(addr: A) -> std::io::Result<()>
where
    A: std::net::ToSocketAddrs + Debug,
{
    // let db = PgPool::connect(url)

    let state = Arc::new(State::new().await);
    HttpServer::new(move || {
        App::new()
            .service(api::ping::ping)
            .service(web::scope("/files").service(files::host_proxy::images))
            .service(
                web::scope("/api/v1")
                    .service(
                        web::scope("/kanji")
                            .service(api::v1::kanji::get)
                            .service(api::v1::kanji::create)
                            .service(api::v1::kanji::from_radical),
                    )
                    .service(
                        web::scope("/radical")
                            .service(api::v1::radical::get)
                            .service(api::v1::radical::create)
                            .service(api::v1::radical::from_kanji)
                            .service(api::v1::radical::upload_radical_symbol),
                    )
                    .service(
                        web::scope("/vocabulary")
                            .service(api::v1::vocabulary::get)
                            .service(api::v1::vocabulary::create)
                            .service(api::v1::vocabulary::from_kanji),
                    ),
            )
            .wrap(TracingLogger::default())
            .wrap(NormalizePath::new(TrailingSlash::Trim))
            .app_data(web::Data::new(state.clone()))
    })
    .bind(&addr)
    .map(|app| {
        info!("Listening on {addr:?}");
        app.run()
    })?
    .await
}
