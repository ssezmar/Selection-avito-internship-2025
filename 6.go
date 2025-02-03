package main

import "fmt"


func m6() {
	ch := make(chan int)
	close(ch)

	val := <-ch
	fmt.Println(val)
}

// Answer 0