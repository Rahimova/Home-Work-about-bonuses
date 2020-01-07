package main

import "fmt"

func main() {

	salary := []int{12000, 8000, 15000, 8000}

	fmt.Println(bonusAmount(5, salary))

}

func bonusAmount(percent int, salary []int) int {

	sumOfBonuses := 0

	for _, employee := range salary {

		bonusBorder := 10_000

		if employee > bonusBorder {

			bonus := employee - bonusBorder

			bonus = bonus * percent / 100

			sumOfBonuses = sumOfBonuses + bonus

		}

	}

	return sumOfBonuses

}
