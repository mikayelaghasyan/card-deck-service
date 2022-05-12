package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mikayelaghasyan/card-deck-service/pkg/api"
	"github.com/mikayelaghasyan/card-deck-service/pkg/handler"
	"github.com/mikayelaghasyan/card-deck-service/pkg/repository"
	"github.com/mikayelaghasyan/card-deck-service/pkg/service"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	e.File("/swagger/raw", "pkg/api/openapi.yml")
	url := echoSwagger.URL("/swagger/raw")
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(url))

	repo, err := repository.NewInMemoryDeckRepository()
	if err != nil {
		log.Fatal("failed to create repository")
	}

	deckService, err := service.NewDeckService(repo)
	if err != nil {
		log.Fatal("failed to create service")
	}

	handler, err := handler.NewHandler(*deckService)
	if err != nil {
		log.Fatal("failed to create handler")
	}

	g := e.Group("/api")
	api.RegisterHandlers(g, handler)

	e.Logger.Fatal(e.Start(":1323"))
}
