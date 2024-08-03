use std::sync::Arc;

use manekani_auth::AuthManager;
use manekani_db::Database;

#[derive(Debug, Clone)]
pub struct AppState {
    pub db: Arc<Database>,
    pub auth_manager: Arc<AuthManager>,
}
