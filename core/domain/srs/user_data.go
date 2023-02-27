package srs

import "github.com/sixels/manekani/core/domain/cards"

type SRSUserData struct {
	Cards []*cards.Card `json:"cards"`
}
