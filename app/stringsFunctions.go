package main

import (
	"fmt"
	"sync"
)

const (
	Monday = iota;
	Tuesday
	WednesDay
	ThrusDay
	Friday
	Saturday
	Sunday
)

const (
	a = 1
	b
	c
)
var mu sync.Mutex
var val int = 0
var wg sync.WaitGroup
func main() {
	increment:= func (){
		mu.Lock()
		defer mu.Unlock()
		val++
		defer wg.Done()
	}	

	decrement:= func(){
		mu.Lock()
		defer mu.Unlock()
		val--;
		defer wg.Done()
	}
	for i:=0 ; i < 10000000 ; i++ {
		wg.Add(2)
		go increment()
		go decrement()
	}

	wg.Wait()
	fmt.Println("shi")
	fmt.Println("this is the final value", val)
	fmt.Println(Monday, Tuesday, WednesDay, Friday, ThrusDay, Saturday, Sunday)
	fmt.Println(a,b,c)
}
