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

// #[cfg(test)]
// pub(crate) fn radical_ground() -> InsertRadical {
//     InsertRadical::builder()
//         .name("ground")
//         .symbol("一")
//         .meaning_mnemonic("This radical consists of a single, horizontal stroke. What's the biggest, single, horizontal stroke? That's the ground. Look at the ground, look at this radical, now look at the ground again. Kind of the same, right?")
//         .build()
// }

#[cfg(test)]
pub(crate) fn radical_barb() -> InsertRadical {
    InsertRadical::builder()
        .name("barb")
        .symbol("亅")
        .meaning_mnemonic(r#"This radical is shaped like a barb. Like the kind you'd see on barb wire. Imagine one of these getting stuck to your arm or your clothes. Think about how much it would hurt with that little hook on the end sticking into you. Say out loud, "Oh dang, I got a barb stuck in me!""#)
        .build()
}

// #[cfg(test)]
// pub(crate) fn radical_slide() -> InsertRadical {
//     InsertRadical::builder()
//         .name("slide")
//         .symbol("丿")
//         .meaning_mnemonic("Close your eyes and imagine you're a kid again. Now open them and... look! It's a slide! Imagine little you sliding down this slide over and over, having a great time. Imagination + emotion makes for good memorization!")
//         .build()
// }

#[cfg(test)]
pub(crate) fn radical_middle() -> InsertRadical {
    InsertRadical::builder()
        .name("middle")
        .symbol("中")
        .meaning_mnemonic("There's a stick going right through the middle of someone's mouth! Imagine that mouth being yours. Measure the placement of the stick. Perfectly aligned, right in the middle. That's amazing, though to be honest, you probably have bigger problems than measuring the location of this stick in your mouth.")
        .build()
}

#[cfg(test)]
pub(crate) fn radical_stop() -> InsertRadical {
    InsertRadical::builder()
        .name("stop")
        .symbol("止")
        .meaning_mnemonic("There's a giant toe and a stick in the ground in front of you. You're driving your car toward them, but you don't see them until your lights hit them. What would these things cause you to do? Most likely stop your car right away.")
        .build()
}
