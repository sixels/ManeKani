use sqlx::{pool::PoolConnection, PgPool, Postgres};

pub mod kanji;
pub mod radical;
pub mod vocabulary;

pub struct Repository {
    database: PgPool,
}

impl Repository {
    pub fn new(database: PgPool) -> Self {
        Self { database }
    }

    pub async fn connection(&self) -> PoolConnection<Postgres> {
        self.database
            .acquire()
            .await
            .expect("Could not get a database connection")
    }

    pub async fn connect(db_url: String) -> Self {
        let pool = PgPool::connect(&db_url).await.unwrap();
        Self { database: pool }
    }
}
