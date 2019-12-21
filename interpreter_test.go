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
