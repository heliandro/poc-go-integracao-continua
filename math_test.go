package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15);

	if total != 30 {
		t.Errorf("Resultado da soma é inválido. Resultado: %d Esperado: %d", total, 30)
	}
	
}

func TestSub(t *testing.T) {
	total := sub(10, 7);

	if total != 3 {
		t.Errorf("Resultado da sub é inválido. Resultado: %d Esperado: %d", total, 3)
	}
	
}

func TestMulti(t *testing.T) {
	total := multi(5, 5);

	if total != 25 {
		t.Errorf("Resultado da multi é inválido. Resultado: %d Esperado: %d", total, 25)
	}
	
}