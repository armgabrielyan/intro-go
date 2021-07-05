package main

import "fmt"

func sum(numbers []float64) float64 {
	total := 0.0

	for _, num := range numbers {
		total += num
	}

	return total
}

func main() {
	numbers := []float64{1.2, 0.5, 12.9, 24.7, 6.3}

	total := sum(numbers)

	fmt.Println(total)
}
