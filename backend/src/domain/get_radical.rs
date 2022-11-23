use crate::{
    entities::radical::{GetRadical, Radical},
    repositories::{radical::RadicalRepository, RepositoryError},
};

pub async fn execute<K>(db: &mut K, radical: &GetRadical) -> Result<Radical, RepositoryError>
where
    K: RadicalRepository,
{
    db.get(radical).await
}
