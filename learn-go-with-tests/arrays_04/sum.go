package arrays //

// Sum retona o cálculo da somas dos elementos de um slice de números inteiros.
func Sum(numbers []int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*
Funções variádicas (http://en.wikipedia.org/wiki/Variadic_function)
https://go.dev/ref/spec#Function_types
https://go.dev/ref/spec#Passing_arguments_to_..._parameters
*/
/*
/* Para ococrrer o erro "invalid operation: result != expected (slice can only be compared to nil)"
func SumAll(numbersToSum ...[]int) []int {
	return nil
}
*/

func SumAll(numbersToSum ...[]int) []int {
	return nil
}