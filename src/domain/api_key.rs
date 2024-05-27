use serde::{Deserialize, Serialize};
use time::OffsetDateTime;

pub struct ApiKey {
    pub id: String,
    pub name: String,
    pub prefix: String,
    pub claims: ApiKeyClaims,
    pub used_at: Option<OffsetDateTime>,
    pub revoked_at: Option<OffsetDateTime>,
    pub updated_at: OffsetDateTime,
    pub created_at: OffsetDateTime,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct ApiKeyClaims {
    #[serde(default)]
    pub deck_write: bool,
    #[serde(default)]
    pub deck_delete: bool,
    #[serde(default)]
    pub subject_write: bool,
    #[serde(default)]
    pub subject_delete: bool,
    #[serde(default)]
    pub review_create: bool,
    #[serde(default)]
    pub study_data_write: bool,
    #[serde(default)]
    pub study_data_delete: bool,
}
