package main

import (
	"bytes"
	"fmt"
	"testing"
)

var cases = []struct {
	Number int
	Value  string
}{
	{Number: 1, Value: "1"},
	{Number: 3, Value: "Fizz"},
	{Number: 5, Value: "Buzz"},
	{Number: 10, Value: "Buzz"},
	{Number: 29, Value: "29"},
	{Number: 30, Value: "FizzBuzz"},
	{Number: 60, Value: "FizzBuzz"},
	{Number: 81, Value: "Fizz"},
}

func TestIsItFizzOrBuzz(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d should give %s", test.Number, test.Value), func(t *testing.T) {
			var output bytes.Buffer
			IsItFizzOrBuzz(&output, test.Number)
			if output.String() != test.Value {
				t.Errorf("got %s, want %s", output.String(), test.Value)
			}
		})
	}
}
