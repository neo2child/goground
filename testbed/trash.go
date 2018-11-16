package main

import "fmt"

func outerFunction() {
	outer := "I'm outer!"
	innerFunction := func() {
		inner := "I'm inner!"
		reinnerFunction := func() {
			reinner := "I'm reinner"
			fmt.Println(reinner)
		}
		reinnerFunction()
		fmt.Println(inner)
	}
	innerFunction()
	fmt.Println(outer)
}

func main() {
	outerFunction()

}
