use manekani_types::repository::{RepoInsertable, RepoQueryable};

use crate::entity::{
    kanji::GetKanji,
    vocabulary::{GetVocabulary, InsertVocabulary, Vocabulary, VocabularyPartial},
};

use super::Error;

pub async fn query<R>(repo: &R, vocab: GetVocabulary) -> Result<Vocabulary, Error>
where
    R: RepoQueryable<GetVocabulary, Vocabulary>,
{
    Ok(repo.query(vocab).await?)
}

pub async fn query_by_kanji<R>(repo: &R, kanji: GetKanji) -> Result<Vec<VocabularyPartial>, Error>
where
    R: RepoQueryable<GetKanji, Vec<VocabularyPartial>>,
{
    Ok(repo.query(kanji).await?)
}

pub async fn insert<R>(repo: &R, vocabulary: InsertVocabulary) -> Result<Vocabulary, Error>
where
    R: RepoInsertable<InsertVocabulary, Vocabulary>,
{
    Ok(repo.insert(vocabulary).await?)
}

#[cfg(test)]
mod tests {

    use crate::{
        entity::{
            kanji::{kanji_middle, kanji_stop},
            radical::{radical_middle, radical_stop},
        },
        repository::Repository,
    };

    use super::*;

    use sqlx::PgPool;

    #[sqlx::test]
    async fn it_should_create_a_new_vocabulary(pool: PgPool) -> Result<(), Error> {
        let repo = Repository::new(pool);

        let kanji_middle = {
            use crate::domain::{kanji, radical};
            let _ = radical::insert(&repo, radical_middle()).await?;
            kanji::insert(&repo, kanji_middle()).await?
        };
        let kanji_stop = {
            use crate::domain::{kanji, radical};
            let _ = radical::insert(&repo, radical_stop()).await?;
            kanji::insert(&repo, kanji_stop()).await?
        };

        let radical = InsertVocabulary::builder()
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

        let created_radical = insert(&repo, radical.clone()).await?;

        assert_eq!(created_radical.name, radical.name);
        assert_eq!(created_radical.word, radical.word);
        assert_eq!(created_radical.word_type, radical.word_type);
        assert_eq!(created_radical.reading, radical.reading);
        assert_eq!(created_radical.meaning_mnemonic, radical.meaning_mnemonic);
        assert_eq!(created_radical.reading_mnemonic, radical.reading_mnemonic);
        assert_eq!(created_radical.user_synonyms, None);
        assert_eq!(created_radical.user_meaning_note, None);

        Ok(())
    }

    #[sqlx::test]
    async fn it_should_collide_with_an_existing_vocabulary(pool: PgPool) -> Result<(), Error> {
        let repo = Repository::new(pool);

        let kanji_middle = {
            use crate::domain::{kanji, radical};
            let _ = radical::insert(&repo, radical_middle()).await?;
            kanji::insert(&repo, kanji_middle()).await?
        };
        let kanji_stop = {
            use crate::domain::{kanji, radical};
            let _ = radical::insert(&repo, radical_stop()).await?;
            kanji::insert(&repo, kanji_stop()).await?
        };
        let radical = InsertVocabulary::builder()
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

        let _ = insert(&repo, radical.clone()).await?;
        let collision = insert(&repo, radical.clone()).await;

        assert!(matches!(collision, Err(Error::Conflict)));

        Ok(())
    }
}
