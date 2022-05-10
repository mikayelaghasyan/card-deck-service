package handler

import "github.com/labstack/echo/v4"

type Handler struct {
}

func NewHandler() (*Handler, error) {
	return &Handler{}, nil
}

func (h *Handler) PostDecks(c echo.Context) error {
	return nil
}
