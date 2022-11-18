use time::OffsetDateTime;
use typed_builder::TypedBuilder;
use uuid::Uuid;

/// Represents a Kanji radical.
#[derive(sqlx::FromRow, Debug, Clone)]
pub struct Radical {
    pub id: Uuid,
    pub created_at: OffsetDateTime,
    pub updated_at: OffsetDateTime,

    /// The radical name.
    pub name: String,
    /// The radical symbol.
    pub symbol: String,
    /// Mnemonics to help you remember the radical meaning.
    /// They are stored in a string using markdown syntax.
    pub meaning_mnemonic: String,
    /// Radical name synonyms defined by the user.
    pub user_synonyms: Option<Vec<String>>,
    /// User notes on this radical.
    pub user_meaning_note: Option<String>,
}

/// A subset of `Radical` used for database insertion.
#[derive(Debug, Clone, Default, TypedBuilder)]
pub struct InsertRadical {
    #[builder(setter(into))]
    pub name: String,
    #[builder(setter(into))]
    pub symbol: String,
    #[builder(setter(into))]
    pub meaning_mnemonic: String,
}

impl InsertRadical {
    pub fn new(name: String, symbol: String, meaning_mnemonic: String) -> Self {
        Self {
            name,
            symbol,
            meaning_mnemonic,
        }
    }
}
