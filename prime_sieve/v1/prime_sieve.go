package v1

func GetPrimeNumbers(n int) []int {

	primes := []int{}
	pipe := GenerateNumbers(n)
	for {
		prime, ok := <-pipe
		if !ok {
			break
		}

		primes = append(primes, prime)
		pipe = Filter(pipe, prime)
	}

	return primes
}

func GenerateNumbers(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 2; i <= n; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func Filter(in <-chan int, n int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			if v%n == 0 {
				continue
			}
			out <- v
		}
		close(out)
	}()
	return out
}
