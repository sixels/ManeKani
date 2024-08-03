use time::OffsetDateTime;
use uuid::Uuid;

#[derive(Debug, sqlx::FromRow)]
pub struct DeckModel {
    pub id: Uuid,
    pub name: String,
    pub slug: String,
    pub image: Option<String>,
    pub description: Option<String>,
    pub tags: Option<Vec<String>>,
    pub is_featured: bool,
    pub is_public: bool,
    pub created_at: OffsetDateTime,
    pub updated_at: OffsetDateTime,
    pub created_by_user_id: Uuid,
    pub allowed_users: Option<Vec<Uuid>>,
}
