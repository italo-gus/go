package main

import "fmt"

// https://www.youtube.com/c/AprendaGo
// https://github.com/vkorbes/aprendago/blob/master/OUTLINE.md

/*
 Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
 42
 "James Bond"
 true
 Agora demonstre os valores nestas variáveis, com:

 Uma única declaração print
 Múltiplas declarações print
*/
func main() {
	// := gopher / declarador curto  https://go.dev/ref/spec#Short_variable_declarations
	// nestes casos o tipo da variável é inferido pelo compilador,	o compilador decide o tipo da variável com base no valor.
	// Só pode ser usado dentro de bloco de códigos
	// erro de utilizar fora de um bloco de código / função:
	// "syntax error: non-declaration statement outside function body"
	// erro de sintaxe: declaração não declarada fora do corpo da função

	x := 42           // int
	y := "James Bond" // string
	z := true         // bool

	fmt.Printf("Valor de x: %d , Tipo de x: %T \nValor de y: %s , Tipo de y: %T \nValor de z: %t , Tipo de z: %T \n\n", x, x, y, y, z, z)
	// https://pkg.go.dev/fmt

	fmt.Println("Valor de x: ", x)
	fmt.Println("Valor de y: ", y)
	fmt.Println("Valor de z: ", z)

}
