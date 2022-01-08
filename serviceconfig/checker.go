package serviceconfig

import (
	"strconv"
	"strings"
)

type NumberRule struct {
	number int
	fraze  string
}

func NewRule(number int, fraze string) NumberRule {
	return NumberRule{number: number, fraze: fraze}
}

type Checker interface {
	CheckSequence(max int) string
	AssertNumber(n int) string
}

type fizzBuzzChecker struct {
	rules []NumberRule
}

func NewFizzBuzzChecker(rules []NumberRule) Checker {
	return fizzBuzzChecker{rules: rules}
}

func (c fizzBuzzChecker) CheckSequence(max int) string {
	res := make([]string, 0, max)
	for i := 1; i <= max; i++ {
		res = append(res, c.AssertNumber(i))
	}

	return strings.Join(res, ",")
}

func (c fizzBuzzChecker) AssertNumber(n int) string {
	var res string
	for _, rule := range c.rules {
		if n%rule.number == 0 {
			res = res + " " + rule.fraze
		}
	}

	if len(res) != 0 {
		return res[1:]
	}

	return strconv.Itoa(n)
}
