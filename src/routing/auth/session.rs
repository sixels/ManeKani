use actix_web::{dev::ServiceRequest, HttpMessage};
use manekani_auth::{AuthManager, SsoSession};
use uuid::Uuid;

pub(super) const MANEKANI_SESSION_COOKIE: &str = "manekani_session";

#[derive(Debug, Clone)]
pub struct AuthSession(pub(super) SsoSession);

impl AuthSession {
    pub async fn set_session(
        req: &ServiceRequest,
        auth_manager: &AuthManager,
        token: Option<&str>,
    ) -> Result<(), actix_web::Error> {
        match token {
            Some(token) => {
                let cookie = format!("{}={}", MANEKANI_SESSION_COOKIE, token);
                let session = match auth_manager.get_sso_session(&cookie).await {
                    Some(session) => session,
                    None => {
                        return Err(actix_web::error::ErrorUnauthorized("Invalid session token"));
                    }
                };

                req.extensions_mut().insert(AuthSession(session));
                Ok(())
            }
            None => Err(actix_web::error::ErrorUnauthorized("Missing session token")),
        }
    }

    pub fn user_id(&self) -> Uuid {
        self.0
            .metadata_public
            .as_ref()
            .map(|metadata| metadata.id)
            .unwrap_or_default()
    }
}
