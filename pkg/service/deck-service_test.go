package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var deckService *DeckService

func setUp(t *testing.T) {
	service, err := NewDeckService()
	assert.NoError(t, err)
	deckService = service
}

func tearDown(t *testing.T) {

}

func TestCreateDeckDefault(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	deck := deckService.CreateDeck(false, nil)

	orderedCards := newDefaultCardList()

	assert.Equal(t, false, deck.Shuffled)
	assert.Equal(t, orderedCards, deck.Cards)
}

func TestCreateDeckShuffled(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	deck := deckService.CreateDeck(true, nil)

	orderedCards := newDefaultCardList()

	assert.Equal(t, true, deck.Shuffled)
	assert.Equal(t, len(orderedCards), len(deck.Cards))
	assert.NotEqual(t, orderedCards, deck.Cards)
}
