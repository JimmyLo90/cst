package main

import "fmt"

func main() {
	var a map[int]int
	fmt.Println(a, &a, a == nil)

	a = make(map[int]int)

	fmt.Println(a, &a, a == nil)
}
