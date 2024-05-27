-- Add allowed users to private decks
CREATE TABLE IF NOT EXISTS allowed_users_deck (
    deck_id UUID NOT NULL,
    user_id TEXT NOT NULL,
    CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT deck_id_fk FOREIGN KEY (deck_id) REFERENCES decks(id) ON DELETE CASCADE,
    CONSTRAINT allowed_users_deck_pk PRIMARY KEY (user_id, deck_id)
);