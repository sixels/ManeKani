use crate::{model::deck::DeckModel, Database};

pub struct CreateDeck {
    pub name: String,
    pub description: Option<String>,
    pub slug: String,
    pub image: Option<String>,
    pub tags: Vec<String>,
    pub is_featured: bool,
    pub visibility: DeckVisibility,
    pub created_by_user_id: String,
}

#[derive(Debug, PartialEq)]
pub enum DeckVisibility {
    Public,
    Private { allowed_users: Vec<String> },
}

pub async fn create_deck(db: &Database, deck: CreateDeck) -> Result<DeckModel, sqlx::Error> {
    let mut tx = db.begin().await?;

    let result = sqlx::query!(
        r#"
        INSERT INTO decks (
            name,
            description,
            slug,
            image,
            tags,
            is_featured,
            is_public,
            created_by_user_id
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING *
        "#,
        deck.name,
        deck.description,
        deck.slug,
        deck.image,
        &deck.tags,
        deck.is_featured,
        deck.visibility == DeckVisibility::Public,
        deck.created_by_user_id,
    )
    .fetch_one(&mut *tx)
    .await?;

    let mut allowed_users: Option<Vec<String>> = None;
    if let DeckVisibility::Private {
        allowed_users: users,
    } = deck.visibility
    {
        let allowed_users_result = sqlx::query!(
            r#"
            INSERT INTO allowed_users_deck (deck_id, user_id)
                SELECT $1::UUID AS deck_id, * FROM UNNEST($2::TEXT[]) AS user_id
            RETURNING user_id
            "#,
            &result.id,
            &users,
        )
        .fetch_all(&mut *tx)
        .await?;

        allowed_users = Some(
            allowed_users_result
                .into_iter()
                .map(|row| row.user_id)
                .collect(),
        )
    }

    tx.commit().await?;

    Ok(DeckModel {
        id: result.id,
        name: result.name,
        slug: result.slug,
        image: result.image,
        description: result.description,
        tags: result.tags,
        is_featured: result.is_featured,
        is_public: result.is_public,
        created_at: result.created_at,
        updated_at: result.updated_at,
        created_by_user_id: result.created_by_user_id,
        allowed_users,
    })
}
