package main

import (
	"fmt"
	"sync"
)

func m5() {
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}


// Answer 5