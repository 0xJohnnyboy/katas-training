package domain

import (
	"errors"
)

type Rental struct {
	movie      Movie
	daysRented float64
}

func NewRental(movie Movie, daysRented float64) *Rental {
	return &Rental{
		movie,
		daysRented,
	}
}

func (r Rental) Movie() Movie {
	return r.movie
}

func (r Rental) DaysRented() float64 {
	return r.daysRented
}

var CostCalculationMap = map[uint]func(r Rental) float64{
	REGULAR:     calculateRegularCost,
	NEW_RELEASE: calculateNewReleaseCost,
	CHILDREN:    calculateChildrenCost,
}

func calculateRegularCost(r Rental) float64 {
	amount := 2.0
	if r.DaysRented() > 2 {
		amount += (r.DaysRented() - 2) * 1.5
	}

	return amount
}

func calculateNewReleaseCost(r Rental) float64 {
	return r.DaysRented() * 3
}

func calculateChildrenCost(r Rental) float64 {
	amount := 1.5
	if r.DaysRented() > 3 {
		amount += (r.DaysRented() - 3) * 1.5
	}

	return amount
}

func (r Rental) Cost() (float64, error) {
	if calculus, ok := CostCalculationMap[r.Movie().PriceCode()]; ok {
		return calculus(r), nil
	}
	return 0, errors.New("Price code doesn't exist")
}

func (r Rental) RenterPoints() int {
	points := 1

	isNewRelease := r.Movie().PriceCode() == NEW_RELEASE
	isRentedMoreThanOneDay := r.DaysRented() > 1

	if isNewRelease && isRentedMoreThanOneDay {
		points++
	}

	return points
}
