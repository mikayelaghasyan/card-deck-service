package integrationtests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
	"github.com/mikayelaghasyan/card-deck-service/pkg/api"
	"github.com/mikayelaghasyan/card-deck-service/pkg/handler"
	"github.com/stretchr/testify/assert"
)

func SendCreateDeckRequest(t *testing.T, handler handler.Handler, shuffled *bool, cards *[]api.CardCode) (response api.CreateDeckResponse) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/decks", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, handler.PostDecks(ctx, api.PostDecksParams{Shuffled: shuffled, Cards: cards})) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		response = api.CreateDeckResponse{}
		json.Unmarshal(rec.Body.Bytes(), &response)
	}

	return
}

func SendOpenDeckRequest(t *testing.T, handler handler.Handler, deckId types.UUID) (response api.OpenDeckResponse) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/decks", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, handler.GetDecksId(ctx, api.DeckId(deckId))) {
		assert.Equal(t, http.StatusOK, rec.Code)

		response = api.OpenDeckResponse{}
		json.Unmarshal(rec.Body.Bytes(), &response)
	}

	return
}

func SendDrawCardsRequest(t *testing.T, handler handler.Handler, deckId types.UUID, count api.NumberOfCards) (response api.DrawCardsResponse) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/decks/:id/draw", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, handler.PutDecksIdDraw(ctx, api.DeckId(deckId), api.PutDecksIdDrawParams{Count: count})) {
		assert.Equal(t, http.StatusOK, rec.Code)

		response = api.DrawCardsResponse{}
		json.Unmarshal(rec.Body.Bytes(), &response)
	}

	return
}

func CreateOrderedCards() []api.Card {
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
			orderedCards = append(orderedCards, handler.NewApiCard(suit, value))
		}
	}
	return orderedCards
}

func CreateSampleCards() []api.Card {
	return []api.Card{
		handler.NewApiCard(api.CardSuitSPADES, api.CardValueACE),
		handler.NewApiCard(api.CardSuitCLUBS, api.CardValueN10),
		handler.NewApiCard(api.CardSuitDIAMONDS, api.CardValueN2),
	}
}

func ToCardCodes(cards []api.Card) []api.CardCode {
	cardCodes := []api.CardCode{}
	for _, card := range cards {
		cardCodes = append(cardCodes, card.Code)
	}
	return cardCodes
}
