package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(abc())
}

func abc() (s string, err error) {
	m, err := 1, errors.New("fuck")
	if true {
		fmt.Println(m)
		// err = errors.New("you")
	}
	return
}
