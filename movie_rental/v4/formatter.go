package movie_rental

type Formatter interface {
	FormatStatement(c Customer) string
}

type FormatterType int

const (
	PlainText FormatterType = iota
	HTML
)

var FormatterRegistry map[FormatterType]Formatter = map[FormatterType]Formatter{
	PlainText: NewPlainTextFormatter(),
	HTML: NewHTMLFormatter(),
}
