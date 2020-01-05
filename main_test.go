package main

import "testing"

func Test_calculateTotalBonus(t *testing.T) {
	type args struct {
		sales []int
	}
	tests := []struct {
		name string
		sales []int
		totalBonus int
		want int
	}{
		// TODO: Add test cases.
	}
	for _, test := range tests{
		        got := calculateTotalBonus(test.sales);
		        if test.want != got {
				t.Error("want: ", test.want,  "but got: ", got)

			}
		}
	}




