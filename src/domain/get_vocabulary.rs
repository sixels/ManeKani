use crate::{
    entities::vocabulary::{GetVocabulary, Vocabulary},
    repositories::{vocabulary::VocabularyRepository, RepositoryError},
};

pub async fn execute<K>(
    db: &mut K,
    vocabulary: &GetVocabulary,
) -> Result<Vocabulary, RepositoryError>
where
    K: VocabularyRepository,
{
    db.get(vocabulary).await
}
