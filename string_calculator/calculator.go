package string_calculator

import (
	"strconv"
	"strings"
	"errors"
)

var ErrNegativeNumber = errors.New("negative number not allowed")


func Add(input string) (int, error) {
	parsed := parse(input)

	if parsed == nil {
		return 0, nil
	}

	sum := 0
	for _, number := range parsed {
		if number < 0 {
			return 0, ErrNegativeNumber
		}

		sum += number
	}

	return sum, nil
}

func parse(input string) []int {
	input = strings.ReplaceAll(input, "\n", ",")
	splitted := strings.Split(input, ",")
	numbers := []int{}

	for _, number := range splitted {
		parsed, err := strconv.Atoi(number)
		if err != nil {
			return nil
		}
		numbers = append(numbers, parsed)
	}

	return numbers
}
