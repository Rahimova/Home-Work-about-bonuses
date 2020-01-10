package main

import "fmt"

func main() {

	salary := []int{12000, 8000, 15000, 8000}

	fmt.Println(bonusAmount(5, salary))

}

func bonusAmount(percent int, sales []int) int {

	sumOfBonuses := 0

	for _, sale := range sales {

		bonusBorder := 10_000

		if sale > bonusBorder {

			bonus := sale - bonusBorder

			bonus = bonus * percent / 100

			sumOfBonuses = sumOfBonuses + bonus

		}

	}

	return sumOfBonuses

}
