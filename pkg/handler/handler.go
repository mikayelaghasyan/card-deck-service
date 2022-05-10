package handler

import (
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mikayelaghasyan/card-deck-service/pkg/api"
)

type Handler struct {
}

func NewHandler() (*Handler, error) {
	return &Handler{}, nil
}

func (h *Handler) PostDecks(c echo.Context) error {
	response := &api.CreateDeckResponse{
		DeckId:    types.UUID(uuid.NewString()),
		Shuffled:  false,
		Remaining: 52,
	}
	return c.JSON(http.StatusCreated, response)
}
