ALTER TABLE decks
ADD CONSTRAINT created_by_user_id_fk FOREIGN KEY (created_by_user_id) REFERENCES users(id) ON DELETE CASCADE;