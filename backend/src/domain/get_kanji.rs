use crate::{
    entities::kanji::{GetKanji, Kanji},
    repositories::{kanji::KanjiRepository, RepositoryError},
};

pub async fn execute<K>(db: &mut K, kanji: &GetKanji) -> Result<Kanji, RepositoryError>
where
    K: KanjiRepository,
{
    db.get(kanji).await
}
