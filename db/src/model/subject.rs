// TODO: move comments from here

use time::OffsetDateTime;
use uuid::Uuid;

#[derive(Debug, sqlx::FromRow)]
pub struct SubjectModel {
    pub id: Uuid,
    pub category: String,
    pub level: i32,
    pub name: String,
    pub slug: String,

    /// The data of this subject.
    ///
    /// Example:
    /// ```json
    /// {
    ///   "value": "おやつ",
    ///   "image": null,
    ///   "pronunciation_audio": {
    ///     "path": "/files/ByzjB4BoUr",
    ///     "metadata": {
    ///      "type": "audio/ogg",
    ///      "voiceActor": "John Doe"
    ///     }
    ///   },
    ///   "sentences": [
    ///     {"jp": "今日はおやつにマフィンを食べた。", "en": "Today I ate a muffin as a snack."}
    ///   ],
    ///   "patterns": [
    ///     {
    ///       "pattern": "おやつに~",
    ///       "examples": [ {"jp": "おやつに食べる", "en": "to eat something as a snack"} ]
    ///     }
    ///   ]
    /// }
    /// ```
    pub data: serde_json::Value,
    /// The data used to help studying this subject.
    ///
    /// Example:
    /// ```json
    /// [
    ///   {
    ///     "name": "meaning",
    ///     "mnemonic": "...",
    ///     "answers": [
    ///       {
    ///         "value": "Correct",
    ///         "is_primary": true,
    ///         "is_valid": true,
    ///         "is_hidden": false,
    ///       },
    ///       ...
    ///     ]
    ///   },
    ///   {
    ///     "name": "reading",
    ///     "mnemonic": "...",
    ///     "answers": [ ... ]
    ///   }
    /// ]
    /// ```
    pub study_data: serde_json::Value,

    pub priority: i32,
    pub created_at: OffsetDateTime,
    pub updated_at: OffsetDateTime,
    pub deck_id: Uuid,
    pub similars: Option<Vec<Uuid>>,
    pub depends_on: Option<Vec<Uuid>>,
    pub depended_by: Option<Vec<Uuid>>,
}
