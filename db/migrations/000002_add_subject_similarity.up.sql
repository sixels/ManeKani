CREATE TABLE IF NOT EXISTS subjects_similarity (
    similar_to_id UUID NOT NULL,
    similar_from_id UUID NOT NULL,
    CONSTRAINT similar_to_id_fk FOREIGN KEY (similar_to_id) REFERENCES subjects(id) ON DELETE CASCADE,
    CONSTRAINT similar_from_id_fk FOREIGN KEY (similar_from_id) REFERENCES subjects(id) ON DELETE CASCADE,
    CONSTRAINT subject_similarity_pk PRIMARY KEY (similar_to_id, similar_from_id)
);