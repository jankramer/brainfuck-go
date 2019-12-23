package main

import (
	"fmt"
	"io"
)

// Run executes a given brainfuck program.
func Run(program []byte, input io.Reader, output io.Writer) error {
	memory := make(map[int]byte)
	var pInstr, pMemory int

	loopPositions, err := scanLoops(program)
	if err != nil {
		return fmt.Errorf("unable to run program: %w", err)
	}

	loops := NewStack()

	for pInstr < len(program) {

		switch program[pInstr] {
		case '+':
			memory[pMemory]++
		case '-':
			memory[pMemory]--
		case '>':
			pMemory++
		case '<':
			pMemory--
		case '.':
			if _, err := output.Write([]byte{memory[pMemory]}); err != nil {
				return err
			}
		case ',':
			buf := make([]byte, 1)
			_, err := input.Read(buf)

			switch err {
			case nil:
				memory[pMemory] = buf[0]
			case io.EOF:
				memory[pMemory] = 0
			default:
				return fmt.Errorf("unable to read from input: %w", err)
			}

		case '[':
			if memory[pMemory] == 0 {
				pInstr = loopPositions[pInstr]
			} else {
				loops.Push(pInstr)
			}
		case ']':
			pInstr, _ = loops.Pop()
			continue
		}

		pInstr++
	}

	return nil
}

// scanLoops creates a dictionary which maps start positions of loops to their respective end positions
func scanLoops(program []byte) (map[int]int, error) {
	loopPositions := make(map[int]int)

	s := NewStack()
	for i, instr := range program {
		switch instr {
		case '[':
			s.Push(i)
		case ']':
			start, err := s.Pop()
			if err != nil {
				return nil, fmt.Errorf("unbalanced loop end at index %d", i)
			}

			loopPositions[start] = i
		}
	}

	start, err := s.Pop()

	if err == nil {
		return nil, fmt.Errorf("unbalanced loop start at index %d", start)
	}

	return loopPositions, nil
}
