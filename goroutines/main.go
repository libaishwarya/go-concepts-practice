package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}

	m2 := sync.RWMutex{}

	m := sync.Mutex{}
	a := []int{1, 2, 3, 4, 5}

	result := []int{}
	processed := make(chan int)

	wg2.Add(1)
	go func() {
		for v := range processed {
			m.Lock()
			result = append(result, v)
			m.Unlock()
		}
		wg2.Done()
	}()

	for _, v := range a {
		wg.Add(1)
		go slowProcess(v, &wg, processed)
	}

	wg.Wait()
	close(processed)
	wg2.Wait()

	fmt.Println(result)
}

func slowProcess(i int, wg *sync.WaitGroup, processed chan int) {
	time.Sleep(1 * time.Second)
	processed <- i * 2
	wg.Done()
}
