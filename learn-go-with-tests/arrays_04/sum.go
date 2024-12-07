package arrays // Funções variádicas,  reflect.DeepEqual , trabalhando com slice

// Sum retona o cálculo da somas dos elementos de um slice de números inteiros.
func Sum(numbers []int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*
/* Para ococrrer o erro "invalid operation: result != expected (slice can only be compared to nil)"
func SumAll(numbersToSum ...[]int) []int {
	return nil
}
*/

/*
Funções variádicas (http://en.wikipedia.org/wiki/Variadic_function)
https://go.dev/ref/spec#Function_types
https://go.dev/ref/spec#Passing_arguments_to_..._parameters
*/

/*
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum) // len(slice) comprimento do slice https://go.dev/ref/spec#Length_and_capacity
	sums := make([]int, lengthOfNumbers) // make([]Tipo dos elementos, comprimento, capacidade) cria um slice https://go.dev/ref/spec#Slice_types
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers) // utiliza os elementos do slice como o array com índices
	}
	return sums
}
*/

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers)) // append(parâmetros, n ) retorna slice adicionado os n elementos do parâmetros
	}
	return sums
}

/*
slice append https://go.dev/ref/spec#Appending_and_copying_slices
*/

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			//tail := numbers[1:] // parte do slice
			tail := numbers
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

/*
slice copiando parte
https://go.dev/ref/spec#Slice_expressions
Lembrando que como slice é interessante criar (make) um para cópia,
ou será associado o primeiro slice (ponteiro).
Exemplo em
https://go.dev/play/p/ICCWcRGIO68
https://go.dev/play/p/bTrRmYfNYCp
*/
