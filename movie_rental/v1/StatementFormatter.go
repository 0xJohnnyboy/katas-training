package movie_rental

import "fmt"

type StatementFormatter interface {
	Format(c Customer) string
}

type Format int

const (
	PlainText Format = iota
	HTML
)

type PlainTextFormatter struct{}

func (f PlainTextFormatter) Format(c Customer) string {
	totalAmount := 0.0

	result := "Rental Record for " + c.Name() + "\n"

	for _, rental := range c.rentals {
		rentalCost := rental.CalculateCost()

		result += fmt.Sprintf("\t"+rental.Movie().Title()+"\t%.1f\n", rentalCost)
		totalAmount += rentalCost
	}

	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %d frequent renter points", c.FrequentRenterPoints())

	return result
}

type HTMLFormatter struct{}

func (f HTMLFormatter) Format(c Customer) string {
	totalAmount := 0.0

	result := "<h1>Rental Record for <em>" + c.Name() + "</em></h1>"
	result += "<table>"

	for _, rental := range c.rentals {
		rentalCost := rental.CalculateCost()

		result += fmt.Sprintf("<tr><td>"+rental.Movie().Title()+"</td><td>%.1f</td></tr>", rentalCost)
		totalAmount += rentalCost
	}
	result += "</table>"

	result += fmt.Sprintf("<p>Amount owed is <em>%.1f</em></p>", totalAmount)
	result += fmt.Sprintf("<p>You earned <em>%d</em> frequent renter points</p>", c.FrequentRenterPoints())

	return result
}

var formatterMap = map[Format]StatementFormatter{
	PlainText: PlainTextFormatter{},
	HTML:      HTMLFormatter{},
}

func GetStatementFormatter(format Format) StatementFormatter {
	if formatter, ok := formatterMap[format]; ok {
		return formatter
	}
	return PlainTextFormatter{}
}
