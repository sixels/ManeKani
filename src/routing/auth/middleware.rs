use std::rc::Rc;

use actix_web::{
    dev::{forward_ready, Service, ServiceRequest, ServiceResponse, Transform},
    web,
};
use futures_util::future::{ready, LocalBoxFuture, Ready};
use manekani_auth::AuthManager;

use super::session::{AuthSession, MANEKANI_SESSION_COOKIE};

// pub async fn auth_guard(
//     mut req: Request,
//     next: axum::middleware::Next,
//     auth_manager: Arc<AuthManager>,
//     db: Arc<Database>,
// ) -> axum::response::Response {
//     let session_token = req
//         .headers()
//         .get_all("Cookie")
//         .into_iter()
//         .filter_map(|cookies| {
//             cookies.to_str().ok().map(|cookies| {
//                 cookies.split(';').filter_map(|cookie| {
//                     cookie.parse::<cookie::Cookie>().ok().and_then(|cookie| {
//                         if cookie.name() == "manekani_session" {
//                             Some(cookie.to_string())
//                         } else {
//                             None
//                         }
//                     })
//                 })
//             })
//         })
//         .flatten()
//         .collect::<Vec<String>>();

//     let session_token = (!session_token.is_empty()).then(|| session_token.join("; "));

//     req.extensions_mut().insert(AuthLayer {
//         session: session_token.map(|token| AuthLayerInner { token, db }),
//         auth_manager: auth_manager.clone(),
//     });

//     next.run(req).await
// }

pub struct AuthMiddleware;

impl<S, B> Transform<S, ServiceRequest> for AuthMiddleware
where
    S: Service<ServiceRequest, Response = ServiceResponse<B>, Error = actix_web::Error> + 'static,
    S::Future: 'static,
    B: 'static,
{
    type Response = ServiceResponse<B>;
    type Error = S::Error;
    type InitError = ();
    type Transform = InnerAuthMiddleware<S>;
    type Future = Ready<Result<Self::Transform, Self::InitError>>;

    fn new_transform(&self, service: S) -> Self::Future {
        ready(Ok(InnerAuthMiddleware {
            service: Rc::new(service),
        }))
    }
}

pub struct InnerAuthMiddleware<S> {
    service: Rc<S>,
}

impl<S, B> Service<ServiceRequest> for InnerAuthMiddleware<S>
where
    S: Service<ServiceRequest, Response = ServiceResponse<B>, Error = actix_web::Error> + 'static,
    S::Future: 'static,
    B: 'static,
{
    type Response = ServiceResponse<B>;
    type Error = S::Error;
    type Future = LocalBoxFuture<'static, Result<Self::Response, Self::Error>>;

    forward_ready!(service);

    fn call(&self, req: ServiceRequest) -> Self::Future {
        let service = Rc::clone(&self.service);

        Box::pin(async move {
            let auth_manager = req.app_data::<web::Data<AuthManager>>().ok_or_else(|| {
                actix_web::error::ErrorInternalServerError("AuthManager not found in request")
            })?;

            let session_token = extract_session_cookie(&req);

            AuthSession::set_session(&req, auth_manager, session_token.as_deref()).await?;

            let res = service.call(req).await?;

            if session_token.is_none() {
                return Err(actix_web::error::ErrorUnauthorized("missing session token"));
            }

            Ok(res)
        })
    }
}

fn extract_session_cookie(req: &ServiceRequest) -> Option<String> {
    req.cookies().ok().and_then(|cookies| {
        cookies.iter().find_map(|cookie| {
            if cookie.name() == MANEKANI_SESSION_COOKIE {
                Some(cookie.value().to_string())
            } else {
                None
            }
        })
    })
}
