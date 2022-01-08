package basic

import (
	"strconv"
	"strings"
)

func FizzBuzz(max int) string {
	res := make([]string, 0, max)
	for i := 1; i <= max; i++ {
		res = append(res, AssertNumber(i))
	}

	return strings.Join(res, ",")
}

func AssertNumber(n int) string {
	var res string

	if n%3 == 0 {
		res += " Fizz"
	}
	if n%5 == 0 {
		res += " Buzz"
	}

	if len(res) != 0 {
		return res[1:]
	}

	return strconv.Itoa(n)
}
