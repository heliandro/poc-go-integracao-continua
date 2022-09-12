package main

import "fmt"

func main() {
	fmt.Println(soma(17, 25))

	fmt.Println(sub(15, 10))

	fmt.Println(multi(2, 3))
}

func soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func multi(a int, b int) int {
	return a * b
}

func divi(a int, b int) int {
	return a / b
}