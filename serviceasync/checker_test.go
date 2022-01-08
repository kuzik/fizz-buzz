package serviceasync_test

import (
	"testing"

	"github.com/kuzik/fizz-buzz/serviceasync"
)

func TestAssertNumber(t *testing.T) {
	tests := []struct {
		name   string
		arg    int
		config []serviceasync.NumberRule
		want   string
	}{
		{
			"1",
			1,
			[]serviceasync.NumberRule{
				serviceasync.NewRule(3, "Fizz"),
				serviceasync.NewRule(5, "Buzz"),
			},
			"1",
		},
		{
			"Fizz",
			3,
			[]serviceasync.NumberRule{
				serviceasync.NewRule(3, "Fizz"),
				serviceasync.NewRule(5, "Buzz"),
			},
			"Fizz",
		},
		{
			"Buzz",
			5,
			[]serviceasync.NumberRule{
				serviceasync.NewRule(3, "Fizz"),
				serviceasync.NewRule(5, "Buzz"),
			},
			"Buzz",
		},
		{
			"Fizz Buzz",
			15,
			[]serviceasync.NumberRule{
				serviceasync.NewRule(3, "Fizz"),
				serviceasync.NewRule(5, "Buzz"),
			},
			"Fizz Buzz",
		},
		{
			"Fizz Buzz",
			10,
			[]serviceasync.NumberRule{
				serviceasync.NewRule(3, "Fizz"),
				serviceasync.NewRule(5, "Buzz"),
			},
			"Buzz",
		},
		{
			"Fizz Buzz",
			6,
			[]serviceasync.NumberRule{
				serviceasync.NewRule(3, "Fizz"),
				serviceasync.NewRule(5, "Buzz"),
			},
			"Fizz",
		},
		{
			"Fizz Buzz",
			6,
			[]serviceasync.NumberRule{
				serviceasync.NewRule(1, "One"),
				serviceasync.NewRule(2, "Two"),
				serviceasync.NewRule(3, "Three"),
				serviceasync.NewRule(6, "Four"),
			},
			"One Two Three Four",
		},
		{
			"Fizz Buzz",
			6,
			nil,
			"6",
		},
	}
	for _, tt := range tests {
		configChecker := serviceasync.NewFizzBuzzChecker(tt.config)
		t.Run(tt.name, func(t *testing.T) {

			if got := configChecker.AssertNumber(tt.arg); got != tt.want {
				t.Errorf("serviceasync() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckSequence(t *testing.T) {
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
		{
			"0",
			0,
			"",
		},
	}
	configChecker := serviceasync.NewFizzBuzzChecker([]serviceasync.NumberRule{
		serviceasync.NewRule(3, "Fizz"),
		serviceasync.NewRule(5, "Buzz"),
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configChecker.CheckSequence(tt.arg); got != tt.want {
				t.Errorf("serviceasync() = %v, want %v", got, tt.want)
			}
		})
	}
}
