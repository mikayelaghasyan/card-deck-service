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
	cardSuits  []model.CardSuit
	cardValues []model.CardValue
}

func NewDeckService(repository repository.DeckRepository) (*DeckService, error) {
	rand.Seed(time.Now().UnixNano())
	return &DeckService{
		repository: repository,
		cardSuits: []model.CardSuit{
			model.SPADES,
			model.DIAMONDS,
			model.CLUBS,
			model.HEARTS,
		},
		cardValues: []model.CardValue{
			model.ACE,
			model.TWO,
			model.THREE,
			model.FOUR,
			model.FIVE,
			model.SIX,
			model.SEVEN,
			model.EIGHT,
			model.NINE,
			model.TEN,
			model.JACK,
			model.QUEEN,
			model.KING,
		},
	}, nil
}

func (service *DeckService) CreateDeck(shuffled bool, cards []model.Card) (*model.Deck, error) {
	id, _ := uuid.NewRandom()
	var cardList []model.Card
	if cards != nil {
		cardList = cards
	} else {
		cardList = service.newDefaultCardList()
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

func (service *DeckService) newDefaultCardList() (cards []model.Card) {
	var result []model.Card
	for _, suit := range service.cardSuits {
		for _, value := range service.cardValues {
			result = append(result, model.NewCard(suit, value))
		}
	}
	return result
}
