pub mod kanji;
pub mod radical;
pub mod vocabulary;

pub use manekani_service_common::repository::error::Error;

pub use self::{
    kanji::Repository as KanjiRepository, radical::Repository as RadicalRepository,
    vocabulary::Repository as VocabularyRepository,
};
