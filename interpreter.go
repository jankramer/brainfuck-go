package main

import "bytes"

// Run executes a brainfuck program and returns its output.
func Run(program []byte) []byte {
	buf := new(bytes.Buffer)
	tape := make(map[int]byte)

	end := len(program)
	programPointer := 0
	tapePointer := 0

	for programPointer < end {

		switch program[programPointer] {
		case '+':
			tape[tapePointer]++
			programPointer++
		case '.':
			buf.WriteByte(tape[tapePointer])
			programPointer++
		}

	}

	return buf.Bytes()
}
