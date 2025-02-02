package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go printMessage("hello 1", &wg)
	wg.Add(1)
	go printMessage("hello 2", &wg)
	wg.Add(1)
	go printMessage("hello 3", &wg)
	wg.Add(1)
	go printMessage("hello 4", &wg)

	wg.Wait()
	fmt.Println("All goroutines are done")

}

func printMessage(s string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println(s)
	time.Sleep(1 * time.Second)
}
