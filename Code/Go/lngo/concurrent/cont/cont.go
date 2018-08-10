package main

import (
	"fmt"
	"sync"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	var x, y int
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(2)
	go func() {
		mu.Lock()
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
		wg.Done()
		mu.Unlock()
	}()
	go func() {
		mu.Lock()
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
		wg.Done()
		mu.Unlock()
	}()
	wg.Wait()
}
