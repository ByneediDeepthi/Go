package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(input string, n int) (int, error) {
	length := len(input)
	if n > length || n < 0 {
		return -1, errors.New("Invalid input value for n")
	}
	for _, d := range input {
		if d < '0' || d > '9' {
			return -1, errors.New("Invalid input for string")
		}
	}
	product := 1
	for i := 0; i < (n); i++ {
		x := string(input[i])
		y, _ := strconv.Atoi(x)
		product *= y
	}
	pro := product
	for i := 1; i <= (length - n); i++ {
		product := 1
		for j := 0; j < n; j++ {
			x := string(input[i+j])
			y, _ := strconv.Atoi(x)
			product *= y
		}
		if product > pro {
			pro = product
		}
	}
	return pro, nil
}
