package main

import "fmt"

/*
 Utilizando a solução do exercício anterior:
 Em package-level scope, atribua os seguintes valores às variáveis:
 para "x" atribua 42
 para "y" atribua "James Bond"
 para "z" atribua true
 Na função main:
 Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
 Demonstre a variável "s".
*/

var x int = 42
var y string = "James Bond" // string
var z bool = true           // bool

func main() {
	s := fmt.Sprintf("Valor de x: %d , Tipo de x: %T \nValor de y: %s , Tipo de y: %T \nValor de z: %t , Tipo de z: %T \n\n", x, x, y, y, z, z)
	fmt.Println("Valor de s:", s)
	fmt.Printf("Tipo de s: %T \n", s)

}
