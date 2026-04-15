package movie_rental

type Formatter interface {
	FormatStatement(c Customer) string
}
