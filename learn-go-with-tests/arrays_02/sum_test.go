package arrays // Refactor - Refatoração / for range / _ blank identifier (identificador em branco)

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	received := Sum(numbers)
	expected := 15

	if received != expected {
		t.Errorf("expected '%d' but received '%d' given, '%v'", expected, received, numbers)
	}

}
