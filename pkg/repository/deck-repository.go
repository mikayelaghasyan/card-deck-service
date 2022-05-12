package repository

import (
	"github.com/google/uuid"
	"github.com/mikayelaghasyan/card-deck-service/pkg/model"
)

type DeckRepository interface {
	Save(model.Deck) (*model.Deck, error)
	GetById(deckId uuid.UUID) (*model.Deck, error)
}
