package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15);

	if total != 30 {
		t.Errorf("Resultado da soma é inválido. Resultado: %d Esperado: %d", total, 30)
	}
	
}

func TestSub(t *testing.T) {
	total := multi(6, 4);

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