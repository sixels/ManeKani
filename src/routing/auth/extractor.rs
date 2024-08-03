use actix_web::{dev::Payload, web, FromRequest, HttpMessage, HttpRequest};
use futures_util::future::LocalBoxFuture;
use manekani_db::{query::user, Database};
use time::OffsetDateTime;
use uuid::Uuid;

use super::session::AuthSession;

pub struct CurrentUser {
    pub id: Uuid,
    pub email: String,
    pub username: Option<String>,
    pub display_name: String,
    pub is_verified: bool,
    pub is_complete: bool,

    pub created_at: OffsetDateTime,
    pub updated_at: OffsetDateTime,
}

impl FromRequest for CurrentUser {
    type Error = actix_web::Error;
    type Future = LocalBoxFuture<'static, Result<Self, Self::Error>>;

    #[inline]
    fn from_request(req: &HttpRequest, _: &mut Payload) -> Self::Future {
        let db = req
            .app_data::<web::Data<Database>>()
            .ok_or_else(|| {
                actix_web::error::ErrorInternalServerError("Database not found in request")
            })
            .cloned();

        let session = req
            .extensions()
            .get::<AuthSession>()
            .ok_or_else(|| actix_web::error::ErrorUnauthorized("Session not found in request"))
            .cloned();

        Box::pin(async move {
            let db = db?;
            let session = session?;

            let user = user::read::get_user(&db, session.user_id())
                .await
                .map_err(|e| actix_web::error::ErrorInternalServerError(e))?;

            Ok(Self {
                id: user.id,
                email: user.email,
                username: user.username,
                // display_name: user.display_name,
                display_name: Default::default(),
                is_verified: user.is_verified,
                is_complete: user.is_complete,
                created_at: user.created_at,
                updated_at: user.updated_at,
            })
        })
    }
}
