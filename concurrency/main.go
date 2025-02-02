package main

import (
	"fmt"
	"time"
)

//Function is defined separatey and called as goroutines which can aso be reused in code if needed

func main() {
	go sayHello()
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Main function done")
}

func sayHello() {
	fmt.Println("Hello")
}

// Anonymous function is defined and executed immediately
// func main() {
// 	go func() {
// 		fmt.Println("Hello")
// 	}()
// 	time.Sleep(1 * time.Millisecond)
// 	fmt.Println("Main function done")
// }
