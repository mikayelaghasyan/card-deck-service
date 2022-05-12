package handler

import (
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
	"github.com/mikayelaghasyan/card-deck-service/pkg/api"
	"github.com/mikayelaghasyan/card-deck-service/pkg/service"
)

type Handler struct {
	deckService service.DeckService
}

func NewHandler(deckService service.DeckService) (*Handler, error) {
	return &Handler{
		deckService: deckService,
	}, nil
}

func (h *Handler) PostDecks(ctx echo.Context, params api.PostDecksParams) error {
	shuffled := false
	if params.Shuffled != nil {
		shuffled = *params.Shuffled
	}
	deck, err := h.deckService.CreateDeck(shuffled, nil)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "deck not found")
	}

	response := &api.CreateDeckResponse{
		DeckId:    types.UUID(deck.Id.String()),
		Shuffled:  deck.Shuffled,
		Remaining: api.NumberOfCards(len(deck.Cards)),
	}
	return ctx.JSON(http.StatusCreated, response)
}

func (h *Handler) GetDecksId(ctx echo.Context, deckId api.DeckId) error {
	return ctx.JSON(http.StatusOK, nil)
}
