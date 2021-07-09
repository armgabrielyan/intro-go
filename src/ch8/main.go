package main

import (
	"fmt"
	"intro-go/src/ch8/math"
)

func main() {
	numbers := []float64{1.6, 4.7, -7.6, 12.4, -9.2}

	average := math.Average(numbers)
	min := math.Min(numbers)
	max := math.Max(numbers)

	fmt.Printf("The average is %f\n", average)
	fmt.Printf("The min is %f\n", min)
	fmt.Printf("The max is %f\n", max)
}
