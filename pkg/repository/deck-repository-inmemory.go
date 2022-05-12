package repository

import (
	"github.com/google/uuid"
	"github.com/mikayelaghasyan/card-deck-service/pkg/common"
	"github.com/mikayelaghasyan/card-deck-service/pkg/model"
)

type InMemoryDeckRepository struct {
	deckMap map[uuid.UUID]model.Deck
}

func NewInMemoryDeckRepository() (*InMemoryDeckRepository, error) {
	return &InMemoryDeckRepository{
		deckMap: map[uuid.UUID]model.Deck{},
	}, nil
}

func (repository *InMemoryDeckRepository) Save(deck model.Deck) (*model.Deck, error) {
	repository.deckMap[deck.Id] = deck
	return &deck, nil
}

func (repository *InMemoryDeckRepository) GetById(deckId uuid.UUID) (*model.Deck, error) {
	deck, exists := repository.deckMap[deckId]
	if exists {
		return &deck, nil
	} else {
		return nil, common.ErrNotFound
	}
}
