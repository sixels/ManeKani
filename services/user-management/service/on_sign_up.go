package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/sixels/manekani/services/user-management/db"
)

type orySignUpPayload struct {
	Flow struct {
		ID string `json:"id"`
	} `json:"flow"`
	Identity struct {
		Traits struct {
			Email string `json:"email"`
		} `json:"traits"`
		MetadataPublic *oryMetadataPublic `json:"metadata_public,omitempty"`
	} `json:"identity"`
}

type oryMetadataPublic struct {
	ID string `json:"id"`
}

func (s *Service) onSignUp(w http.ResponseWriter, r *http.Request) {
	var payload orySignUpPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := payload.Identity.Traits.Email
	log.Printf("received on-sign-up event for user %s\n", email)

	// attempt to create a new user
	user, err := s.state.db.CreateUser(r.Context(), db.CreateUserInput{
		Email: email,
	})

	if errors.Is(err, db.ErrCreateUserDuplicateEmail) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(
			fmt.Sprintf(`{"messages":[{"instance_ptr":"#/traits/email","messages":[{"id":123,"text":"email already in use","type":"validation","context":{"value":"%v"}}]}]}`,
				email)))
		return
	}

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload.Identity.MetadataPublic = &oryMetadataPublic{
		ID: user.ID.String(),
	}
	jsonResponse(w, http.StatusOK, payload)
}

func jsonResponse[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
