package service

import (
	"testing"

	"github.com/mikayelaghasyan/card-deck-service/pkg/model"
	"github.com/mikayelaghasyan/card-deck-service/pkg/repository"
	"github.com/stretchr/testify/assert"
)

var deckService *DeckService

func setUp(t *testing.T) {
	repo, err := repository.NewInMemoryDeckRepository()
	assert.NoError(t, err)
	service, err := NewDeckService(repo)
	assert.NoError(t, err)
	deckService = service
}

func tearDown(t *testing.T) {

}

func TestCreateDeckDefault(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	deck, _ := deckService.CreateDeck(false, nil)

	orderedCards := newDefaultCardList()

	assert.Equal(t, false, deck.Shuffled)
	assert.Equal(t, orderedCards, deck.Cards)
}

func TestCreateDeckShuffled(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	deck, _ := deckService.CreateDeck(true, nil)

	orderedCards := newDefaultCardList()

	assert.Equal(t, true, deck.Shuffled)
	assert.Equal(t, len(orderedCards), len(deck.Cards))
	assert.NotEqual(t, orderedCards, deck.Cards)
}

func TestCreateDeckWithCards(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	sampleCards := []model.Card{
		model.NewCard(model.SPADES, model.ACE),
		model.NewCard(model.CLUBS, model.TEN),
		model.NewCard(model.DIAMONDS, model.TWO),
	}

	deck, _ := deckService.CreateDeck(true, &sampleCards)

	assert.Equal(t, false, deck.Shuffled)
	assert.Equal(t, sampleCards, deck.Cards)
}

func TestGetDeck(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	createdDeck, _ := deckService.CreateDeck(false, nil)

	deck := deckService.GetDeck(createdDeck.Id)

	assert.Equal(t, createdDeck, deck)
}
