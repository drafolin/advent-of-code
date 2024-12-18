package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"slices"
	"strconv"
	"strings"
)

func main() {
	//data := strings.Split("Register A: 729\nRegister B: 0\nRegister C: 0\n\nProgram: 0,1,5,4,3,0", "\n\n")
	data := strings.Split(utils.ReadInput("17th"), "\n\n")

	registerA := strings.Split(data[0], "\n")[0]
	program := strings.Split(data[1], ",")
	program[0] = strings.Split(program[0], " ")[1]
	programUint := make([]uint64, len(program))
	for i, v := range program {
		var err error
		programUint[i], err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	seed, err := strconv.ParseUint(strings.Split(registerA, ": ")[1], 10, 64)
	if err != nil {
		panic("Invalid seed value")
	}
	fmt.Println(programUint, seed)
	fmt.Println(partOne(programUint, seed))
	fmt.Println(partTwo(programUint))
}

func partTwo(program []uint64) (seed uint64) {
	for itr := len(program) - 1; itr >= 0; itr-- {
		seed <<= 3
		for !slices.Equal(partOne(program, seed), program[itr:]) {
			seed++
		}
	}

	return
}

func partOne(program []uint64, seed uint64) (res []uint64) {
	execPtr := 0

	registers := map[rune]uint64{
		'A': seed,
		'B': 0,
		'C': 0,
	}

	for execPtr < len(program)-1 {
		operand := program[execPtr+1]

		switch opcode := program[execPtr]; opcode {
		case 0: // adv
			registers['A'] >>= comboOperand(operand, registers)
		case 1: // bxl
			registers['B'] ^= operand
		case 2: //bst
			registers['B'] = comboOperand(operand, registers) & 7
		case 3: // jnz
			if registers['A'] != 0 {
				execPtr = int(operand)
				continue
			}
		case 4: //bxc
			registers['B'] ^= registers['C']
		case 5: // out
			res = append(res, comboOperand(operand, registers)&7)
		case 6: // bdv
			registers['B'] = registers['A'] >> comboOperand(operand, registers)
		case 7: // cdv
			registers['C'] = registers['A'] >> comboOperand(operand, registers)
		}

		execPtr += 2
	}
	return
}

func comboOperand(operand uint64, registers map[rune]uint64) uint64 {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return registers['A']
	case 5:
		return registers['B']
	case 6:
		return registers['C']
	default:
		panic("Reserved operand")
	}
}
