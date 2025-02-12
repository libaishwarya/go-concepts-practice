// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	start := time.Now() // Start timer

// 	arr := []int{1, 2, 3}

// 	for i, num := range arr {
// 		newNum := num * 2
// 		filename := fmt.Sprintf("file%d.txt", i+1)
// 		err := ioutil.WriteFile(filename, []byte(strconv.Itoa(newNum)), 0644)
// 		if err != nil {
// 			fmt.Println("Error", err)
// 			return

// 		}
// 		fmt.Println("✅ Stored", newNum, "in", filename)

// 	}
// 	elapsed := time.Since(start) // End timer
// 	fmt.Println("Total Execution Time:", elapsed)
// }

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"
	"time"
)

func main() {
	numbers := []int{1, 2, 3}
	start := time.Now() // Start timer

	result := make(chan string) //unbuffered channel for storing results
	var wg sync.WaitGroup       //waitgroups to synchronise goroutines

	//loop through each number and start a go routine
	for i, num := range numbers {
		wg.Add(1)
		go processNum(num, i, &wg, result)
	}

	go func() { // start a go routine and wait for al the task to complete and cose the channel
		wg.Wait()     //wait for al goroutine to complete
		close(result) //close result channel to signal completion
	}()

	for res := range result { //read and print result from channel
		fmt.Println(res)

	}

	elapsed := time.Since(start) // End timer
	fmt.Println("Total Execution Time:", elapsed)
}

func processNum(num, i int, wg *sync.WaitGroup, result chan<- string) {
	defer wg.Done()
	//Add 1 to number
	newNum := num + 1
	filename := fmt.Sprintf("file%d.txt", i+1) //create a file

	//writing to file
	err := ioutil.WriteFile(filename, []byte(strconv.Itoa(newNum)), 0644)

	if err != nil {
		result <- fmt.Sprintf("❌ Error writing to %s: %v", filename, err) //sends error message if not success
	} else {
		result <- fmt.Sprintf("✅ Stored %d in %s", newNum, filename) //else sends success message
	}

}
