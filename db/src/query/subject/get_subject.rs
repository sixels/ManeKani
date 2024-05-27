use sqlx::{Postgres, QueryBuilder};
use uuid::Uuid;

use crate::{model::subject::SubjectModel, Database};

/// Get a subject by its ID.
pub async fn get_subject(db: &Database, subject_id: Uuid) -> Result<SubjectModel, sqlx::Error> {
    let subjects = sqlx::query_as!(SubjectModel,
        r#"
        SELECT
            id,
            category,
            level,
            name,
            slug,
            priority,
            data,
            study_data,
            created_at,
            updated_at,
            deck_id,
            ARRAY_AGG(similars.similar_to_id) similars,
            ARRAY_AGG(dependencies.depends_on_id) depends_on,
            ARRAY_AGG(dependents.depended_by_id) depended_by
        FROM subjects
            INNER JOIN subjects_similarity AS similars ON subjects.id = similars.similar_to_id
            INNER JOIN subjects_dependency AS dependencies ON subjects.id = dependencies.depends_on_id
            INNER JOIN subjects_dependency AS dependents ON subjects.id = dependents.depended_by_id
        WHERE id = $1
        GROUP BY subjects.id
        "#,
        subject_id
    )
    .fetch_one(db)
    .await?;

    Ok(subjects)
}

pub struct GetDeckSubjectsFilter {
    pub categories: Option<Vec<String>>,
    pub levels: Option<Vec<i32>>,
    pub names: Option<Vec<String>>,
    pub slugs: Option<Vec<String>>,
    pub similar_to: Option<Uuid>,
}

pub async fn get_deck_subjects(
    db: &Database,
    deck_id: Uuid,
    filter: GetDeckSubjectsFilter,
) -> Result<Vec<SubjectModel>, sqlx::Error> {
    let mut query = QueryBuilder::<Postgres>::new(
        r#"
            SELECT
                id,
                category,
                level,
                name,
                slug,
                priority,
                data,
                study_data,
                created_at,
                updated_at,
                deck_id,
                ARRAY_AGG(similars.similar_to_id) similars,
                ARRAY_AGG(dependencies.depends_on_id) depends_on,
                ARRAY_AGG(dependents.depended_by_id) depended_by
            FROM subjects
            WHERE
                deck_id = $1
            "#,
    );
    query.push_bind(deck_id);

    let mut arg = 2;
    if let Some(categories) = filter.categories {
        query.push(format!(" AND category = ANY(${arg})"));
        query.push_bind(categories);
        arg += 1;
    }
    if let Some(levels) = filter.levels {
        query.push(format!(" AND level = ANY(${arg})"));
        query.push_bind(levels);
        arg += 1;
    }
    if let Some(names) = filter.names {
        query.push(format!(" AND name = ANY(${arg})"));
        query.push_bind(names);
        arg += 1
    }
    if let Some(slugs) = filter.slugs {
        query.push(format!(" AND slug = ANY(${arg})"));
        query.push_bind(slugs);
        arg += 1;
    }
    if let Some(similar_to) = filter.similar_to {
        query.push(format!(" AND id = ANY(SELECT similar_to_id FROM subjects_similarity WHERE similar_to_id = ${arg} OR similar_from_id = ${arg})"));
        query.push_bind(similar_to);
    }

    let subjects = query.build_query_as().fetch_all(db).await?;

    Ok(subjects)
}
