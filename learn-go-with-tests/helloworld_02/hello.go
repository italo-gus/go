package main // Definição do pacote

/* fmt O pacote fmt implementa E/S formatadas com funções análogas a printf e scanf do C. Os 'verbos' de formato são derivados do C, mas são mais simples.
https://pkg.go.dev/fmt@go1.23.3
*/

import "fmt" // Importar Biblioteca / Pacote

// Criação da Função Hello (sem parâmetros ) tipo de retorno string {}
func Hello() string {
	return " Hello, world"
}

// Criação de Função func nome_da_Função (parâmetros nome e tipo ) tipo de retorno {}
// Criação de Função main (sem parâmetros) sem retorno {}
func main() {
	fmt.Println(Hello()) // imprimir em tela (stdout [saída padrão do terminal]) o retorno da função Hello
}
