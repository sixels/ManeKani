use std::sync::Arc;

use axum::extract::FromRef;
use manekani_db::Database;

#[derive(Debug, Clone)]
pub struct AppState {
    pub db: Arc<Database>,
}

impl FromRef<AppState> for Arc<Database> {
    fn from_ref(state: &AppState) -> Self {
        state.db.clone()
    }
}
