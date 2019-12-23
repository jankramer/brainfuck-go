package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

var runTests = []struct {
	descr   string
	program []byte
	input   io.Reader
	out     []byte
}{
	{
		"output single byte",
		[]byte("+++."),
		new(bytes.Buffer),
		[]byte{3},
	},

	{
		"output input",
		[]byte(",."),
		bytes.NewBuffer([]byte{10}),
		[]byte{10},
	},

	{
		"handles input error",
		[]byte(","),
		new(bytes.Buffer),
		nil,
	},

	{
		"adds and subtracts",
		[]byte("+-++-+++-."),
		new(bytes.Buffer),
		[]byte{3},
	},

	{
		"skips unknown instructions",
		[]byte("++++ foobar ."),
		new(bytes.Buffer),
		[]byte{4},
	},

	{
		"output multiple bytes by moving pointer forward",
		[]byte("+.>++.>+++."),
		new(bytes.Buffer),
		[]byte{1, 2, 3},
	},

	{
		"moves pointer back and forth",
		[]byte("+>++>+++.<.<."),
		new(bytes.Buffer),
		[]byte{3, 2, 1},
	},

	{
		"handles simple loops",
		[]byte("+++++[>>+<<-]>>."),
		new(bytes.Buffer),
		[]byte{5},
	},
	{
		"hello world",
		[]byte("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."),
		new(bytes.Buffer),
		[]byte("Hello World!\n"),
	},
}

func TestRun(t *testing.T) {
	for _, test := range runTests {
		t.Run(test.descr, func(t *testing.T) {

			output := new(bytes.Buffer)
			err := Run(test.program, test.input, output)

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got := output.Bytes()
			if bytes.Compare(got, test.out) != 0 {
				t.Errorf("got %x, want %x", got, test.out)
			}
		})
	}
}

var scanLoopTests = []struct {
	descr string
	in    []byte
	out   map[int]int
	err   bool
}{
	{
		"simple loop",
		[]byte("[]"),
		map[int]int{0: 1},
		false,
	},

	{
		"1 nested loop",
		[]byte("[[]]"),
		map[int]int{0: 3, 1: 2},
		false,
	},

	{
		"multiple nested loops",
		[]byte("[[234[6]8[0]2]4]"),
		map[int]int{0: 15, 1: 13, 5: 7, 9: 11},
		false,
	},

	{
		"unbalanced loop",
		[]byte("[[]"),
		nil,
		true,
	},
}

func TestScanLoops(t *testing.T) {
	for _, test := range scanLoopTests {
		t.Run(test.descr, func(t *testing.T) {

			loops, err := scanLoops(test.in)

			if !reflect.DeepEqual(loops, test.out) {
				t.Errorf("got %v, want %v", loops, test.out)
			}

			if test.err && err == nil {
				t.Errorf("got nil, want error")
			}
		})
	}
}
