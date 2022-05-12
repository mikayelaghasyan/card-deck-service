package integrationtests

import (
	"testing"

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

	createDeckResponse := SendCreateDeckRequest(t, *hand, nil, nil)

	assert.NotNil(t, uuid.FromStringOrNil(string(createDeckResponse.DeckId)))
	assert.Equal(t, false, bool(createDeckResponse.Shuffled))
	assert.Equal(t, 52, int(createDeckResponse.Remaining))
}

func TestCreateDeckShuffled(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	shuffled := true
	response := SendCreateDeckRequest(t, *hand, &shuffled, nil)

	assert.NotNil(t, uuid.FromStringOrNil(string(response.DeckId)))
	assert.Equal(t, true, bool(response.Shuffled))
	assert.Equal(t, 52, int(response.Remaining))
}

func TestCreateDeckWithCards(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	sampleCards := CreateSampleCards()
	sampleCardCodes := ToCardCodes(sampleCards)
	createDeckResponse := SendCreateDeckRequest(t, *hand, nil, &sampleCardCodes)

	expectedCards := CreateSampleCards()
	expectedCardCodes := ToCardCodes(sampleCards)

	assert.NotNil(t, uuid.FromStringOrNil(string(createDeckResponse.DeckId)))
	assert.Equal(t, false, bool(createDeckResponse.Shuffled))
	assert.Equal(t, len(expectedCardCodes), int(createDeckResponse.Remaining))

	openDeckResponse := SendOpenDeckRequest(t, *hand, createDeckResponse.DeckId)

	assert.Equal(t, expectedCards, openDeckResponse.Cards.Cards)
}

func TestCreateDeckWithCardsShuffled(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	shuffled := true
	sampleCards := CreateSampleCards()
	sampleCardCodes := ToCardCodes(sampleCards)
	createDeckResponse := SendCreateDeckRequest(t, *hand, &shuffled, &sampleCardCodes)

	expectedCards := CreateSampleCards()
	expectedCardCodes := ToCardCodes(sampleCards)

	assert.NotNil(t, uuid.FromStringOrNil(string(createDeckResponse.DeckId)))
	assert.Equal(t, true, bool(createDeckResponse.Shuffled))
	assert.Equal(t, len(expectedCardCodes), int(createDeckResponse.Remaining))

	openDeckResponse := SendOpenDeckRequest(t, *hand, createDeckResponse.DeckId)

	assert.Equal(t, len(expectedCards), len(openDeckResponse.Cards.Cards))
	assert.NotEqual(t, expectedCards, openDeckResponse.Cards.Cards)
}

func TestOpenDeckDefault(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	createDeckResponse := SendCreateDeckRequest(t, *hand, nil, nil)
	openDeckResponse := SendOpenDeckRequest(t, *hand, createDeckResponse.DeckId)

	orderedCards := CreateOrderedCards()

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

func TestDrawCards(t *testing.T) {
	setUp(t)
	defer tearDown(t)

	shuffled := true
	createDeckResponse := SendCreateDeckRequest(t, *hand, &shuffled, nil)
	openDeckResponse := SendOpenDeckRequest(t, *hand, createDeckResponse.DeckId)

	expectedCards := openDeckResponse.Cards.Cards[:3]

	drawCardsResponse := SendDrawCardsRequest(t, *hand, createDeckResponse.DeckId, 3)

	assert.Equal(t, expectedCards, drawCardsResponse.Cards)
}
