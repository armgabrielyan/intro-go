package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := 1
	y := 2

	fmt.Printf("x = %d | y = %d\n", x, y)

	swap(&x, &y)

	fmt.Printf("x = %d | y = %d", x, y)
}
