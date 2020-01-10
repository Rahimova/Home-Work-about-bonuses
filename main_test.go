package main

import (
	"testing"
)

func TestBonusAmount(t *testing.T) {

	tests := []struct {
		name string
		sales []int
		want int
	}{

		{name: "Have bonus", sales: []int{15_000, 12_000, 11_000}, want: 400},
		{name: "No bonus", sales: []int{8000, 9000, 5, 4, 400}, want: 0},
		{name: "Equal to border - no bonus", sales: []int{10_000, 10_000, 10_000}, want: 0},

		{name: "Mixed", sales: []int{10_000, 15_000, 12_000, 8000}, want: 350},
	}

	for _, tests := range tests {


		got := bonusAmount(5, tests.sales)

		if got != tests.want {

			t.Error("bonusAmount test", tests.name, "got", got, "want", tests.want)

		}

	}

}
