package main

import (
	"fmt"
	"github.com/drafolin/advent-of-code/2024/utils"
	"math"
	"strconv"
	"strings"
)

func main() {
	//data := strings.Split("Register A: 729\nRegister B: 0\nRegister C: 0\n\nProgram: 0,1,5,4,3,0", "\n\n")
	data := strings.Split(utils.ReadInput("17th"), "\n\n")

	registersStrs := strings.Split(data[0], "\n")
	program := strings.Split(data[1], ",")
	program[0] = strings.Split(program[0], " ")[1]

	execPtr := 0
	output := ""

	registers := make(map[rune]int)
	for _, reg := range registersStrs {
		regSplit := strings.Split(reg, ": ")
		regName := strings.Split(regSplit[0], " ")[1]
		var err error
		registers[rune(regName[0])], err = strconv.Atoi(regSplit[1])
		if err != nil {
			panic("Invalid register value")
		}
	}

	for execPtr < len(program) {
		opcode := program[execPtr]
		operand := rune(program[execPtr+1][0])

		switch opcode {
		case "0":
			// adv
			numerator := registers['A']
			denominator := math.Pow(2, float64(comboOperand(operand, registers)))
			registers['A'] = numerator / int(denominator)
		case "1":
			// bxl
			op1 := registers['B']
			op2 := literalOperand(operand)
			registers['B'] = op1 ^ op2
		case "2":
			//bst
			registers['B'] = comboOperand(operand, registers) % 8
		case "3":
			// jnz
			if registers['A'] != 0 {
				execPtr = literalOperand(operand)
				continue
			}
		case "4":
			//bxc
			registers['B'] = registers['B'] ^ registers['C']
		case "5":
			// out
			if output != "" {
				output += ","
			}
			output += strconv.Itoa(comboOperand(operand, registers) % 8)
		case "6":
			// bdv
			numerator := registers['A']
			denominator := math.Pow(2, float64(comboOperand(operand, registers)))
			registers['B'] = numerator / int(denominator)
		case "7":
			// cdv
			numerator := registers['A']
			denominator := math.Pow(2, float64(comboOperand(operand, registers)))
			registers['C'] = numerator / int(denominator)
		}

		execPtr += 2
	}

	fmt.Println(output)
}

func comboOperand(operand rune, registers map[rune]int) int {
	switch operand {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return registers['A']
	case '5':
		return registers['B']
	case '6':
		return registers['C']
	case '7':
		panic("Reserved operand")
	}
	return 0
}

func literalOperand(operand rune) int {
	i, err := strconv.Atoi(string(operand))
	if err != nil {
		panic("Invalid operand")
	}
	return i
}
