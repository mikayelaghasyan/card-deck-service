package service

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/mikayelaghasyan/card-deck-service/pkg/model"
	"github.com/mikayelaghasyan/card-deck-service/pkg/repository"
)

type DeckService struct {
	repository repository.DeckRepository
}

func NewDeckService(repository repository.DeckRepository) (*DeckService, error) {
	rand.Seed(time.Now().UnixNano())
	return &DeckService{
		repository: repository,
	}, nil
}

func (service *DeckService) CreateDeck(shuffled bool, cards *[]model.Card) (*model.Deck, error) {
	id, _ := uuid.NewRandom()
	var cardList []model.Card
	if cards != nil {
		cardList = *cards
	} else {
		cardList = newDefaultCardList()
	}
	if shuffled {
		rand.Shuffle(len(cardList), func(i, j int) { cardList[i], cardList[j] = cardList[j], cardList[i] })
	}
	deck := model.Deck{
		Id:       id,
		Shuffled: shuffled,
		Cards:    cardList,
	}
	return service.repository.Save(deck)
}

func (service *DeckService) GetDeck(deckId uuid.UUID) *model.Deck {
	deck, _ := service.repository.GetById(deckId)
	return deck
}

func newDefaultCardList() (cards []model.Card) {
	var result []model.Card
	for suit := 1; suit <= 4; suit++ {
		for value := 1; value <= 13; value++ {
			result = append(result, model.NewCard(model.CardSuit(suit), model.CardValue(value)))
		}
	}
	return result
}
