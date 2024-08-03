use std::sync::Arc;

use actix_web::{
    dev::{ServiceFactory, ServiceRequest},
    web, App, HttpResponse, Scope,
};
use async_graphql::{http::GraphiQLSource, EmptyMutation, EmptySubscription, Object, Schema};
use async_graphql_actix_web::GraphQL;
use manekani_db::Database;

use crate::routing::auth::middleware::AuthMiddleware;

use super::ui::settings;

pub struct Query;

#[Object]
impl Query {
    async fn howdy(&self) -> &'static str {
        "partner"
    }
}

// pub fn setup_routes(state: AppState) -> Router {

// let cors = CorsLayer::new()
//     .allow_methods([Method::GET, Method::POST])
//     // allow requests from   any origin
//     .allow_origin(AllowOrigin::list(["http://127.0.0.1:3000"
//         .parse()
//         .unwrap()]));

// Router::new()
//     // web app routes
//     .route("/settings/api-keys", get(api_keys::index::get))
//     .route(
//         "/settings/api-keys/create",
//         get(api_keys::create::get).post(api_keys::create::post),
//     )
//     .route(
//         "/settings/api-keys/manage/:api_key_id",
//         get(api_keys::manage::get).delete(api_keys::manage::delete),
//     )
//     // intern api routes
//     .route("/intern/users/on-sign-up", post(intern::users::on_sign_up))
//     // static assets
//     // .nest_service("/assets", ServeDir::new("assets"))
//     // graphql api
//     .route("/", post_service(GraphQL::new(graphql_schema)))
//     .route("/graphiql", get(graphiql))
//     .layer(middleware::from_fn({
//         let db = state.db.clone();
//         let auth_manager = state.auth_manager.clone();
//         move |req, next| auth_guard(req, next, auth_manager.clone(), db.clone())
//     }))
// .with_state(state)
// }

pub struct SettingsRouter;

impl SettingsRouter {
    pub fn scope(self) -> Scope {
        web::scope("/settings").service(
            web::scope("/api-keys")
                .wrap(AuthMiddleware)
                .service(settings::api_keys::index::get)
                .service(settings::api_keys::create::get)
                .service(settings::api_keys::create::post)
                .service(settings::api_keys::manage::get)
                .service(settings::api_keys::manage::delete),
        )
    }
}

pub trait ManeKaniGQLApi {
    fn setup_graphql(self, db: Arc<Database>) -> Self;
}

pub async fn graphiql() -> HttpResponse {
    HttpResponse::Ok()
        .insert_header(("content-type", "text/html; charset=utf-8"))
        .body(GraphiQLSource::build().endpoint("/").finish())
}

impl<T> ManeKaniGQLApi for App<T>
where
    T: ServiceFactory<ServiceRequest, Config = (), Error = actix_web::Error, InitError = ()>,
{
    fn setup_graphql(self, db: Arc<Database>) -> Self {
        let graphql_schema = Schema::build(Query, EmptyMutation, EmptySubscription)
            .data(db)
            .finish();

        let graphql = GraphQL::new(graphql_schema.clone());

        self.app_data(graphql_schema)
            .service(web::resource("/").route(web::post().to(graphql)))
            .service(web::resource("/graphiql").route(web::get().to(graphiql)))
    }
}
