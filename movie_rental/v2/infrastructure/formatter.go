package infrastructure

import (
	domain "movie-rental-golang/domain"
)

type StatementFormat uint

const (
	PlainText StatementFormat = iota
	HTML
)

var FormattersRegistry = map[StatementFormat]func() StatementFormatter{
	PlainText: func() StatementFormatter { return NewPlainTextFormatter() },
	HTML:      func() StatementFormatter { return NewHTMLFormatter() },
}

type StatementFormatter interface {
	Format(c *domain.Customer) (string, error)
}
