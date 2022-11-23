use crate::{
    entities::vocabulary::{InsertVocabulary, Vocabulary},
    repositories::{vocabulary::VocabularyRepository, RepositoryError},
};

pub async fn execute<R>(db: &mut R, vocab: &InsertVocabulary) -> Result<Vocabulary, RepositoryError>
where
    R: VocabularyRepository,
{
    db.insert(vocab).await
}

#[cfg(test)]
mod tests {
    use crate::entities::{
        kanji::{kanji_middle, kanji_stop},
        radical::{radical_middle, radical_stop},
    };

    use super::*;

    use sqlx::PgPool;

    #[sqlx::test]
    async fn it_should_create_a_new_vocabulary(pool: PgPool) -> Result<(), RepositoryError> {
        let mut conn = pool.acquire().await?;

        let kanji_middle = {
            use crate::domain::{create_kanji, create_radical};
            let _ = create_radical::execute(&mut conn, &radical_middle()).await?;
            create_kanji::execute(&mut conn, &kanji_middle()).await?
        };
        let kanji_stop = {
            use crate::domain::{create_kanji, create_radical};
            let _ = create_radical::execute(&mut conn, &radical_stop()).await?;
            create_kanji::execute(&mut conn, &kanji_stop()).await?
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

        let created_radical = execute(&mut conn, &radical).await?;

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
    async fn it_should_collide_with_an_existing_vocabulary(
        pool: PgPool,
    ) -> Result<(), RepositoryError> {
        let mut conn = pool.acquire().await?;

        let kanji_middle = {
            use crate::domain::{create_kanji, create_radical};
            let _ = create_radical::execute(&mut conn, &radical_middle()).await?;
            create_kanji::execute(&mut conn, &kanji_middle()).await?
        };
        let kanji_stop = {
            use crate::domain::{create_kanji, create_radical};
            let _ = create_radical::execute(&mut conn, &radical_stop()).await?;
            create_kanji::execute(&mut conn, &kanji_stop()).await?
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

        let _ = execute(&mut conn, &radical).await?;
        let collision = execute(&mut conn, &radical).await;

        assert!(matches!(collision, Err(RepositoryError::Conflict)));

        Ok(())
    }
}
