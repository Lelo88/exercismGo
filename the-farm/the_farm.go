package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	numberOfCows int
	message      string
}
func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.message)
}

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, cows int) (float64, error) {

	amount, fodderAmountErr := f.FodderAmount(cows)
	if fodderAmountErr != nil {
		return 0, fodderAmountErr
	}
	fatteningFactor, fatteningFactorErr := f.FatteningFactor()
	if fatteningFactorErr != nil {
		return 0, fatteningFactorErr
	}
	return (amount * fatteningFactor) / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(f, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return &InvalidCowsError{
			numberOfCows:  cows,
			message:      "there are no negative cows",
		}
	case cows == 0:
		return &InvalidCowsError{
			numberOfCows: cows,
			message:      "no cows don't need food",
		}
	default:
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
