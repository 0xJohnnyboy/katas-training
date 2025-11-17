package main

import (
	"fmt"

	ps "katas/prime_sieve/v1"
)

func main() {
	primes := ps.GetPrimeNumbers(100)
	for _, p := range primes {
		fmt.Println(p)
	}
}
