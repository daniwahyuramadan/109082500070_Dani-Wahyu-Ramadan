package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	result1 := f(g(h(a)))
	result2 := g(h(f(b)))
	result3 := h(f(g(c)))

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}