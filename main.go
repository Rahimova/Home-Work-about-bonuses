package main

import "fmt"

func main()  {
		sales :=[]int{12_000, 8_000, 15_000, 12_000}
		totalBonus := calculateTotalBonus(sales)
		// Считаем общее количество бонусов
		fmt.Println("Общее количество бонусов" , totalBonus)

		//.....
}


	func calculateTotalBonus (sales []int) int {
		totalBonus := 0 // Общее количество бонусов
		for _, sales := range sales {
			if sales > 10_000 {
				upSales := sales - 10_000
				bonus := (upSales * 5) / 100
				fmt.Println(bonus)
				totalBonus += bonus
			}

		}
		return totalBonus
	}

