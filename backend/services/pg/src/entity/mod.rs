pub(crate) mod kanji;
pub(crate) mod radical;
pub(crate) mod vocabulary;

pub use kanji::{Kanji, ReqKanjiInsert, ReqKanjiQuery};
pub use radical::{Radical, ReqRadicalInsert, ReqRadicalQuery, ReqRadicalUpdate};
pub use vocabulary::{ReqVocabularyInsert, ReqVocabularyQuery, Vocabulary};

pub type KanjiPartial = kanji::Partial;
pub type RadicalPartial = radical::Partial;
pub type VocabularyPartial = vocabulary::Partial;
