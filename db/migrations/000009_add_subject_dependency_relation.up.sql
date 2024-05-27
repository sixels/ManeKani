CREATE TABLE IF NOT EXISTS subjects_dependency (
    depends_on_id UUID NOT NULL,
    depended_by_id UUID NOT NULL,
    CONSTRAINT depends_on_id_fk FOREIGN KEY (depends_on_id) REFERENCES subjects(id) ON DELETE CASCADE,
    CONSTRAINT depended_by_id_fk FOREIGN KEY (depended_by_id) REFERENCES subjects(id) ON DELETE CASCADE,
    CONSTRAINT subject_dependency_pk PRIMARY KEY (depends_on_id, depended_by_id)
);