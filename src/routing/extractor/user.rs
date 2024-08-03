// use axum::{
//     extract::FromRequestParts,
//     response::{IntoResponse, Redirect, Response},
// };
// use http::{request::Parts, HeaderMap, StatusCode};
// use manekani_auth::AuthManager;

// use crate::routing::middleware::auth::{AuthLayer, AuthUser, GetUserError};

// pub struct User(pub AuthUser);

// #[axum::async_trait]
// impl<S> FromRequestParts<S> for User {
//     type Rejection = Response;

//     async fn from_request_parts(parts: &mut Parts, _: &S) -> Result<Self, Self::Rejection> {
//         let auth = parts
//             .extensions
//             .get::<AuthLayer>()
//             .ok_or_else(|| UserRejection::AuthExtensionNotInitialized.to_response(parts))?;

//         match auth.get_current_user().await {
//             Ok(user) => Ok(User(user)),
//             Err(e) => match e {
//                 GetUserError::MissingSessionCookies => {
//                     Err(UserRejection::MissingSessionCookie(auth).to_response(parts))
//                 }
//                 GetUserError::InvalidSession | GetUserError::UserNotFound => {
//                     Err(UserRejection::InvalidSession(auth).to_response(parts))
//                 }
//             },
//         }
//     }
// }

// impl<'a> UserRejection<'a> {
//     fn to_response(self, parts: &Parts) -> Response {
//         let return_url = parts.uri.to_string();

//         match self {
//             Self::AuthExtensionNotInitialized => {
//                 tracing::error!("auth extension was not found in request's context");
//                 StatusCode::INTERNAL_SERVER_ERROR.into_response()
//             }
//             UserRejection::MissingSessionCookie(auth) => {
//                 Redirect::to(&auth.browser_login_url(&return_url)).into_response()
//             }
//             UserRejection::InvalidSession(auth) => {
//                 let mut headers = HeaderMap::new();
//                 // clean manekani_session cookie
//                 headers.insert(
//                     "Set-Cookie",
//                     "manekani_session=; Max-Age=0; Expires=Thu, 01 Jan 1970 00:00:00 GMT"
//                         .parse()
//                         .unwrap(),
//                 );
//                 (headers, Redirect::to(&auth.browser_login_url(&return_url))).into_response()
//             }
//         }
//     }
// }

// pub enum UserRejection<'a> {
//     AuthExtensionNotInitialized,
//     MissingSessionCookie(&'a AuthManager),
//     InvalidSession(&'a AuthManager),
// }
