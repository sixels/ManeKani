#[warn(clippy::pedantic)]
#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let subscriber = tracing_subscriber::FmtSubscriber::builder()
        //
        .finish();
    tracing::subscriber::set_global_default(subscriber).unwrap();

    dotenvy::dotenv().unwrap();

    manekani::api::serve("127.0.0.1:8081").await
}
