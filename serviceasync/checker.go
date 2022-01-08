package serviceasync

import (
	"strconv"
	"strings"
	"sync"
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
	AssertNumber(index int) string
}

type fizzBuzzChecker struct {
	rules []NumberRule
}

func NewFizzBuzzChecker(rules []NumberRule) Checker {
	return fizzBuzzChecker{rules: rules}
}

type assertResult struct {
	result string
	index  int
}

func (c fizzBuzzChecker) CheckSequence(max int) string {
	if max <= 0 {
		return ""
	}

	ch := make(chan assertResult)
	wg := &sync.WaitGroup{}

	for i := 1; i <= max; i++ {
		wg.Add(1)
		go func(ch chan assertResult, wg *sync.WaitGroup, index int) {
			defer wg.Done()

			ch <- assertResult{
				result: c.AssertNumber(index),
				index:  index,
			}
		}(ch, wg, i)
	}

	quit := make(chan bool)
	go func() {
		wg.Wait()
		quit <- true
	}()

	sorted := make([]string, max+1)
	for {
		select {
		case output := <-ch:
			sorted[output.index] = output.result
		case <-quit:
			return strings.Join(sorted, ",")[1:]
		}
	}
}

func (c fizzBuzzChecker) AssertNumber(index int) string {
	var res string
	for _, rule := range c.rules {
		if index%rule.number == 0 {
			res = res + " " + rule.fraze
		}
	}

	if len(res) != 0 {
		return res[1:]
	}

	return strconv.Itoa(index)
}
