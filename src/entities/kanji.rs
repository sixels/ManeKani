use time::OffsetDateTime;
use typed_builder::TypedBuilder;
use uuid::Uuid;

/// Represents a Kanji.
#[derive(sqlx::FromRow, Debug, Clone)]
pub struct Kanji {
    pub id: Uuid,
    pub created_at: OffsetDateTime,
    pub updated_at: OffsetDateTime,

    /// The kanji name.
    pub name: String,
    /// Alternative names for the kanji.
    pub alt_names: Vec<String>,
    /// The kanji symbol.
    pub symbol: String,
    /// The main reading of the kanji.
    pub reading: String,
    /// Onyomi readings for the kanji.
    pub onyomi: Vec<String>,
    /// Kunyomi readings for the kanji.
    pub kunyomi: Vec<String>,
    /// Nanori readings for the kanji.
    pub nanori: Vec<String>,
    /// Mnemonics to help you remember the kanji meaning.
    /// They are stored in a string using markdown syntax.
    pub meaning_mnemonic: String,
    /// Mnemonics to help you remember the kanji reading.
    /// They are stored in a string using markdown syntax.
    pub reading_mnemonic: String,
    /// Kanji name synonyms defined by the user.
    pub user_synonyms: Option<Vec<String>>,
    /// User notes on this kanji meaning.
    pub user_meaning_note: Option<String>,
    /// User notes on this kanji reading.
    pub user_reading_note: Option<String>,
}

#[derive(Debug, Clone, TypedBuilder)]
pub struct InsertKanji {
    #[builder(setter(into))]
    pub name: String,
    #[builder(default)]
    pub alt_names: Vec<String>,
    #[builder(setter(into))]
    pub symbol: String,
    #[builder(setter(into))]
    pub reading: String,
    #[builder(default)]
    pub onyomi: Vec<String>,
    #[builder(default)]
    pub kunyomi: Vec<String>,
    #[builder(default)]
    pub nanori: Vec<String>,
    #[builder(setter(into))]
    pub meaning_mnemonic: String,
    #[builder(setter(into))]
    pub reading_mnemonic: String,
    pub radical_composition: Vec<String>,
}
