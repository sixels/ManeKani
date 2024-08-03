use std::sync::Arc;

use manekani::{app::ManeKani, state::AppState};
use manekani_auth::{AuthManager, AuthOptions};

#[actix_web::main]
async fn main() {
    tracing_subscriber::fmt().init();

    // TODO: get configuration from env

    let db = manekani_db::Database::new(
        "postgres://manekani:secret@postgres-manekani/manekani-test?sslmode=disable",
    )
    .await
    .unwrap();

    let auth_manager = AuthManager::new(AuthOptions {
        base_url: String::from("http://kratos:4433"),
        proxy_url: Some(String::from("http://127.0.0.1:4433")),
    })
    .await
    .unwrap();

    let state = AppState {
        db: Arc::new(db),
        auth_manager: Arc::new(auth_manager),
    };

    ManeKani::new(state).listen("0.0.0.0:9999").await.unwrap();
}
