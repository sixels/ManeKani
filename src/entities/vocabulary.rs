use serde::{Deserialize, Serialize};
use time::OffsetDateTime;
use typed_builder::TypedBuilder;
use uuid::Uuid;

/// Represents a vocabulary word.
#[derive(sqlx::FromRow, Debug, Clone, Serialize)]
pub struct Vocabulary {
    pub id: Uuid,
    pub created_at: OffsetDateTime,
    pub updated_at: OffsetDateTime,

    /// The vocabulary name/meaning.
    pub name: String,
    /// The vocabulary level.
    pub level: i32,
    /// Alternative names for the vocabulary.
    pub alt_names: Vec<String>,
    /// The vocabulary writing.
    pub word: String,
    /// The vocabulary word type (e.g: noun, verb)
    pub word_type: Vec<String>,
    /// The word reading.
    pub reading: String,
    /// Mnemonics to help you remember the vocabulary meaning.
    /// They are stored in a string using markdown syntax.
    pub meaning_mnemonic: String,
    /// Mnemonics to help you remember the vocabulary reading.
    /// They are stored in a string using markdown syntax.
    pub reading_mnemonic: String,
    /// Vocabulary name synonyms defined by the user.
    pub user_synonyms: Option<Vec<String>>,
    /// User notes on this kanji meaning.
    pub user_meaning_note: Option<String>,
    /// User notes on this kanji reading.
    pub user_reading_note: Option<String>,
}

#[derive(Debug, Clone, TypedBuilder, Deserialize)]
pub struct InsertVocabulary {
    #[builder(setter(into))]
    pub name: String,
    pub level: i32,
    #[builder(default)]
    pub alt_names: Vec<String>,
    #[builder(setter(into))]
    pub word: String,
    pub word_type: Vec<String>,
    #[builder(setter(into))]
    pub reading: String,
    #[builder(setter(into))]
    pub meaning_mnemonic: String,
    #[builder(setter(into))]
    pub reading_mnemonic: String,
    pub kanji_composition: Vec<String>,
}

#[derive(Debug, Clone, TypedBuilder, Deserialize)]
pub struct GetVocabulary {
    pub word: String,
}
