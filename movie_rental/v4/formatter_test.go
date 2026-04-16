package movie_rental

import "testing"

func makeCustomer(name string, rentals []Rental) Customer {
	customer := Customer{name, []Rental{}}
	for _, rental := range rentals {
		customer.AddRental(rental)
	}
	return customer
}

func TestFormatterRegistry(t *testing.T) {
	if _, ok := FormatterRegistry[PlainText]; !ok {
		t.Fatalf("expected formatter registry to contain PlainText formatter")
	}
	if _, ok := FormatterRegistry[HTML]; !ok {
		t.Fatalf("expected formatter registry to contain HTML formatter")
	}
}

func TestFormatters(t *testing.T) {
	type formatterExpectation struct {
		name string
		want string
	}

	type statementCase struct {
		name    string
		customer Customer
		expects []formatterExpectation
	}

	cases := []statementCase{
		{
			name: "should return statement for mixed rentals",
			customer: makeCustomer("Bob", []Rental{
				{Movie{"Jaws", REGULAR}, 2},
				{Movie{"Golden Eye", REGULAR}, 3},
				{Movie{"Short New", NEW_RELEASE}, 1},
				{Movie{"Long New", NEW_RELEASE}, 2},
				{Movie{"Bambi", CHILDREN}, 3},
				{Movie{"Toy Story", CHILDREN}, 4},
			}),
			expects: []formatterExpectation{
				{
					name: "plain text",
					want: "" +
						"Rental Record for Bob\n" +
						"\tJaws\t2.0\n" +
						"\tGolden Eye\t3.5\n" +
						"\tShort New\t3.0\n" +
						"\tLong New\t6.0\n" +
						"\tBambi\t1.5\n" +
						"\tToy Story\t3.0\n" +
						"Amount owed is 19.0\n" +
						"You earned 7 frequent renter points",
				},
				{
					name: "html",
					want: "" +
						"<h1>Rental Record for <em>Bob</em></h1>\n" +
						"<table>\n" +
						"<tr><td>Jaws</td><td>2.0</td></tr>\n" +
						"<tr><td>Golden Eye</td><td>3.5</td></tr>\n" +
						"<tr><td>Short New</td><td>3.0</td></tr>\n" +
						"<tr><td>Long New</td><td>6.0</td></tr>\n" +
						"<tr><td>Bambi</td><td>1.5</td></tr>\n" +
						"<tr><td>Toy Story</td><td>3.0</td></tr>\n" +
						"</table>\n" +
						"<p>Amount owed is <em>19.0</em></p>\n" +
						"<p>You earned <em>7</em> frequent renter points</p>",
				},
			},
		},
		{
			name:     "should return zero totals when customer has no rentals",
			customer: makeCustomer("Alice", []Rental{}),
			expects: []formatterExpectation{
				{
					name: "plain text",
					want: "" +
						"Rental Record for Alice\n" +
						"Amount owed is 0.0\n" +
						"You earned 0 frequent renter points",
				},
				{
					name: "html",
					want: "" +
						"<h1>Rental Record for <em>Alice</em></h1>\n" +
						"<table>\n" +
						"</table>\n" +
						"<p>Amount owed is <em>0.0</em></p>\n" +
						"<p>You earned <em>0</em> frequent renter points</p>",
				},
			},
		},
		{
			name: "should format fractional costs with one decimal and apply new release bonus points over one day",
			customer: makeCustomer("Charlie", []Rental{
				{Movie{"Almost Three Days", REGULAR}, 2.5},
				{Movie{"One Day Plus", NEW_RELEASE}, 1.5},
			}),
			expects: []formatterExpectation{
				{
					name: "plain text",
					want: "" +
						"Rental Record for Charlie\n" +
						"\tAlmost Three Days\t2.8\n" +
						"\tOne Day Plus\t4.5\n" +
						"Amount owed is 7.2\n" +
						"You earned 3 frequent renter points",
				},
				{
					name: "html",
					want: "" +
						"<h1>Rental Record for <em>Charlie</em></h1>\n" +
						"<table>\n" +
						"<tr><td>Almost Three Days</td><td>2.8</td></tr>\n" +
						"<tr><td>One Day Plus</td><td>4.5</td></tr>\n" +
						"</table>\n" +
						"<p>Amount owed is <em>7.2</em></p>\n" +
						"<p>You earned <em>3</em> frequent renter points</p>",
				},
			},
		},
	}

	type formatterTarget struct {
		name string
		kind FormatterType
	}

	targets := []formatterTarget{
		{name: "plain text", kind: PlainText},
		{name: "html", kind: HTML},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			for _, target := range targets {
				target := target
				t.Run(target.name, func(t *testing.T) {
					formatter := FormatterRegistry[target.kind]
					actual := formatter.FormatStatement(tc.customer)

					var expected string
					for _, exp := range tc.expects {
						if exp.name == target.name {
							expected = exp.want
							break
						}
					}

					if expected == "" {
						t.Fatalf("missing expected output for formatter %q in test %q", target.name, tc.name)
					}

					if expected != actual {
						t.Fatalf("Expected:\n%s\nActual:\n%s\n", expected, actual)
					}
				})
			}
		})
	}
}
