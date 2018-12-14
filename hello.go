package main

import (
	"fmt"
)

type rect struct {
	name string
	age  int
}

func display(s interface{}) {
	fmt.Println(s)
}

func main() {
	display(s)
}
