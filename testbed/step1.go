package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func take1() {
	fmt.Println("Ready Go!")
}

func take2() {
	fmt.Println(math.Pi)
}

func take3() {
	fmt.Printf("You have %g problems", math.Pi)
	fmt.Println()
}

func take4() {
	fmt.Println("Now time is", time.Now())
}

func take5() {
	fmt.Println("My favorite number is", rand.Intn(91))
}

func take6() {
	fmt.Printf("9의 제곱근은 %g 이다", math.Sqrt(9))
	fmt.Println()
}

func add(x, y int) int {
	return x + y
}

func seek() {
	x := 5
	var found int
	table := [][]int{
		{1, 5, 9},
		{2, 6, 5, 13},
		{5, 3, 7, 4},
	}
	for row := 0; row < len(table); row++ {
		for col := 0; col < len(table[row]); col++ {
			if table[row][col] == x {
				found += 1 //found = true
				fmt.Printf("%d를 찾았습니다. (row: %d, col: %d)\n", x, row, col)
			}
		}
	}
	if found > 0 {
		fmt.Printf("검색하신 '%d'를(을) '%d'개 찾았습니다.", x, found)
	} else {
		fmt.Printf("검색하신 '%d'을(를) 찾을 수 없습니다.", x)
	}
}

func main() {
	take1()
	take2()
	take3()
	take4()
	take5()
	take6()
	fmt.Println("함수 add의 값", add(2, 5))
	seek()
}
