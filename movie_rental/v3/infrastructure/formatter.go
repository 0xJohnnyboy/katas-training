package infrastructure

import (
	domain "movie-rental-golang/domain"
)

type Formatter interface {
	Format(customer domain.Customer) string
}
