package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15);

	if total != 30 {
		t.Errorf("Resultado da soma é inválido. Resultado: %d Esperado: %d", total, 30)
	}
	
}

func TestSub(t *testing.T) {
	total := sub(6, 4);

	if total != 2 {
		t.Errorf("Resultado da sub é inválido. Resultado: %d Esperado: %d", total, 2)
	}
	
}

func TestMulti(t *testing.T) {
	total := multi(5, 5);

	if total != 25 {
		t.Errorf("Resultado da multi é inválido. Resultado: %d Esperado: %d", total, 25)
	}
	
}

func TestDivi(t *testing.T) {
	total := divi(2, 2);

	if total != 1 {
		t.Errorf("Resultado da div é inválido. Resultado: %d Esperado: %d", total, 1)
	}
	
}

func TestHelloWithNoName(t *testing.T) {
	result := hello("", "", "")
	expected :=  "hello!"

	if result != expected {
		t.Errorf("Resultado do teste hello é inválido. Resultado: %s Esperado: %s", result, expected)
	}
}


func TestHelloWithArg2(t *testing.T) {
	nome := ""
	result := hello(nome, "Ola: ", "")
	expected :=  "Ola: "

	if result != expected {
		t.Errorf("Resultado do teste hello é inválido. Resultado: %s Esperado: %s", result, expected)
	}
}

func TestHelloWithArg3(t *testing.T) {
	nome := ""
	result := hello(nome, "", "| WORLD")
	expected :=  "| WORLD"

	if result != expected {
		t.Errorf("Resultado do teste hello é inválido. Resultado: %s Esperado: %s", result, expected)
	}
}

func TestHelloWithArg4(t *testing.T) {
	nome := "teste"
	result := hello(nome, "", "")
	expected :=  "teste"

	if result != expected {
		t.Errorf("Resultado do teste hello é inválido. Resultado: %s Esperado: %s", result, expected)
	}
}

func TestHelloWithArg5(t *testing.T) {
	nome := "teste"
	result := hello(nome, "222", "")
	expected :=  "teste222"

	if result != expected {
		t.Errorf("Resultado do teste hello é inválido. Resultado: %s Esperado: %s", result, expected)
	}
}

func TestHelloWithArg6(t *testing.T) {
	nome := "teste"
	result := hello(nome, "", "333")
	expected :=  "teste333"

	if result != expected {
		t.Errorf("Resultado do teste hello é inválido. Resultado: %s Esperado: %s", result, expected)
	}
}

func TestHelloWithArg7(t *testing.T) {
	nome := "teste"
	result := hello(nome, "222", "333")
	expected :=  "teste222333"

	if result != expected {
		t.Errorf("Resultado do teste hello é inválido. Resultado: %s Esperado: %s", result, expected)
	}
}

func TestHelloWithArg8(t *testing.T) {
	nome := ""
	result := hello(nome, "222", "333")
	expected :=  "222333"

	if result != expected {
		t.Errorf("Resultado do teste hello é inválido. Resultado: %s Esperado: %s", result, expected)
	}
}

// func TestHelloWithNoArgs(t *testing.T) {
// 	hello.nome := 
// 	result := hello("a", "b", "c")
// 	expected :=  "abc"

// 	if result != expected {
// 		t.Errorf("Resultado do teste hello é inválido. Resultado: %s Esperado: %s", result, expected)
// 	}
// }