use manekani_service_pg::Repository as ManeKaniRepo;
use manekani_service_s3::repository::S3Repo;
use std::env;

pub struct State {
    pub manekani: ManeKaniRepo,
    pub s3: S3Repo,
}

impl State {
    pub async fn new() -> Self {
        let manekani_db = env::var("MANEKANI_DB_URL").expect("MANEKANI_DB_URL must be set");
        let manekani = ManeKaniRepo::connect(manekani_db).await;

        let s3 = S3Repo::new("manekani".to_owned()).await;

        Self { manekani, s3 }
    }
}
