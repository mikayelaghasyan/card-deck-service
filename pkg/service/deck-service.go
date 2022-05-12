package service

import (
	"github.com/google/uuid"
	"github.com/mikayelaghasyan/card-deck-service/pkg/model"
)

type DeckService struct {
}

func NewDeckService() (*DeckService, error) {
	return &DeckService{}, nil
}

func (service *DeckService) CreateDeck(shuffled bool, cards *[]model.Card) model.Deck {
	id, _ := uuid.NewRandom()
	return model.Deck{
		Id:       id,
		Shuffled: shuffled,
		Cards:    newDefaultCardList(),
	}
}

func newDefaultCardList() (cards []model.Card) {
	var result []model.Card
	for suit := 1; suit <= 4; suit++ {
		for value := 1; value <= 13; value++ {
			result = append(result, model.Card{
				Suit:  model.CardSuit(suit),
				Value: model.CardValue(value),
			})
		}
	}
	return result
}
