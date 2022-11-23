use std::env;

use sqlx::PgPool;

#[derive(Clone)]
pub struct State {
    pub db: PgPool,
}

impl State {
    pub async fn new() -> Self {
        let db_url = env::var("DATABASE_URL").expect("DATABASE_URL environment is not set");

        let pool = PgPool::connect(&db_url).await.unwrap();

        Self { db: pool }
    }
}
