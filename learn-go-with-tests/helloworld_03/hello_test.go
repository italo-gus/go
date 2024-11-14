package maingotdd // Definição do pacote

/* testing O teste de pacote fornece suporte para teste automatizado de pacotes Go. Ele é destinado a ser usado em conjunto com o comando "go test", que automatiza a execução de qualquer função do formulário.
https://pkg.go.dev/testing@go1.23.3

O nome do arquivo deve ser formado pelo arquivo a ser testado adicionado _test
*/

import (
	"testing"
)

/*
	Função para realizar o teste de ser formado por Test adicionado o nome da função a ser testada

necessitando do parâmetro t *testing.T
*/
func TestHello(t *testing.T) {
	got := Hello()         // variáve que recebe o retorno do resultdo da função a ser testada
	want := "Hello, world" // variáve que armazena o valor esperado de retorno para comparação do teste

	if got != want { // verifica se há diferença entre as variáveis o resultado da função e o esperado
		t.Errorf("got %q want %q", got, want) //
	}
}
