package serviceconfig_test

import (
	"testing"

	"github.com/kuzik/fizz-buzz/serviceconfig"
)

func TestAssertNumber(t *testing.T) {
	tests := []struct {
		name   string
		arg    int
		config []serviceconfig.NumberRule
		want   string
	}{
		{
			"1",
			1,
			[]serviceconfig.NumberRule{
				serviceconfig.NewRule(3, "Fizz"),
				serviceconfig.NewRule(5, "Buzz"),
			},
			"1",
		},
		{
			"Fizz",
			3,
			[]serviceconfig.NumberRule{
				serviceconfig.NewRule(3, "Fizz"),
				serviceconfig.NewRule(5, "Buzz"),
			},
			"Fizz",
		},
		{
			"Buzz",
			5,
			[]serviceconfig.NumberRule{
				serviceconfig.NewRule(3, "Fizz"),
				serviceconfig.NewRule(5, "Buzz"),
			},
			"Buzz",
		},
		{
			"Fizz Buzz",
			15,
			[]serviceconfig.NumberRule{
				serviceconfig.NewRule(3, "Fizz"),
				serviceconfig.NewRule(5, "Buzz"),
			},
			"Fizz Buzz",
		},
		{
			"Fizz Buzz",
			10,
			[]serviceconfig.NumberRule{
				serviceconfig.NewRule(3, "Fizz"),
				serviceconfig.NewRule(5, "Buzz"),
			},
			"Buzz",
		},
		{
			"Fizz Buzz",
			6,
			[]serviceconfig.NumberRule{
				serviceconfig.NewRule(3, "Fizz"),
				serviceconfig.NewRule(5, "Buzz"),
			},
			"Fizz",
		},
		{
			"Fizz Buzz",
			6,
			[]serviceconfig.NumberRule{
				serviceconfig.NewRule(1, "One"),
				serviceconfig.NewRule(2, "Two"),
				serviceconfig.NewRule(3, "Three"),
				serviceconfig.NewRule(6, "Four"),
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
		configChecker := serviceconfig.NewFizzBuzzChecker(tt.config)
		t.Run(tt.name, func(t *testing.T) {
			if got := configChecker.AssertNumber(tt.arg); got != tt.want {
				t.Errorf("serviceconfig() = %v, want %v", got, tt.want)
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
	configChecker := serviceconfig.NewFizzBuzzChecker([]serviceconfig.NumberRule{
		serviceconfig.NewRule(3, "Fizz"),
		serviceconfig.NewRule(5, "Buzz"),
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configChecker.CheckSequence(tt.arg); got != tt.want {
				t.Errorf("serviceconfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
