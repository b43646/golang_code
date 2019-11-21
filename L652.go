package gotest

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad = 10
)

var wg sync.WaitGroup

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func worker(tasks chan string, worker int)  {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : Shuting Down\n", worker)
			return
		}
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}

func main() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Tasks : %d", post)
	}

	close(tasks)
	wg.Wait()
}
