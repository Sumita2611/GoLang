package main
import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // signal that this goroutine is done
	fmt.Printf("worker %d started\n", i)
	//some task is running
	fmt.Printf("worker %d ended\n", i)
}

func main() {
	fmt.Println("wait sync")

	var wg sync.WaitGroup
	wg.Add(1) // injcrement the waitgroup counter

	// start 3 worker goroutines
	for i:=0; i<3; i++ {
		go worker(i, &wg);
	}
	wg.Wait() // wait for all goroutines to finish
	fmt.Println("workers task completed")
}