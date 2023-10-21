package main

import (
	"fmt"
	"io"
	"os"
)

const countdownStart = 3
const finalWord = "Go!"

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}