use manekani::server;

#[warn(clippy::pedantic)]
#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenvy::dotenv().unwrap();

    let subscriber = tracing_subscriber::FmtSubscriber::builder()
        .log_internal_errors(true)
        .finish();
    tracing::subscriber::set_global_default(subscriber).unwrap();

    server::serve("127.0.0.1:8081").await
}
