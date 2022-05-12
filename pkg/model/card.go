package model

type Card struct {
	Suit  CardSuit
	Value CardValue
}

type CardSuit int8

const (
	_ CardSuit = iota
	SPADES
	DIAMONDS
	CLUBS
	HEARTS
)

type CardValue int8

const (
	_ CardValue = iota
	ACE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)
