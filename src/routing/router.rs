use async_graphql::{http::GraphiQLSource, EmptyMutation, EmptySubscription, Object, Schema};
use async_graphql_axum::GraphQL;
use axum::{
    response::Html,
    routing::{get, post_service},
    Router,
};
use tower_http::services::ServeDir;

use crate::{state::AppState, ui::settings::api_keys};

pub struct Query;

#[Object]
impl Query {
    async fn howdy(&self) -> &'static str {
        "partner"
    }
}

async fn graphiql() -> Html<String> {
    Html(
        GraphiQLSource::build()
            .endpoint("/")
            .subscription_endpoint("/ws")
            .finish(),
    )
}

pub fn setup_router(state: AppState) -> Router {
    let graphql_schema = Schema::build(Query, EmptyMutation, EmptySubscription)
        .data(state.db.clone())
        .finish();

    Router::new()
        .route("/", post_service(GraphQL::new(graphql_schema)))
        .route("/graphiql", get(graphiql))
        // static assets
        .nest_service("/assets", ServeDir::new("assets"))
        // web app routes
        .route("/settings/api-keys", get(api_keys::index::get))
        .route(
            "/settings/api-keys/create",
            get(api_keys::create::get).post(api_keys::create::post),
        )
        .route(
            "/settings/api-keys/manage/:api_key_id",
            get(api_keys::manage::get).delete(api_keys::manage::delete),
        )
        .with_state(state)
}
