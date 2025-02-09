package main

import (
	"fmt"
	"sync"
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
	fmt.Println("this is the final value", val)
}
