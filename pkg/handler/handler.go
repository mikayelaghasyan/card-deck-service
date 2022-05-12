package handler

import (
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mikayelaghasyan/card-deck-service/pkg/api"
	"github.com/mikayelaghasyan/card-deck-service/pkg/model"
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

	var cards []model.Card
	if params.Cards != nil {
		cardCodes := *params.Cards
		cards = toModelCards(cardCodes)
	}

	deck, err := h.deckService.CreateDeck(shuffled, cards)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not create deck")
	}

	response := &api.CreateDeckResponse{
		DeckId:    types.UUID(deck.Id.String()),
		Shuffled:  deck.Shuffled,
		Remaining: api.NumberOfCards(len(deck.Cards)),
	}
	return ctx.JSON(http.StatusCreated, response)
}

func (h *Handler) GetDecksId(ctx echo.Context, deckId api.DeckId) error {
	id, _ := uuid.Parse(string(deckId))
	deck := h.deckService.GetDeck(id)
	if deck == nil {
		return echo.NewHTTPError(http.StatusNotFound, "deck not found")
	}

	response := api.OpenDeckResponse{
		DeckBrief: api.DeckBrief{
			DeckId:    types.UUID(deck.Id.String()),
			Shuffled:  deck.Shuffled,
			Remaining: api.NumberOfCards(len(deck.Cards)),
		},
		Cards: api.Cards{
			Cards: toApiCards(deck.Cards),
		},
	}
	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) PutDecksIdDraw(ctx echo.Context, id api.DeckId, params api.PutDecksIdDrawParams) error {
	return nil
}

func NewApiCard(suit api.CardSuit, value api.CardValue) api.Card {
	return api.Card{
		Suit:  suit,
		Value: value,
		Code:  cardCode(suit, value),
	}
}

func toApiCards(cardList []model.Card) []api.Card {
	result := []api.Card{}
	for _, card := range cardList {
		result = append(result, toApiCard(card))
	}
	return result
}

func toApiCard(card model.Card) api.Card {
	suit := toApiCardSuit(card.Suit)
	value := toApiCardValue(card.Value)
	return NewApiCard(suit, value)
}

func toApiCardSuit(suit model.CardSuit) api.CardSuit {
	switch suit {
	case model.SPADES:
		return api.CardSuitSPADES
	case model.DIAMONDS:
		return api.CardSuitDIAMONDS
	case model.CLUBS:
		return api.CardSuitCLUBS
	case model.HEARTS:
		return api.CardSuitHEARTS
	}
	panic("unknown card suit: " + suit)
}

func toApiCardValue(value model.CardValue) api.CardValue {
	switch value {
	case model.ACE:
		return api.CardValueACE
	case model.TWO:
		return api.CardValueN2
	case model.THREE:
		return api.CardValueN3
	case model.FOUR:
		return api.CardValueN4
	case model.FIVE:
		return api.CardValueN5
	case model.SIX:
		return api.CardValueN6
	case model.SEVEN:
		return api.CardValueN7
	case model.EIGHT:
		return api.CardValueN8
	case model.NINE:
		return api.CardValueN9
	case model.TEN:
		return api.CardValueN10
	case model.JACK:
		return api.CardValueJACK
	case model.QUEEN:
		return api.CardValueQUEEN
	case model.KING:
		return api.CardValueKING
	}
	panic("unknown card value: " + value)
}

func cardCode(suit api.CardSuit, value api.CardValue) api.CardCode {
	var code string
	if value == api.CardValueN10 {
		code += string(value)
	} else {
		code += string(value[:1])
	}
	code += string(suit[:1])
	return api.CardCode(code)
}

func toModelCards(cardCodes []api.CardCode) []model.Card {
	cards := []model.Card{}
	for _, code := range cardCodes {
		cards = append(cards, model.NewCardFromCode(string(code)))
	}
	return cards
}
