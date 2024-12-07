package arrays // Slice / Cobertura de teste no Go

import "testing"

func TestSum(t *testing.T) {
	/*
		t.Run("coleção de 5 números", func(t *testing.T) {
			numbers := []int{1, 2, 3, 4, 5}
			received := Sum(numbers)
			expected := 15

			if received != expected {
				t.Errorf("expected '%d' but received '%d' given, '%v'", expected, received, numbers)
			}
		})
	*/
	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		/* 	Slice permite ter coleções de qualquer tamanho
		https://go.dev/doc/effective_go#slices
		https://go.dev/ref/spec#Slice_types
		https://pkg.go.dev/slices#pkg-overview
		https://go.dev/blog/slices-intro */

		received := Sum(numbers)
		expected := 6

		if received != expected {
			t.Errorf("expected '%d' but received '%d' given, '%v'", expected, received, numbers)
		}
	})
}

/*	Cobertura de teste no Go
	no terminal go test -cover
	https://go.dev/blog/cover#test-coverage
PASS
coverage: 100.0% of statements
ok      learn-go-with-tests/arrays_03   9.297s
*/
