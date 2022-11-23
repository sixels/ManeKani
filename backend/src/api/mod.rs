mod ping;
mod state;
mod v1;

use std::sync::Arc;

use actix_web::{
    middleware::{NormalizePath, TrailingSlash},
    web, App, HttpServer,
};
use tracing_actix_web::TracingLogger;

use self::state::State;

pub async fn serve<A>(addr: A) -> std::io::Result<()>
where
    A: std::net::ToSocketAddrs,
{
    // let db = PgPool::connect(url)

    let state = Arc::new(State::new().await);
    HttpServer::new(move || {
        App::new()
            .wrap(TracingLogger::default())
            .wrap(NormalizePath::new(TrailingSlash::Trim))
            .app_data(web::Data::new(state.clone()))
            .service(ping::ping)
            .service(
                web::scope("/api/v1")
                    .service(
                        web::scope("/kanji")
                            .service(v1::kanji::get)
                            .service(v1::kanji::create),
                    )
                    .service(
                        web::scope("radical")
                            .service(v1::radical::get)
                            .service(v1::radical::create),
                    )
                    .service(
                        web::scope("/vocabulary")
                            .service(v1::vocabulary::get)
                            .service(v1::vocabulary::create),
                    ),
            )
    })
    .bind(addr)?
    .run()
    .await
}
