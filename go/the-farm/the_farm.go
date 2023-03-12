package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type ErrSillyNephew int

func (e ErrSillyNephew) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", int(e))
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}
	switch err {
	case ErrScaleMalfunction:
		if fodder < 0 {
			return 0.0, errors.New("negative fodder")
		} else if cows < 0 {
			return 0.0, ErrSillyNephew(cows)
		} else if fodder > 0 {
			return fodder * 2 / float64(cows), nil
		}
		return 0.0, err
	case nil:
		if fodder < 0 {
			return 0.0, errors.New("negative fodder")
		} else if cows < 0 {
			return 0.0, ErrSillyNephew(cows)
		} else if fodder > 0 {
			return fodder / float64(cows), err
		}
		return 0.0, nil
	default:
		return 0.0, err
	}
}
