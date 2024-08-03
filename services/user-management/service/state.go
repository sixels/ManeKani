package service

import "github.com/sixels/manekani/services/user-management/db"

type State struct {
	db *db.DB
}

// NewState creates a new instance of the user management service.
func NewState(db *db.DB) State {
	return State{
		db: db,
	}
}
