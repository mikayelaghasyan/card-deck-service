package common

import "errors"

var ErrGeneral = errors.New("something went wrong")
var ErrNotFound = errors.New("not found")
var ErrNotEnoughCards = errors.New("not enough cards in the deck")
