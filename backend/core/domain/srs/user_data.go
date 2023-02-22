package srs

import "sixels.io/manekani/core/domain/cards"

type SRSUserData struct {
	Cards []*cards.Card `json:"cards"`
}
