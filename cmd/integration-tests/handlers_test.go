package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikayelaghasyan/card-deck-service/pkg/api"
	"github.com/mikayelaghasyan/card-deck-service/pkg/handler"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateDeck(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/decks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h, _ := handler.NewHandler()

	if assert.NoError(t, h.PostDecks(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		actual := api.CreateDeckResponse{}
		json.Unmarshal(rec.Body.Bytes(), &actual)

		assert.NotNil(t, uuid.FromStringOrNil(string(actual.DeckId)))
		assert.Equal(t, false, bool(actual.Shuffled))
		assert.Equal(t, 52, int(actual.Remaining))
	}
}
