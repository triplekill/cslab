package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	wg      sync.WaitGroup
	counter int64
	mutex   sync.Mutex
)

const (
	numberGoroutines = 4  // Number of goroutines to use.
	taskLoad         = 10 // Amount of work to process.
)

func test24() {
	tasks := make(chan string, taskLoad)

	// Launch goroutines to handle the work.
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// Add a bunch of work to get done.
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// Close the channel so the goroutines will quit
	// when all the work is done.
	close(tasks)

	// Wait for all the work to get done.
	wg.Wait()
}
func worker(tasks chan string, worker int) {
	// Report that we just returned.
	defer wg.Done()

	for {
		// Wait for work to be assigned.
		task, ok := <-tasks
		if !ok {
			// This means the channel is empty and closed.
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// Display we are starting the work.
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// Randomly wait to simulate work time.
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// Display we finished the work.
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}

func test16() {
	wg.Add(2)

	// Create two goroutines.
	go incCounter3(1)
	go incCounter3(2)

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}
func incCounter3(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Only allow one goroutine through this
		// critical section at a time.
		mutex.Lock()
		{
			// Capture the value of counter.
			value := counter

			// Yield the thread and be placed back in queue.
			runtime.Gosched()

			// Increment our local value of counter.
			value++

			// Store the value back into counter.
			counter = value
		}
		mutex.Unlock()
		// Release the lock and allow any
		// waiting goroutine through.
	}
}

// main is the entry point for all Go programs.
func test9() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	go incCounter(1)
	go incCounter(2)

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable.
func incCounter(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Capture the value of Counter.
		value := counter

		// Yield the thread and be placed back in queue.
		runtime.Gosched()

		// Increment our local value of Counter.
		value++

		// Store the value back into Counter.
		counter = value
	}
}
func test13() {
	wg.Add(2)

	// Create two goroutines.
	go incCounter2(1)
	go incCounter2(2)

	// Wait for the goroutines to finish.
	wg.Wait()

	// Display the final value.
	fmt.Println("Final Counter:", counter)
}
func incCounter2(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Safely Add One To Counter.
		atomic.AddInt64(&counter, 1)

		// Yield the thread and be placed back in queue.
		runtime.Gosched()
	}
}

// main is the entry point for all Go programs.
func test1() {
	// Allocate 1 logical processor for the scheduler to use.
	// runtime.GOMAXPROCS(2)
	// 使用 1 和 2 的结果不同

	// wg is used to wait for the program to finish.
	// Add a count of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()

		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()

		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

func main() {
	// test1()
	// test9()
	// test13()
	// test16()
	test24()
}
