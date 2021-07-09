package main

import "fmt"

func max(numbers ...float64) float64 {
	m := numbers[0]

	for _, v := range numbers {
		if v > m {
			m = v
		}
	}

	return m
}

func main() {
	fmt.Println(max(4, 8.2, 5, 1, 6.7))
}
