package domain

import (
	"fmt"
)

type Rental struct {
	movie      *Movie
	daysRented uint
}

func NewRental(movie *Movie, daysRented uint) *Rental {
	return &Rental{movie, daysRented}
}

func (r Rental) Movie() *Movie {
	return r.movie
}

func (r Rental) DaysRented() uint {
	return r.daysRented
}

var CostCalculationMap = map[PriceCode]func(r Rental) float64{
	REGULAR:     calculateRegularCost,
	NEW_RELEASE: calculateNewReleaseCost,
	CHILDREN:    calculateChildrenCost,
}

func calculateRegularCost(r Rental) float64 {
	var cost float64
	cost = 2
	if r.daysRented > 2 {
		cost += float64(r.DaysRented()-2) * 1.5
	}

	return cost
}

func calculateNewReleaseCost(r Rental) float64 {
	var cost float64
	cost = float64(r.daysRented) * 3
	return cost
}

func calculateChildrenCost(r Rental) float64 {
	var cost float64
	cost = 1.5
	if r.daysRented > 3 {
		cost += float64(r.daysRented-3) * 1.5
	}
	return cost
}

func (r Rental) Cost() (float64, error) {
	pc := r.Movie().PriceCode()
	fn, ok := CostCalculationMap[pc]
	if !ok {
		return 0, fmt.Errorf("unknown price code: %d", pc)
	}
	return fn(r), nil
}

func (r Rental) RenterPoints() uint {

	if r.Movie().PriceCode() == NEW_RELEASE && r.daysRented > 1 {
		return 2
	}

	return 1
}
