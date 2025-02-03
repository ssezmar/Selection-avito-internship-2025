package main

import (
	"fmt"
	"sync"
)

func main() {
	f()
}


func f() {
	var m = make([]int, 500)

	wg := new(sync.WaitGroup)

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(i int) {
			m[i] = i
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Print("f done")
}

// Answer Может вывести f done, а может выдать панику