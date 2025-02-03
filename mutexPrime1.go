package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu      sync.Mutex
	counter int
}

func (counter *Counter) Inc() {
	counter.mu.Lock()
	counter.counter++
	counter.mu.Unlock()
}

func main() {
	var c1, c2 Counter
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				c1.Inc()
				c2.Inc()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c1.counter, c2.counter)
}
