use std::sync::Arc;

use manekani::{app::App, state::AppState};

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt().init();

    let db = manekani_db::Database::new(
        "postgres://manekani:secret@localhost:5433/manekani-test?sslmode=disable",
    )
    .await
    .unwrap();

    let state = AppState { db: Arc::new(db) };

    App::new(state).listen("0.0.0.0:9999").await.unwrap();
}
