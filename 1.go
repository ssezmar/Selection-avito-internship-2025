package main

import "fmt"

type user struct {
	balance int64
}

func m1() {
	users := []user{
		{balance: 1000},
		{balance: 2000},
	}

	for _, u := range users {
		u.balance += 1000
		fmt.Println(u.balance)
		fmt.Println(users)
	}

	fmt.Println(users)
}

// Answer [{1000} {2000}]