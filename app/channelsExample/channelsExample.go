package main

import (
	"fmt"
	"sync"
	"time"
)

type Integer interface {
	int | int32 | int64
}

func generatorFunc[K any, T Integer](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func taken[K any, T Integer](done <-chan K, outCh <-chan T, n int) <-chan T {
	take := make(chan T)

	go func() {
		defer close(take)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case t, ok := <-outCh:
				if !ok {
					return
				}
				take <- t
			}
		}
	}()

	return take
}

func primeFinder[K any, T Integer](done <-chan K, st <-chan T) <-chan T {
	isPrime := func(el T) bool {
		for i := el - 1; i > 1; i-- {
			if el%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan T)

	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case val, ok := <-st:
				if !ok {
					return
				}
				if isPrime(val) {
					primes <- val
				}
			}
		}
	}()

	return primes
}

func fanIn[T any, K any](done <-chan K, channels ...<-chan T) <-chan T {
	wg := sync.WaitGroup{}
	fannedInStream := make(chan T)
	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}
	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}

func main() {
	done := make(chan int)

	now := time.Now()
	cpuCount := 3

	primeFindersChannels := make([]<-chan int, cpuCount)
	timeFunc := func() int { return 1000000007 }
	phase1 := generatorFunc(done, timeFunc)
	phase2 := taken(done, phase1, 2)

	//Fan-Out
	for i := 0; i < cpuCount; i++ {
		primeFindersChannels[i] = primeFinder(done, phase2)
	}

	//Fan-In
	for i := range fanIn(done, primeFindersChannels...) {
		fmt.Println(i)
	}

	fmt.Println("time taken:: ", time.Since(now))
}
