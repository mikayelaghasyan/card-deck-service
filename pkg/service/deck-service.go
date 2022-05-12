package service

import "github.com/mikayelaghasyan/card-deck-service/pkg/model"

type DeckService struct {
}

func NewDeckService() (*DeckService, error) {
	return &DeckService{}, nil
}

func (service *DeckService) CreateDeck(shuffled bool, cards *[]model.Card) model.Deck {
	return model.Deck{}
}

func newDefaultCardList() (cards []model.Card) {
	var result []model.Card
	for suit := 1; suit <= 4; suit++ {
		for value := 1; suit <= 13; suit++ {
			result = append(result, model.Card{
				Suit:  model.CardSuit(suit),
				Value: model.CardValue(value),
			})
		}
	}
	return result
}
