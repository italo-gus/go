package arrays // Declarar função como variável

import (
	"reflect"
	"slices"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	t.Run("faz as somas de 1 slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		received := SumAll(numbers)
		expected := []int{6}

		if slices.Equal(received, expected) != true {
			t.Errorf("expected '%v' but the received was '%v' given, '%v'", expected, received, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, received, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(received, expected) {
			t.Errorf("received %v expected %v", received, expected)
		}
	}
	/*
		// nome_da_variável := func(parâmetros){bloco_da_função}
		utiliza a declaração curta de variável e a declaração de função
		https://go.dev/ref/spec#Short_variable_declarations
		https://go.dev/ref/spec#Function_declarations
	*/

	t.Run("faz as somas de alguns slices", func(t *testing.T) {
		received := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}
		checkSums(t, received, expected)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		received := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 12}
		checkSums(t, received, expected)
	})

}
