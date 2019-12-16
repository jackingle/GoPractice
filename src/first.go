package main

import (
	"fmt"
)

var thirtytwo int = 32

func main() {
	var i int = 42
	j := 42
	s := "This is a string"
	fmt.Println("Hello World! This is my first Go program")
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(s)
	fmt.Println(thirtytwo)
	fmt.Println("1 + 1 =", 1+1)
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
