package main

import (
	"fmt"
)

/*
Go fmt tem três funções para gerar texto:

Print() Imprime formatos usando os formatos padrão para seus operandos e grava na saída padrão.
Espaços são adicionados entre operandos quando nenhum deles é uma string.
Ele retorna o número de bytes gravados e qualquer erro de gravação encontrado.

Println() formata usando os formatos padrão para seus operandos e grava na saída padrão.
Espaços são sempre adicionados entre operandos e uma nova linha é anexada.
A função é semelhante, Print() com a diferença de que é adicionado um espaço em branco
e uma nova linha é adicionada no final.
Ele retorna o número de bytes gravados e qualquer erro de gravação encontrado.

Printf() primeiro formata seu argumento com base no verbo de formatação fornecido e depois grava na saída padrão.
Ele retorna o número de bytes gravados e qualquer erro de gravação encontrado.
Lista de verbos de formatação https://pkg.go.dev/fmt#hdr-Printing

https://pkg.go.dev/fmt#pkg-functions
*/
func main() {
	var a string = "Hello"
	var b string = "World"
	var c int = 10
	var d float64 = 3.141
	var e float64 = 52.11141546
	var f = false

	fmt.Println("fmt.Print()")
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print("\n")       // imprime quebra de linha
	fmt.Print(a, b, "\n") // já com imprime quebra de linha
	/*
	   HelloWorld
	   HelloWorld
	*/

	fmt.Println("\nfmt.Println()")
	fmt.Println(a, b)
	x, err := fmt.Println(a, b)

	fmt.Println(x) // o número de bytes gravados
	fmt.Println(err)
	/*
		Hello World
		Hello World
		12
		<nil>
	*/

	fmt.Println("\nfmt.Printf()")
	fmt.Printf("a tem valor: %v e tipo: %T \n", a, a)
	fmt.Printf("c tem valor: %v e tipo: %T \n", c, c)
	fmt.Printf("%v%%\n", c) // adicionando %
	/*
		fmt.Printf()
		a tem valor: Hello e tipo: string
		c tem valor: 10 e tipo: int
		10%
	*/

	fmt.Println("\nPrintf com string ")
	fmt.Printf("%s \n", a)    // %s Imprime o valor como uma string simples
	fmt.Printf("%q \n", a)    // %q Imprime o valor como uma string entre aspas duplas
	fmt.Printf("%20s \n", a)  // %20s Imprime o valor como uma string simples (largura 20, justificado à direita)
	fmt.Printf("%-20s \n", a) // %-20s Imprime o valor como uma string simples (largura 20, justificado à esquerda)
	fmt.Printf("%x \n", a)    // %x Imprime o valor como um dump hexadecimal de valores de bytes
	fmt.Printf("% x \n", a)   // % x Imprime o valor como um dump hexadecimal com espaços
	/*
		Printf com string
		Hello
		"Hello"
		               Hello
		Hello
		48656c6c6f
		48 65 6c 6c 6f
	*/

	fmt.Println("\nPrintf com inteiro")
	fmt.Printf("%b \n", c)    // %b Base 2
	fmt.Printf("%d \n", c)    // %d Base 10
	fmt.Printf("%+d \n", c)   // %+d Base 10 e sempre mostrar sinal
	fmt.Printf("%o \n", c)    // %o Base 8
	fmt.Printf("%O \n", c)    // %O Base 8, com 0o inicial
	fmt.Printf("%x \n", c)    // %x Base 16, minúsculas
	fmt.Printf("%X \n", c)    // %X Base 16, maiúsculas
	fmt.Printf("%#x \n", c)   // %#x Base 16, com 0x inicial
	fmt.Printf("%20d \n", c)  // %20d Preenchimento com espaços  (largura 20, justificado à direita)
	fmt.Printf("%-20d \n", c) // %-20d Preenchimento com espaços (largura 20, justificado à esquerda)
	fmt.Printf("%020d \n", c) // %020d Preenchimento com zeros   (largura 20, justificado à direita)
	fmt.Printf("%T\n", c)
	/*
		Printf com inteiro
		1010
		10
		+10
		12
		0o12
		a
		A
		0xa
						10
		10
		00000000000000000010
		int
	*/

	fmt.Println("\nPrintf com ponto flutuante ")
	fmt.Printf("%e\n", d)     // %e Notação científica com 'e' como expoente
	fmt.Printf("%f\n", d)     // %f Ponto decimal, sem expoente
	fmt.Printf("%.2f\n", d)   // %.2f Largura padrão, precisão 2
	fmt.Printf("%6.2f\n", d)  // %6.2f Largura 6, precisão 2
	fmt.Printf("%20.5f\n", d) // %20.5f Largura 20, precisão 5
	fmt.Printf("%20.5f\n", e) // %20.5f Largura 6, precisão 5
	fmt.Printf("%g\n", e)     // %g Expoente conforme necessário, apenas dígitos necessários
	fmt.Printf("%T\n", e)
	/*
		Printf com ponto flutuante
		3.141000e+00
		3.141000
		3.14
		3.14
					3.14100
					52.11142
		52.11141546
		float64
	*/

	fmt.Println("\nPrintf com booleano ")
	fmt.Printf("%t\n", f)
	/*
		Printf com booleano
		false
	*/
}
