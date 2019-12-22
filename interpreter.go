package main

import (
	"bytes"
	"fmt"
)

// Run executes a brainfuck program and returns its output.
func Run(program []byte) ([]byte, error) {
	output := new(bytes.Buffer)
	tape := make(map[int]byte)

	end := len(program)
	programPointer := 0
	tapePointer := 0

	loopMarkers, err := scanLoops(program)
	if err != nil {
		return nil, fmt.Errorf("unable to run program: %w", err)
	}

	loops := NewStack()

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
		case '[':
			if tape[tapePointer] == 0 {
				programPointer = loopMarkers[programPointer]
			} else {
				loops.Push(programPointer)
			}
		case ']':
			programPointer, _ = loops.Pop()
			continue
		}

		programPointer++
	}

	return output.Bytes(), nil
}

// scanLoops creates a dictionary which maps start positions of loops to their respective end positions
func scanLoops(program []byte) (map[int]int, error) {
	loopMarkers := make(map[int]int)

	s := NewStack()
	for i, v := range program {
		switch v {
		case '[':
			s.Push(i)
		case ']':
			start, err := s.Pop()
			if err != nil {
				return nil, fmt.Errorf("unbalanced loop end at index %d", i)
			}

			loopMarkers[start] = i
		}
	}

	start, err := s.Pop()

	if err == nil {
		return nil, fmt.Errorf("unbalanced loop start at index %d", start)
	}

	return loopMarkers, nil
}
