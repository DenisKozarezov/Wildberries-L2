package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	wg := &sync.WaitGroup{}

	output := make(chan interface{})
	for _, channel := range channels {
		wg.Add(1)
		go func(channel <-chan interface{}) {
			defer wg.Done()
			for value := range channel {
				output <- value
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
		sig(5*time.Second),
	)

	fmt.Printf("done after %v", time.Since(start))
}
