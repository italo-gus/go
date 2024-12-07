package arrays // Arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	/*	var nomeDaArray = [capacidade] {valor1, valorn}
		Arrays têm uma capacidade fixa que você define quando declara a variável,
		podendo ser com a capacidade inferida
		nomeDaArray := [...] {valor1, valorn}
		https://go.dev/doc/effective_go#arrays
		https://go.dev/ref/spec#Array_types
	*/
	received := Sum(numbers)
	expected := 15

	if received != expected {
		t.Errorf("expected '%d' but received '%d' given, '%v'", expected, received, numbers)
	}
	/* %q imprimir strings, %d imprimir inteiro, %v imprimir o valor em um formato padrão.
	Formatação de erros https://pkg.go.dev/fmt#hdr-Printing	*/
}
