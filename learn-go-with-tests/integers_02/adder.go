// Documentação
// https://go.dev/blog/godoc
// https://go.dev/doc/comment
// https://pkg.go.dev/go/doc
// https://tip.golang.org/doc/comment
// Documentação para Overview https://github.com/amalmadhu06/godoc-example#synopsis-and-overview
// "If you publish your code with examples to a public URL"
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers#testable-examples

package integers // Go Doc
import "fmt"

// Documentação para tipos e funções https://github.com/amalmadhu06/godoc-example#documentation-for-types-and-functions
// Adiciona recebe dois inteiros e retorna um inteiro que é a soma deles
func Add(x, y int) int {
	return x + y
}

// Documentação para Exemplo https://github.com/amalmadhu06/godoc-example#add-examples
// Exemplo de teste da função ADD(1,5) e imprime o retorno no terminal
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

/*
no terminal godoc -http=:6060
no navegador http://localhost:6060/pkg/learn-go-with-tests/integers_02/
o terminal fica bloqueado, no menu terminal selecione new terminal
*/
