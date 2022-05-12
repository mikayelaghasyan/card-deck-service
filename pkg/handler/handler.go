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

func (h *Handler) PostDecks(ctx echo.Context, params api.PostDecksParams) error {
	response := &api.CreateDeckResponse{
		DeckId:    types.UUID(uuid.NewString()),
		Shuffled:  false,
		Remaining: 52,
	}
	if params.Shuffled != nil {
		response.Shuffled = *params.Shuffled
	}
	return ctx.JSON(http.StatusCreated, response)
}

func (h *Handler) GetDecksId(ctx echo.Context, id api.DeckId) error {
	return ctx.JSON(http.StatusOK, nil)
}
