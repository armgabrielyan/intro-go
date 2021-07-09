package main

import "fmt"

func main() {
	ptr := new(float64)
	*ptr = 6.5

	fmt.Println(ptr)
	fmt.Println(*ptr)
}
