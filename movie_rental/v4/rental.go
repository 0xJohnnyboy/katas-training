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

func (r Rental) Cost() float64 {
	cost := 0.0

	switch r.Movie().PriceCode() {
	case REGULAR:
		cost += 2
		if r.DaysRented() > 2 {
			cost += (r.DaysRented() - 2) * 1.5
		}
	case NEW_RELEASE:
		cost += r.DaysRented() * 3
	case CHILDREN:
		cost += 1.5
		if r.DaysRented() > 3 {
			cost += (r.DaysRented() - 3) * 1.5
		}
	}
	return cost
}

func (r Rental) RenterPoints() int {
	points := 1
	if (r.Movie().PriceCode() == NEW_RELEASE) && r.DaysRented() > 1 {
		points++
	}

	return points
}
