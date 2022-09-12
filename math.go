package main

import "fmt"

func main() {
	fmt.Println(soma(17, 25))

	fmt.Println(sub(15, 10))

	fmt.Println(multi(2, 3))

	fmt.Println(divi(2, 2))
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

func hello(nome string) string {
	var str1 string
	str1 = "Nome: "

	var str2 string
	str2 = " | hello!"

	result := str1 + nome + str2

	return result
}