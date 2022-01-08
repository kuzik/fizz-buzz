package basic_test

import (
	"testing"

	"github.com/kuzik/fizz-buzz/basic"
)

func TestAssertNumber(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want string
	}{
		{
			"1",
			1,
			"1",
		},
		{
			"Fizz",
			3,
			"Fizz",
		},
		{
			"Buzz",
			5,
			"Buzz",
		},
		{
			"Fizz Buzz",
			15,
			"Fizz Buzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basic.AssertNumber(tt.arg); got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want string
	}{
		{
			"5 len",
			5,
			"1,2,Fizz,4,Buzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basic.FizzBuzz(tt.arg); got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
