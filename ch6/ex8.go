package main

import "fmt"

func main() {
	x := "Hello world!"
	add := &x

	y := "Bye world!"
	add = &y

	fmt.Println(add)
}
