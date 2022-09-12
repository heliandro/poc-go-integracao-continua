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

func hello(nome string, str1 string, str2 string) string {

	if (nome != "" && str1 != "" && str2 != "") {
		return nome + str1 + str2
	}

	if (nome != "" && str1 != "") {
		return nome + str1
	}

	if (nome != "" && str2 != "") {
		return nome + str2
	}

	if (str1 != "" && str2 != "") {
		return str1 + str2
	}

	if (str2 != "") {
		return str2
	}

	if (str1 != "") {
		return str1
	}

	if (nome != "") {
		return nome
	}

	return "hello!"
}
