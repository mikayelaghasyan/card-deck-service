package service

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/mikayelaghasyan/card-deck-service/pkg/model"
)

type DeckService struct {
}

func NewDeckService() (*DeckService, error) {
	rand.Seed(time.Now().UnixNano())
	return &DeckService{}, nil
}

func (service *DeckService) CreateDeck(shuffled bool, cards *[]model.Card) model.Deck {
	id, _ := uuid.NewRandom()
	cardList := newDefaultCardList()
	if shuffled {
		rand.Shuffle(len(cardList), func(i, j int) { cardList[i], cardList[j] = cardList[j], cardList[i] })
	}
	return model.Deck{
		Id:       id,
		Shuffled: shuffled,
		Cards:    cardList,
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
