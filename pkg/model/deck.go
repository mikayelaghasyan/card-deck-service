package model

import "github.com/google/uuid"

type Deck struct {
	Id       uuid.UUID
	Shuffled bool
	Cards    []Card
}
