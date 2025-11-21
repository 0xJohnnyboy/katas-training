package movie_rental

type Customer struct {
	name                 string
	rentals              []Rental
	frequentRenterPoints uint
}

func NewCustomer(name string) *Customer {
	return &Customer{name, []Rental{}, 0}
}

func (c *Customer) addRental(rental Rental) {
	c.frequentRenterPoints += rental.CalculateFrequentRenterPoints()
	c.rentals = append(c.rentals, rental)
}

func (c Customer) FrequentRenterPoints() uint {
	return c.frequentRenterPoints
}

func (c Customer) Name() string {
	return c.name
}

// deprecated
func (c Customer) Statement() string {
	formatter := GetStatementFormatter(PlainText)
	return formatter.Format(c)
}
