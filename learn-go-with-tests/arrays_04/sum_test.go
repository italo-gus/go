package arrays //

import (
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
		result := Sum(numbers)
		expected := 6

		if result != expected {
			t.Errorf("expected '%d' but the result was '%d' given, '%v'", expected, result, numbers)
		}
	})
}

/* Erro
invalid operation: result != expected (slice can only be compared to nil)

func TestSumAll(t *testing.T) {
	t.Run("faz as somas de 1 slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := SumAll(numbers)
		expected := []int{3}

		if result != expected {
			t.Errorf("expected '%v' but the result was '%v' given, '%v'", expected, result, numbers)
		}
	})
}
*/

func TestSumAll(t *testing.T) {
	t.Run("faz as somas de 1 slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := SumAll(numbers)
		expected := []int{3}

		if slices.Equal(result, expected) != true {
			t.Errorf("expected '%v' but the result was '%v' given, '%v'", expected, result, numbers)
		}
	})
}

/*
	t.Run("faz as somas de alguns slices", func(t *testing.T) {
		result := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		//if !reflect.DeepEqual(result, expected) {

			t.Errorf("expected '%v' but result '%v'", expected, result)
		}
	})

}
*/
/*
Comparação entre tipos sem erro e compararação profunda
https://pkg.go.dev/builtin#comparable
https://pkg.go.dev/reflect#DeepEqual

Comparação entre slice
Equal relata se duas fatias são iguais: o mesmo comprimento e todos os elementos iguais.
https://pkg.go.dev/slices#Equal


*/
