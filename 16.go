package main

import (
	"fmt"
)

func m16() {
	inc := func() {
		//v++
		fmt.Print(1, " ") // side
	}

	v := 1

	fmt.Print(v, " ")

	inc()
	fmt.Print(v, " ")
}


// Answer Ошибка компиляции