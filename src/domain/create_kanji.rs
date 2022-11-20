use crate::{
    entities::kanji::{InsertKanji, Kanji},
    repositories::{kanji::KanjiRepository, RepositoryError},
};

pub async fn execute<K>(db: &mut K, kanji: &InsertKanji) -> Result<Kanji, RepositoryError>
where
    K: KanjiRepository,
{
    db.insert(kanji).await
}

#[cfg(test)]
mod tests {
    use crate::entities::radical::radical_barb;

    use super::*;

    use sqlx::PgPool;

    #[sqlx::test]
    async fn it_should_create_a_new_kanji(pool: PgPool) -> Result<(), RepositoryError> {
        let mut conn = pool.acquire().await?;

        let barb = {
            use crate::domain::create_radical;
            create_radical::execute(&mut conn, &radical_barb()).await?
        };

        let kanji = InsertKanji::builder()
            .name("Finish")
            .alt_names(vec!["Complete".to_owned(), "End".to_owned()])
            .symbol("了")
            .meaning_mnemonic(
                r#"Think about it this way. There is a famous fishing lure inventor, working on his best work ever. He’s old, and he’s been trying to make the best fishing lure for the last 50 years, and knows this will be his last invention. Then he figures it out. He takes one barb, and connects the end of it to the top of another barb. When he does this, he knows his work is finally finished."#
            )
            .reading("りょう")
            .onyomi(vec!["りょう".to_owned()])
            .reading_mnemonic(
                r#"How does the fish lure maker test his newly finished lure out? The one he just finished? He gets in a row (りょう) boat and rows out into the sea."#
            )
            .radical_composition(vec![barb.symbol])
            .build();

        let created_kanji = execute(&mut conn, &kanji).await?;

        assert_eq!(created_kanji.name, kanji.name);
        assert_eq!(created_kanji.alt_names, kanji.alt_names);
        assert_eq!(created_kanji.symbol, kanji.symbol);
        assert_eq!(created_kanji.meaning_mnemonic, kanji.meaning_mnemonic);
        assert_eq!(created_kanji.reading, kanji.reading);
        assert_eq!(created_kanji.onyomi, kanji.onyomi);
        assert_eq!(created_kanji.kunyomi, Vec::<String>::new());
        assert_eq!(created_kanji.nanori, Vec::<String>::new());
        assert_eq!(created_kanji.reading_mnemonic, kanji.reading_mnemonic);

        Ok(())
    }

    #[sqlx::test]
    async fn it_should_collide_with_an_existing_kanji(pool: PgPool) -> Result<(), RepositoryError> {
        let mut conn = pool.acquire().await?;

        let barb = {
            use crate::domain::create_radical;
            create_radical::execute(&mut conn, &radical_barb()).await?
        };

        let kanji = InsertKanji::builder()
            .name("finish")
            .alt_names(vec!["Complete".to_owned(), "End".to_owned()])
            .symbol("了")
            .meaning_mnemonic(
                r#"Think about it this way. There is a famous fishing lure inventor, working on his best work ever. He’s old, and he’s been trying to make the best fishing lure for the last 50 years, and knows this will be his last invention. Then he figures it out. He takes one barb, and connects the end of it to the top of another barb. When he does this, he knows his work is finally finished."#
            )
            .reading("りょう")
            .onyomi(vec!["りょう".to_owned()])
            .reading_mnemonic(
                r#"How does the fish lure maker test his newly finished lure out? The one he just finished? He gets in a row (りょう) boat and rows out into the sea."#
            )
            .radical_composition(vec![barb.symbol])
            .build();

        let _ = execute(&mut conn, &kanji).await?;
        let collision = execute(&mut conn, &kanji).await;

        assert!(matches!(collision, Err(RepositoryError::Conflict)));

        Ok(())
    }
}
