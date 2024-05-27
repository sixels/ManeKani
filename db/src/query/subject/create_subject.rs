use uuid::Uuid;

use crate::{model::subject::SubjectModel, Database};

pub struct CreateSubject {
    pub category: String,
    pub level: i32,
    pub name: String,
    pub slug: String,
    pub data: serde_json::Value,
    pub study_data: serde_json::Value,
    pub priority: i32,
    pub depends_on: Vec<Uuid>,
    pub depended_by: Vec<Uuid>,
    pub similars: Vec<Uuid>,
}

// TODO: error handling
pub async fn create_subject(
    db: &Database,
    subject: CreateSubject,
) -> Result<SubjectModel, sqlx::Error> {
    let mut tx = db.begin().await?;

    let result = sqlx::query!(
        r#"
        INSERT INTO subjects (
            category,
            level,
            name,
            slug,
            priority,
            data,
            study_data
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING *
        "#,
        subject.category,
        subject.level,
        subject.name,
        subject.slug,
        subject.priority,
        subject.data,
        subject.study_data
    )
    .fetch_one(&mut *tx)
    .await?;

    let depends_on = sqlx::query!(
        r#"
        INSERT INTO subjects_dependency (depended_by_id, depends_on_id)
            SELECT $1::UUID AS depended_by_id, * FROM UNNEST($2::UUID[]) AS depends_on_id
        RETURNING depended_by_id, depends_on_id
        "#,
        &result.id,
        &subject.depends_on,
    )
    .fetch_all(&mut *tx)
    .await?;

    let depended_by = sqlx::query!(
        r#"
        INSERT INTO subjects_dependency (depends_on_id, depended_by_id)
            SELECT $1::UUID AS depends_on_id, * FROM UNNEST($2::UUID[]) AS depended_by_id
        RETURNING depends_on_id, depended_by_id
        "#,
        &result.id,
        &subject.depended_by,
    )
    .fetch_all(&mut *tx)
    .await?;

    let similars = sqlx::query!(
        r#"
        INSERT INTO subjects_similarity (similar_from_id, similar_to_id)
            SELECT $1::UUID AS similar_from_id, * FROM UNNEST($2::UUID[]) AS similar_to_id
        RETURNING similar_to_id
        "#,
        &result.id,
        &subject.similars,
    )
    .fetch_all(&mut *tx)
    .await?;

    tx.commit().await?;

    let depends_on = depends_on
        .into_iter()
        .map(|d| d.depends_on_id)
        .collect::<Vec<_>>();

    let depended_by = depended_by
        .into_iter()
        .map(|d| d.depended_by_id)
        .collect::<Vec<_>>();

    let similars = similars
        .into_iter()
        .map(|s| s.similar_to_id)
        .collect::<Vec<_>>();

    Ok(SubjectModel {
        id: result.id,
        category: result.category,
        level: result.level,
        name: result.name,
        slug: result.slug,
        priority: result.priority,
        data: result.data,
        study_data: result.study_data,
        created_at: result.created_at,
        updated_at: result.updated_at,
        deck_id: result.deck_id,
        similars: Some(similars),
        depends_on: Some(depends_on),
        depended_by: Some(depended_by),
    })
}
