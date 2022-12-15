use manekani_service_common::repository::{RepoInsertable, RepoQueryable};

use crate::model::{
    ReqKanjiQuery, ReqVocabularyInsert, ReqVocabularyQuery, Vocabulary, VocabularyPartial,
};

use super::Error;

#[async_trait::async_trait]
pub trait Repository:
    RepoQueryable<ReqVocabularyQuery, Vocabulary>
    + RepoQueryable<ReqKanjiQuery, Vec<VocabularyPartial>>
    + RepoInsertable<ReqVocabularyInsert, Vocabulary>
{
    async fn query_vocabulary(&self, vocab: ReqVocabularyQuery) -> Result<Vocabulary, Error> {
        Ok(self.query(vocab).await?)
    }

    async fn query_vocabulary_by_kanji(
        &self,
        kanji: ReqKanjiQuery,
    ) -> Result<Vec<VocabularyPartial>, Error> {
        Ok(self.query(kanji).await?)
    }

    async fn insert_vocabulary(
        &self,
        vocabulary: ReqVocabularyInsert,
    ) -> Result<Vocabulary, Error> {
        Ok(self.insert(vocabulary).await?)
    }
}

impl<T> Repository for T where
    T: RepoQueryable<ReqVocabularyQuery, Vocabulary>
        + RepoQueryable<ReqKanjiQuery, Vec<VocabularyPartial>>
        + RepoInsertable<ReqVocabularyInsert, Vocabulary>
{
}

#[cfg(test)]
mod tests {

    use crate::domain::{KanjiRepository, RadicalRepository, VocabularyRepository};
    use crate::model::{kanji, radical, ReqVocabularyInsert};
    use crate::Repository;

    use manekani_service_common::repository::{error::Error, InsertError};
    use sqlx::PgPool;

    #[sqlx::test]
    async fn it_should_create_a_new_vocabulary(pool: PgPool) -> Result<(), Error> {
        let repo = Repository::new(pool);

        let kanji_middle = {
            repo.insert_radical(radical::middle()).await?;
            repo.insert_kanji(kanji::middle()).await?
        };
        let kanji_stop = {
            repo.insert_radical(radical::stop()).await?;
            repo.insert_kanji(kanji::stop()).await?
        };

        let vocabulary = ReqVocabularyInsert::builder()
            .name("Suspension")
            .level(3)
            .alt_names(vec!["Cancellation".to_owned(), "Discontinuation".to_owned()])
            .word("中止")
            .word_type(vec!["noun".to_owned(), "する verb".to_owned()])
            .reading("ちゅうし")
            .meaning_mnemonic(
                r#"When something is in the middle of an action but it's stopped, this means it's put into suspension. Imagine getting frozen while you're running. It's kind of like that."#
            )
            .reading_mnemonic(
                r#"This is a jukugo word, which usually means on'yomi readings from the kanji. If you know the readings of your kanji you'll know how to read this as well."#
            )
            .kanji_composition(vec![kanji_middle.symbol, kanji_stop.symbol])
            .build();

        let created_vocabulary = repo.insert_vocabulary(vocabulary.clone()).await?;

        assert_eq!(created_vocabulary.name, vocabulary.name);
        assert_eq!(created_vocabulary.word, vocabulary.word);
        assert_eq!(created_vocabulary.word_type, vocabulary.word_type);
        assert_eq!(created_vocabulary.reading, vocabulary.reading);
        assert_eq!(
            created_vocabulary.meaning_mnemonic,
            vocabulary.meaning_mnemonic
        );
        assert_eq!(
            created_vocabulary.reading_mnemonic,
            vocabulary.reading_mnemonic
        );
        assert_eq!(created_vocabulary.user_synonyms, None);
        assert_eq!(created_vocabulary.user_meaning_note, None);

        Ok(())
    }

    #[sqlx::test]
    async fn it_should_collide_with_an_existing_vocabulary(pool: PgPool) -> Result<(), Error> {
        let repo = Repository::new(pool);

        let kanji_middle = {
            repo.insert_radical(radical::middle()).await?;
            repo.insert_kanji(kanji::middle()).await?
        };
        let kanji_stop = {
            repo.insert_radical(radical::stop()).await?;
            repo.insert_kanji(kanji::stop()).await?
        };
        let vocabulary = ReqVocabularyInsert::builder()
            .name("Suspension")
            .level(3)
            .alt_names(vec!["Cancellation".to_owned(), "Discontinuation".to_owned()])
            .word("中止")
            .word_type(vec!["noun".to_owned(), "する verb".to_owned()])
            .reading("ちゅうし")
            .meaning_mnemonic(
                r#"When something is in the middle of an action but it's stopped, this means it's put into suspension. Imagine getting frozen while you're running. It's kind of like that."#
            )
            .reading_mnemonic(
                r#"This is a jukugo word, which usually means on'yomi readings from the kanji. If you know the readings of your kanji you'll know how to read this as well."#
            )
            .kanji_composition(vec![kanji_middle.symbol, kanji_stop.symbol])
            .build();

        repo.insert_vocabulary(vocabulary.clone()).await?;
        let vocab = repo.insert_vocabulary(vocabulary.clone()).await;

        assert!(matches!(vocab, Err(Error::Insert(InsertError::Conflict))));
        Ok(())
    }
}
