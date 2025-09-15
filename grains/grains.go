package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains on a given square.
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("invalid number")
	}

	return uint64(math.Pow(2, float64(number-1))), nil
}

// Total returns the total number of grains on the chessboard.
func Total() uint64 {
	grainsCount, err := Square(64)
	if err != nil {
		return 0
	}

	return grainsCount*2 - 1
}
