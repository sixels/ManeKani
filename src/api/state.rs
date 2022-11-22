use std::env;

use sqlx::{postgres::PgConnectOptions, PgPool};

#[derive(Clone)]
pub struct State {
    pub db: PgPool,
}

impl State {
    pub async fn new() -> Self {
        let db_host = env::var("DB_HOST").unwrap();
        let db_port = env::var("DB_PORT")
            .unwrap_or_default()
            .parse()
            .unwrap_or(5432);
        let db_password = env::var("DB_PASSWORD").unwrap();
        let db_username = env::var("DB_USERNAME").unwrap();
        let db_name = env::var("DB_NAME").unwrap();

        let pool = PgPool::connect_with(
            PgConnectOptions::new()
                .application_name("waniklone")
                .host(&db_host)
                .port(db_port)
                .database(&db_name)
                .username(&db_username)
                .password(&db_password),
        )
        .await
        .unwrap();

        Self { db: pool }
    }
}
