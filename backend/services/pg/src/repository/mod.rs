use sqlx::{pool::PoolConnection, PgPool, Postgres};

pub mod kanji;
pub mod radical;
pub mod vocabulary;

pub struct Repository {
    database: PgPool,
}

impl Repository {
    #[must_use]
    pub fn new(database: PgPool) -> Self {
        Self { database }
    }

    /// Returns the connection of this [`Repository`].
    ///
    /// # Panics
    ///
    /// Panics if it is not connected with the database
    pub async fn connection(&self) -> PoolConnection<Postgres> {
        self.database
            .acquire()
            .await
            .expect("Could not get a database connection")
    }

    /// Starts the connection with the postgres database
    ///
    /// # Panics
    ///
    /// Panics if it could not connect with the database
    pub async fn connect(db_url: String) -> Self {
        let pool = PgPool::connect(&db_url).await.unwrap();
        Self::new(pool)
    }
}
