package movie_rental

type Customer struct {
	name    string
	rentals []Rental
}

func (c *Customer) AddRental(rental Rental) {
	c.rentals = append(c.rentals, rental)
}

func (c Customer) Name() string {
	return c.name
}
