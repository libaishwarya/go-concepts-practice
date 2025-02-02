package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	for j := 1; j <= 5; j++ {
		ch <- j
	}
	close(ch)
	wg.Wait()
}

func worker(id int, ch chan int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for job := range ch {
		fmt.Printf("Worker id is %d performing %d job\n", id, job)
	}

}
