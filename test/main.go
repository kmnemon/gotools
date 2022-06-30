package main

import "fmt"

func fn() int {
	var x int
	if x < 10 {
		x = x + 1
		fmt.Println(x)
		fn()
	} else {
		return x
	}
	return 90
}

func main() {
	fmt.Print(fn())
}
