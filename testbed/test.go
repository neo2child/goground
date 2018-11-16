package main

import "fmt"
import "reflect"

var x = 34

const y = 3.13 // Contstants

func multiplicationTable() {
	for i := 1; i < 10; i++ {
		for m := 1; m < 10; m++ {
			fmt.Println(i, "X", m, "=", i*m)
		}
	}
}

func main() {
	a := 10
	var b string = "20"
	fmt.Println(a, b)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(x)
	fmt.Println(y)
	multiplicationTable()
}
