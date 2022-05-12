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

	orderedCards := deckService.newDefaultCardList()

	assert.Equal(t, false, deck.Shuffled)
	assert.Equal(t, orderedCards, deck.Cards)
}

func TestCreateDeckShuffled(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	deck, _ := deckService.CreateDeck(true, nil)

	orderedCards := deckService.newDefaultCardList()

	assert.Equal(t, true, deck.Shuffled)
	assert.Equal(t, len(orderedCards), len(deck.Cards))
	assert.NotEqual(t, orderedCards, deck.Cards)
}

func TestCreateDeckWithCards(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	expectedCards := createSampleCards()

	deck, _ := deckService.CreateDeck(false, createSampleCards())

	assert.Equal(t, false, deck.Shuffled)
	assert.Equal(t, expectedCards, deck.Cards)
}

func TestCreateDeckWithCardsShuffled(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	expectedCards := createSampleCards()

	deck, _ := deckService.CreateDeck(true, createSampleCards())

	assert.Equal(t, true, deck.Shuffled)
	assert.Equal(t, len(expectedCards), len(deck.Cards))
	assert.NotEqual(t, expectedCards, deck.Cards)
}

func TestGetDeck(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	createdDeck, _ := deckService.CreateDeck(false, nil)

	deck, _ := deckService.GetDeck(createdDeck.Id)

	assert.Equal(t, createdDeck, deck)
}

func TestDrawCards(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	createdDeck, _ := deckService.CreateDeck(true, nil)
	drawnCards, _ := deckService.DrawCards(createdDeck.Id, 3)

	assert.Equal(t, createdDeck.Cards[:3], drawnCards)
}

func createSampleCards() []model.Card {
	return []model.Card{
		model.NewCard(model.SPADES, model.ACE),
		model.NewCard(model.CLUBS, model.TEN),
		model.NewCard(model.DIAMONDS, model.TWO),
		model.NewCard(model.DIAMONDS, model.FIVE),
		model.NewCard(model.HEARTS, model.JACK),
		model.NewCard(model.HEARTS, model.KING),
		model.NewCard(model.DIAMONDS, model.ACE),
	}
}
