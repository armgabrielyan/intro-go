package main

import (
	"fmt"
	"intro-go/src/ch10/sleep"
	"time"
)

func main() {
	fmt.Println("Starting...")

	sleep.Sleep(time.Second * 5)

	fmt.Println("Ending after 5 seconds")
}
