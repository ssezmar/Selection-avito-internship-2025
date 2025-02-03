package main

import (
	"fmt"
)


func m4() {
	b := new(bool)
	modify(b)
	fmt.Println(b)
}

func modify(b *bool) {
	b = nil
}

// Answer Какой-то нулевой указатель