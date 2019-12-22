package main

import "bytes"

// Run executes a brainfuck program and returns its output.
func Run(program []byte) []byte {
	output := new(bytes.Buffer)
	tape := make(map[int]byte)

	end := len(program)
	programPointer := 0
	tapePointer := 0

	for programPointer < end {

		switch program[programPointer] {
		case '+':
			tape[tapePointer]++
		case '-':
			tape[tapePointer]--
		case '>':
			tapePointer++
		case '<':
			tapePointer--
		case '.':
			output.WriteByte(tape[tapePointer])
		}

		programPointer++
	}

	return output.Bytes()
}
