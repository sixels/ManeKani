use serde::{Deserialize, Serialize};
use time::OffsetDateTime;
use typed_builder::TypedBuilder;
use uuid::Uuid;

/// Represents a Kanji.
#[derive(sqlx::FromRow, Debug, Clone, Serialize)]
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

#[derive(Debug, Clone, TypedBuilder, Deserialize)]
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

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct GetKanji {
    pub symbol: String,
}

// #[cfg(test)]
// pub(crate) fn kanji_genius() -> InsertKanji {
//     InsertKanji::builder()
//         .name("genius")
//         .symbol("才")
//         .reading("さい")
//         .onyomi(vec!["さい".to_owned()])
//         .meaning_mnemonic(r#"On the ground you put barbs at the bottom of a slide because you're a genius trying to catch another genius."#)
//         .reading_mnemonic(r#"You check your genius trap, and there's a genius stuck in the barbs! You know it's a genius because they're actually a cyborg (さい)."#)
//         .radical_composition(vec!["一".into(),"亅".into(), "丿".into()])
//         .build()
// }

#[cfg(test)]
pub(crate) fn kanji_middle() -> InsertKanji {
    InsertKanji::builder()
        .name("middle")
        .symbol("中")
        .reading("ちゅう")
        .onyomi(vec!["ちゅう".to_owned()])
        .kunyomi(vec!["なか".to_owned()])
        .meaning_mnemonic(r#"The radical Middle and the kanji Middle are both the same. So if you know one, you know the other."#)
        .reading_mnemonic(r#"To remember the reading for this kanji, we use the word Chewbacca to pull up ちゅう in our memory. If you remember back to the radical 中, the middle of your mouth was stabbed with a stick. You look up to see who did it. There stands Chewbacca, doing his Chewbacca yell. And it isn't a stick in your mouth, it's an arrow from his crossbow (a bowcaster, actually). It just so happens Chewbacca's bowcaster looks just like this kanji too. Go figure."#)
        .radical_composition(vec!["中".into()])
        .build()
}

#[cfg(test)]
pub(crate) fn kanji_stop() -> InsertKanji {
    InsertKanji::builder()
        .name("stop")
        .symbol("止")
        .reading("し")
        .onyomi(vec!["し".to_owned()])
        .kunyomi(vec!["と".to_owned(), "や".to_owned()])
        .meaning_mnemonic(r#"The stop radical is the same as the stop kanji."#)
        .reading_mnemonic(r#"You have to stop because there is a sheep (し) in front of you. You try to walk around the sheep, but it moves to stop in front of you again. Every time the sheep stops you stop too."#)
        .radical_composition(vec!["止".into()])
        .build()
}
