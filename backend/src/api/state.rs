use manekani_pg::Repository as ManeKaniRepo;
use std::env;

pub struct State {
    pub manekani: ManeKaniRepo,
}

impl State {
    pub async fn new() -> Self {
        let manekani_db = env::var("MANEKANI_DB_URL").expect("MANEKANI_DB_URL must be set");
        let manekani = ManeKaniRepo::connect(manekani_db).await;

        Self { manekani }
    }
}
