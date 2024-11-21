package main

import (
	"fmt"
)

/* Declaração de variável
https://go.dev/ref/spec#Variables
https://go.dev/ref/spec#Variable_declarations

Sintaxe
var nome tipo = valor (especificando tipo e valor)

var a string (definindo do tipo string valor padrão "")
var b int    (definindo tipo int valor padrão 0)
var c bool   (definindo tipo bool valor padrão false)

a = "a"      (atribuindo valor após declaração)

** Criação de variável sem declaração

nestes casos o tipo da variável é inferido pelo compilador,	o compilador decide o tipo da variável com base no valor.
Só pode ser usado dentro de funções
erro de utilizar fora de uma função:
	"syntax error: non-declaration statement outside function body"
	erro de sintaxe: declaração não declarada fora do corpo da função

d := 1               (sem var sem definição do tipo da variáve e definindo o valor com " := ")
var e = "inferido"   (com var sem definição do tipo da variáve e definindo o valor com " = ")


*** Declaração de ​múltiplas variáveis
definindo tipo e seus com  respectivos valores
var  f, g, h int = 1, 2, 3

sem definindo tipo (criação) e seus com  respectivos valores
	com var
	var i, j = 6, "Hello"

  	sem var, usando :=
	k, l := 7, "World!"


**** Declaração de variável em um bloco
var (
	m int
	n int = 1
    o string = "hello"
)

***** Regras de nomenclatura de variáveis ​
Uma variável pode ter um nome curto ( x e y) ou um nome mais descritivo (nome, idade, preço, nome_do_carro, e outros).

Regras de nomenclatura de variáveis​:

Um nome de variável deve começar com uma letra ou um caractere sublinhado (_)
Um nome de variável pode conter apenas caracteres alfanuméricos e sublinhados ( a-z, A-Z, 0-9, e _)
Os nomes das variáveis ​​diferenciam maiúsculas de minúsculas case-sensitive (idade, Idade e IDADE são três variáveis ​​diferentes)
Não há limite para o comprimento do nome da variável
Um nome de variável não pode conter espaços
Um nome de variável não pode começar com um dígito .
O nome da variável não pode ser nenhuma palavra-chave Go

Constantes
É variáveis que tem um valor fixo que não pode ser alterado, é imutável e somente leitura. usar palavra-chave a const
const nome tipo = valor (especificando tipo e valor)

Boas práticas
Camel case
	Deve começar com a primeira letra minúscula e  subsequente
	valorInicial
	valorFinal
Pascal case
	A primeira letra maiúscula de cada palavramaiúscula:
	ValorInicial
	ValorFinal
Snake case
	Utiliza underline no lugar do espaço para separar as palavras
	valor_inicial
	valor_final

Padrão visibilidade do GO
	Os nomes na linguagem Go tem um efeito semântico,
	a visibilidade de um identificador fora de um pacote é determinada
	se o nome inicia com letra maiúscula ou minúscula,
	onde: maiúsculas é exportado; minúsculas não é exportado.
	Go não fornece suporte automático para getters e setters.
	Não há nada de errado em fornecer getters e setters você mesmo,
	e geralmente é apropriado fazer isso.
	Mas não é idiomático nem necessário colocar Get no nome do getter.
	Se você tiver um campo chamado owner (minúsculas, não exportado),
	o método getter deve ser chamado Owner (maiúsculas, exportado), não GetOwner.
	O uso de nomes em maiúsculas para exportação fornece o gancho para discriminar o campo do método.
	Uma função setter, se necessária, provavelmente será chamada SetOwner.

	sendo assim:
	PascalCase para exportar (é exportado para fora do pacote)
	camelCase para internos (não é exportado  para fora do pacote).
*/

var a string = "teste" // declara variável com nome "a" do tipo string com valor inicial teste
var b int              // declara variável com nome "b" do tipo int
var c bool             // declara variável com nome "c" do tipo bool

var f, g, h int = 3, 4, 5

var (
	m int
	n int    = 1
	o string = "hello"
)

const pi = 3.14 //  declara constante PI π

func main() {
	b = 1 // atribui o valor 1 a variável "a"
	c := true

	d := 2
	var e = "inferido"

	var i, j = 6, "Hello"
	k, l := 7, "World!"

	fmt.Print("a ")
	fmt.Println(a)
	fmt.Println("b ", b)
	fmt.Println("c ", c)
	fmt.Println("d ", d)
	fmt.Println("e ", e)
	fmt.Println("f ", f)
	fmt.Println("g ", g)
	fmt.Println("h ", h)
	fmt.Println("i ", i)
	fmt.Println("j ", j)
	fmt.Println("k ", k)
	fmt.Println("l ", l)
	fmt.Println("m ", m)
	fmt.Println("n ", n)
	fmt.Println("o ", o)
	fmt.Println("π ", pi)
	// pi = 1111
	/* erro gerado por tentar alterar uma contante
	cannot assign to pi (neither addressable nor a map index expression)*/
}
