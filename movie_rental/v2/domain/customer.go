package domain

type Customer struct {
	name    string
	rentals []*Rental
}

func NewCustomer(name string) *Customer {
	return &Customer{name, []*Rental{}}
}

func (c Customer) Name() string {
	return c.name
}

func (c Customer) Rentals() []*Rental {
	return c.rentals
}

func (c *Customer) AddRental(r *Rental) {
	c.rentals = append(c.rentals, r)
}
