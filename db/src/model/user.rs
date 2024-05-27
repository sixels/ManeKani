use time::OffsetDateTime;
use uuid::Uuid;

pub struct UserModel {
    pub id: String,
    pub username: Option<String>,
    pub is_verified: bool,
    pub is_complete: bool,
    pub created_at: OffsetDateTime,
    pub updated_at: OffsetDateTime,
    pub deck_ids: Option<Vec<Uuid>>,
}
