package main

import "fmt"

/*
 Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
 Identificador "x" deverá ter tipo int
 Identificador "y" deverá ter tipo string
 Identificador "z" deverá ter tipo bool
 Na função main:
 Demonstre os valores de cada identificador
 O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
 valor 0 / padrão / default
 https://go.dev/ref/spec#The_zero_value
*/

var x int32
var y string
var z bool

func main() {
	fmt.Printf("Valor de x: %d\nValor de y: %s\nValor de z: %t\n", x, y, z)
}
