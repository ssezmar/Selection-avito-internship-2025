package main

import (
	"fmt"
)

func m15() {
	v := 1
	fmt.Print(v, " ")

	f1(&v)
	fmt.Print(v, " ")
}

func f1(v *int) {
	*v = *v * 3
}

// Answer 1 3