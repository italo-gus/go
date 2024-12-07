package arrays // Funções variádicas,  reflect.DeepEqual , trabalhando com slice

import (
	"reflect"
	"slices"
	"testing"
)

/*
reflect O pacote reflect implementa reflexão em tempo de execução,
permitindo que um programa manipule objetos com tipos arbitrários.
O uso típico é pegar um valor com interface de tipo estático{} e extrair
suas informações de tipo dinâmico chamando TypeOf, que retorna um Type.
https://pkg.go.dev/reflect
*/

func TestSum(t *testing.T) {
	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		received := Sum(numbers)
		expected := 6

		if received != expected {
			t.Errorf("expected '%d' but received '%d' given, '%v'", expected, received, numbers)
		}
	})
}

/* Para ococrrer o erro "invalid operation: received != expected (slice can only be compared to nil)"

func TestSumAll(t *testing.T) {
	t.Run("faz as somas de 1 slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		received := SumAll(numbers)
		expected := []int{3}

		if received != expected {
			t.Errorf("expected '%v' but the received was '%v' given, '%v'", expected, received, numbers)
		}
	})
}
*/

/*
Para ococrrer o erro "in call to slices.Equal, type string of expected does not match inferred type []int for S"

	func TestSumAll(t *testing.T) {
		t.Run("faz as somas de 1 slices", func(t *testing.T) {
			numbers := []int{1, 2, 3}
			received := SumAll(numbers)
			expected := "bob"

			if slices.Equal(received, expected) != true {
				t.Errorf("expected '%v' but the received was '%v' given, '%v'", expected, received, numbers)
			}
		})
	}
*/
func TestSumAll(t *testing.T) {
	t.Run("faz as somas de 1 slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		received := SumAll(numbers)
		expected := []int{6}

		if slices.Equal(received, expected) != true {
			t.Errorf("expected '%v' but the received was '%v' given, '%v'", expected, received, numbers)
		}
	})

	t.Run("faz as somas de alguns slices", func(t *testing.T) {
		received := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !reflect.DeepEqual(received, expected) {
			t.Errorf("received %v expected %v", received, expected)
		}
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		received := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 12}

		if !reflect.DeepEqual(received, expected) {
			t.Errorf("received %v expected %v", received, expected)
		}
	})

}

/*
Comparação entre tipos sem erro e compararação profunda
https://pkg.go.dev/builtin#comparable
https://pkg.go.dev/reflect#DeepEqual

Comparação entre slice
Equal relata se duas fatias são iguais: o mesmo comprimento e todos os elementos iguais.
https://pkg.go.dev/slices#Equal


*/
