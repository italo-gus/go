package main // Mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

// func Countdown() {}
// .\countdown_test.go:11:12: too many arguments in call to Countdown
// have (*bytes.Buffer)
// want ()

//func Countdown(out *bytes.Buffer) {}
// countdown_test.go:17: received "" expected "3"

/*
func Countdown(out *bytes.Buffer) {
	fmt.Fprint(out, "3")
}
*/

// Usando interface de uso Geral io.Writer
/*
func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
// countdown_test.go:22: received "3" expected "3\n\t2\n\t1\n\tGo!"
*/

/*
func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "Go!")
}
*/

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

/*
func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second) // adiciona espera de 1 segundo
	}
	fmt.Fprint(out, finalWord)
}
*/

/*
	func Countdown(out io.Writer, sleeper Sleeper) {
		for i := countdownStart; i > 0; i-- {
			fmt.Fprintln(out, i)
			time.Sleep(1 * time.Second)
		}

		fmt.Fprint(out, finalWord)
	}
*/
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

/*
func main() {
	Countdown(os.Stdout)
}
*/

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
