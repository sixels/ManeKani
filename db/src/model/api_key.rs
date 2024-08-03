use time::OffsetDateTime;
use uuid::Uuid;

pub struct ApiKeyModel {
    pub id: Uuid,
    pub name: String,
    pub prefix: String,
    pub claims: serde_json::Value,
    pub revoked_at: Option<OffsetDateTime>,
    pub used_at: Option<OffsetDateTime>,
    pub created_at: OffsetDateTime,
    pub updated_at: OffsetDateTime,
    pub created_by_user_id: Uuid,
}
