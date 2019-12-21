package main

import (
	"bytes"
	"testing"
)

var runTests = []struct {
	descr string
	in    []byte
	out   []byte
}{
	{
		"output single byte",
		[]byte("+++."),
		[]byte{3},
	},

	{
		"skips unknown instructions",
		[]byte("++++ foobar ."),
		[]byte{4},
	},

	{
		"output multiple bytes by moving pointer forward",
		[]byte("+.>++.>+++."),
		[]byte{1, 2, 3},
	},

	{
		"moves pointer back and forth",
		[]byte("+>++>+++.<.<."),
		[]byte{3, 2, 1},
	},
}

func TestRun(t *testing.T) {
	for _, test := range runTests {
		t.Run(test.descr, func(t *testing.T) {

			got := Run(test.in)

			if bytes.Compare(got, test.out) != 0 {
				t.Errorf("got %x, want %x", got, test.out)
			}
		})
	}
}
