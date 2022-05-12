package model

type Card struct {
	Suit  CardSuit
	Value CardValue
}

type CardSuit string

const (
	SPADES   = "S"
	DIAMONDS = "D"
	CLUBS    = "C"
	HEARTS   = "H"
)

type CardValue string

const (
	ACE   = "A"
	TWO   = "2"
	THREE = "3"
	FOUR  = "4"
	FIVE  = "5"
	SIX   = "6"
	SEVEN = "7"
	EIGHT = "8"
	NINE  = "9"
	TEN   = "10"
	JACK  = "J"
	QUEEN = "Q"
	KING  = "K"
)

func NewCard(suit CardSuit, value CardValue) Card {
	return Card{
		Suit:  suit,
		Value: value,
	}
}

func NewCardFromCode(code string) Card {
	idx := 1
	if len(code) == 3 {
		idx = 2
	}
	suit := CardSuit(code[idx:])
	value := CardValue(code[:idx])
	return NewCard(suit, value)
}
