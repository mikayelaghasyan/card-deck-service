package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikayelaghasyan/card-deck-service/pkg/handler"
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
	}
}
