package main // Dependency Injection
import (
	"fmt"
	"io"
	"os"
)

//func main() {
//	Greet("Nome")
//}

func main() {
	Greet(os.Stdout, "Elodie")
}

/*
func Greet(name string) {
	// Funciona, mas não tem a Dependency Injection
	fmt.Printf("Hello, %s", name)
}
*/

/*
func Greet(writer *bytes.Buffer, name string) {
	// Funciona , não passa o teste, tem a Dependency Injection
	// A função Printf do pacote fmt. faz a chamada da função Fprintln passando o argumento os.Stdout
	// https://github.com/golang/go/blob/master/src/fmt/print.go
	// func Println(a ...any) (n int, err error) {  return Fprintln(os.Stdout, a...) }
	// não permitindo Dependency Injection.
	// func Fprintln(w io.Writer, a ...any) (n int, err error) { ... }
	// Onde Fprintln necessita de um parametro io.Writer que é uma interface
	// https://github.com/golang/go/blob/master/src/io/io.go
	// type Writer interface { Write(p []byte) (n int, err error) }
	// os.Stdout implementa a interface Writer
	fmt.Printf("Hello, %s \n", name)
	// é impresso em tela
}
*/

/*
func Greet(writer *bytes.Buffer, name string) {
	// Funciona , passa o teste, tem a Dependency Injection, não pode ser usado com os.Stdout
	fmt.Fprintf(writer, "Hello, %s", name)
	// não é impresso em tela ,  é armazenado em writer
}
*/

func Greet(writer io.Writer, name string) {
	// implementando com a interface
	// Funciona , passa o teste, tem a Dependency Injection, pode ser usado com os.Stdout
	fmt.Fprintf(writer, "Hello, %s", name)

}
