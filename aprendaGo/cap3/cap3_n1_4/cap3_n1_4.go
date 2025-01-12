package main

import "fmt"

/*
 Crie um tipo. O tipo subjacente deve ser int.
 Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
 Na função main:
 Demonstre o valor da variável "x"
 Demonstre o tipo da variável "x"
 Atribua 42 à variável "x" utilizando o operador "="
 Demonstre o valor da variável "x"
 Para os aventureiros: https://golang.org/ref/spec#Types
*/

type novotipo int

var x novotipo

func main() {
	fmt.Printf("Valor de x: %d , Tipo de x: %T \n", x, x)
	x = 42
	fmt.Println("Valor de x: ", x)
}
