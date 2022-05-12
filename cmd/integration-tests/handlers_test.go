package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
	"github.com/mikayelaghasyan/card-deck-service/pkg/api"
	"github.com/mikayelaghasyan/card-deck-service/pkg/handler"
	"github.com/mikayelaghasyan/card-deck-service/pkg/repository"
	"github.com/mikayelaghasyan/card-deck-service/pkg/service"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

var hand *handler.Handler

func setUp(t *testing.T) {
	repo, err := repository.NewInMemoryDeckRepository()
	assert.NoError(t, err)
	service, err := service.NewDeckService(repo)
	assert.NoError(t, err)
	h, err := handler.NewHandler(*service)
	assert.NoError(t, err)
	hand = h
}

func tearDown(t *testing.T) {

}

func TestCreateDeckDefault(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	createDeckResponse := sendCreateDeckRequest(t, nil, nil)

	assert.NotNil(t, uuid.FromStringOrNil(string(createDeckResponse.DeckId)))
	assert.Equal(t, false, bool(createDeckResponse.Shuffled))
	assert.Equal(t, 52, int(createDeckResponse.Remaining))
}

func TestCreateDeckShuffled(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	shuffled := true
	response := sendCreateDeckRequest(t, &shuffled, nil)

	assert.NotNil(t, uuid.FromStringOrNil(string(response.DeckId)))
	assert.Equal(t, true, bool(response.Shuffled))
	assert.Equal(t, 52, int(response.Remaining))
}

func TestOpenDeckDefault(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	createDeckResponse := sendCreateDeckRequest(t, nil, nil)
	openDeckResponse := sendOpenDeckRequest(t, createDeckResponse.DeckId)

	orderedCards := createOrderedCards()

	expected := api.OpenDeckResponse{
		DeckBrief: api.DeckBrief{
			DeckId:    createDeckResponse.DeckId,
			Remaining: 52,
			Shuffled:  false,
		},
		Cards: api.Cards{
			Cards: orderedCards,
		},
	}
	assert.Equal(t, expected, openDeckResponse)

}

func sendCreateDeckRequest(t *testing.T, shuffled *bool, cards *[]api.CardCode) (response api.CreateDeckResponse) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/decks", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, hand.PostDecks(ctx, api.PostDecksParams{Shuffled: shuffled, Cards: cards})) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		response = api.CreateDeckResponse{}
		json.Unmarshal(rec.Body.Bytes(), &response)
	}

	return
}

func sendOpenDeckRequest(t *testing.T, deckId types.UUID) (response api.OpenDeckResponse) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/decks", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, hand.GetDecksId(ctx, api.DeckId(deckId))) {
		assert.Equal(t, http.StatusOK, rec.Code)

		response = api.OpenDeckResponse{}
		json.Unmarshal(rec.Body.Bytes(), &response)
	}

	return
}

func createOrderedCards() []api.Card {
	cardSuits := []api.CardSuit{
		api.CardSuitSPADES,
		api.CardSuitDIAMONDS,
		api.CardSuitCLUBS,
		api.CardSuitHEARTS,
	}
	cardValues := []api.CardValue{
		api.CardValueACE,
		api.CardValueN2,
		api.CardValueN3,
		api.CardValueN4,
		api.CardValueN5,
		api.CardValueN6,
		api.CardValueN7,
		api.CardValueN8,
		api.CardValueN9,
		api.CardValueN10,
		api.CardValueJACK,
		api.CardValueQUEEN,
		api.CardValueKING,
	}
	var orderedCards []api.Card
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			orderedCards = append(orderedCards, createCard(suit, value))
		}
	}
	return orderedCards
}

func createCard(suit api.CardSuit, value api.CardValue) api.Card {
	var code string
	if value == api.CardValueN10 {
		code += string(value)
	} else {
		code += string(value[:1])
	}
	code += string(suit[:1])

	return api.Card{
		Suit:  suit,
		Value: value,
		Code:  api.CardCode(code),
	}
}
