package main

import (
	"fmt"
	"testing"
)

func TestFuelRequired(t *testing.T) {
	tt := []struct {
		input int
		want  int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("test-%d", tc.input), func(t *testing.T) {
			if want, got := tc.want, fuelRequired(tc.input); want != got {
				t.Fatalf("wanted %d, but got %d", want, got)
			}
		})
	}
}
