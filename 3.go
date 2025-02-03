package main

import (
	"fmt"
)

func  m3() {
	m := map[string]int {
		"a": 1, 
		"b": 2,
		"c": 3,
	}

	for _, v := range m {
		fmt.Println(v)
	}
}

// Answer Невозможно предсказать