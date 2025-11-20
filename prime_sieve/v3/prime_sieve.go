package v3

func Generate(end int) []int {
	primes := []int{}
	ch := generateNumbers(end)

	for {
		prime, ok := <-ch

		if !ok {
			break
		}

		primes = append(primes, prime)
		ch = filter(ch, prime)
	}

	return primes
}

func generateNumbers(end int) chan int {
	out := make(chan int)

	go func() {
		for i := 2; i <= end; i++ {
			out <- i
		}

		close(out)
	}()
	return out
}

func filter(in <-chan int, prime int) chan int {
	out := make(chan int)

	go func() {
		for i := range in {
			if i%prime == 0 {
				continue
			}

			out <- i
		}
		close(out)
	}()
	return out
}
