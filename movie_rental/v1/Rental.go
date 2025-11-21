package movie_rental

type Rental struct {
	movie      Movie
	daysRented float64
}

func (r Rental) Movie() Movie {
	return r.movie
}

func (r Rental) DaysRented() float64 {
	return r.daysRented
}

func (r Rental) CalculateFrequentRenterPoints() uint {
	pc := r.Movie().PriceCode()
	strategy := GetPricingStrategy(pc)

	return strategy.CalculateFrequentRenterPoints(r.DaysRented())
}

func (r Rental) CalculateCost() float64 {
	pc := r.Movie().PriceCode()

	strategy := GetPricingStrategy(pc)
	return strategy.CalculateCost(r.DaysRented())
}
