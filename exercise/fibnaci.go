package main

import "fmt"

func main() {
	var i int
	var t = 1

	for i = 0; i < 100; {
		t = i + t
		fmt.Println(t)
		i = i + t
		fmt.Println(i)
	}

	fmt.Println(fibonacci(7))

}

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}

}
