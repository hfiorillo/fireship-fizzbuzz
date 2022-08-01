package main

import (
	"fmt"
	"io"
	"os"
)

func FizzOrBuzz(writer io.Writer, number int) string {
	for i := 0; i < number; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Fprintf(writer, "FizzBuzz\n")
		} else if i%3 == 0 {
			fmt.Fprintf(writer, "Fizz\n")
		} else if i%5 == 0 {
			fmt.Fprintf(writer, "Buzz\n")
		} else {
			fmt.Fprintf(writer, "%d", i)
		}
	}
	return "Finished"
}

func IsItFizzOrBuzz(writer io.Writer, i int) string {
	if i%3 == 0 && i%5 == 0 {
		fmt.Fprintf(writer, "FizzBuzz")
	} else if i%3 == 0 {
		fmt.Fprintf(writer, "Fizz")
	} else if i%5 == 0 {
		fmt.Fprintf(writer, "Buzz")
	} else {
		fmt.Fprintf(writer, "%d", i)
	}
	return ""
}

func main() {
	FizzOrBuzz(os.Stdout, 100)
	IsItFizzOrBuzz(os.Stdout, 43)
}
