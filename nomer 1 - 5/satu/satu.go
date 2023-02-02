package main

import (
	"fmt"
	"math"
)

var denominations = []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

func calculateDenomination(amount int) map[int]int {
	result := make(map[int]int)
	for _, denomination := range denominations {
		if denomination <= amount {
			result[denomination] = int(math.Ceil(float64(amount) / float64(denomination)))
			amount %= denomination
		}
	}
	return result
}
func main() {
	var amount int
	fmt.Println("Masukan Inputan")
	fmt.Scanln(&amount)

	denominationMap := calculateDenomination(amount)

	output := ""
	for denomination, count := range denominationMap {
		output += fmt.Sprintf("\nRp. %d: %d\n", denomination, count)
	}

	result := []string{output}
	fmt.Println(result)
}
