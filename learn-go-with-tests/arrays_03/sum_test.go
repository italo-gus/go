package arrays // Slice / Cobertura de teste no Go

import "testing"

func TestSum(t *testing.T) {

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		/* 	Slice permite ter coleções de qualquer tamanho
		https://go.dev/doc/effective_go#slices
		https://go.dev/ref/spec#Slice_types	*/

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
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
