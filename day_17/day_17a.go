// Jim Lawless
// https://github.com/jimlawless/aoc2024
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var patGrammar = "(Register[\\s][A|B|C]|Program|[\\d]+)"
var reGrammar *regexp.Regexp

func main() {
	file, err := os.Open(os.Args[1])
	handleError(err)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	reGrammar, _ = regexp.Compile(patGrammar)

	regA := 0
	regB := 0
	regC := 0

	next := 0
	needComma := false

	var code []string
	powers := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for scanner.Scan() {

		line := scanner.Text()
		tokens := reGrammar.FindAllString(line, -1)
		if len(tokens) == 0 {
			continue
		}
		switch tokens[0] {
		case "Register A":
			regA, _ = strconv.Atoi(tokens[1])
			continue

		case "Register B":
			regB, _ = strconv.Atoi(tokens[1])
			continue

		case "Register C":
			regC, _ = strconv.Atoi(tokens[1])
			continue

		case "Program":
			code = tokens[1:]
		}
		for i := 0; i < len(code); i = next {
			next = i + 2
			//fmt.Println("\n",i,regA,regB,regC)
			switch code[i] {
			case "0": // adv
				regA = regA / powers[combo(code[i+1], regA, regB, regC)]

			case "1": // bxl
				op, _ := strconv.Atoi(code[i+1])
				regB = regB ^ op

			case "2": // bst
				regB = combo(code[i+1], regA, regB, regC) % 8

			case "3": // jnz
				if regA != 0 {
					op, _ := strconv.Atoi(code[i+1])
					next = op
				}

			case "4": // bxc
				regB = regB ^ regC

			case "5": // out
				if needComma {
					fmt.Print(",")
				} else {
					needComma = true
				}
				fmt.Print(combo(code[i+1], regA, regB, regC) % 8)

			case "6": // bdv
				regB = regA / powers[combo(code[i+1], regA, regB, regC)]

			case "7": // cdv
				regC = regA / powers[combo(code[i+1], regA, regB, regC)]

			default:
				fmt.Println("Invalid opcode ", code[i])
				os.Exit(1)
			}
		}
	}
	fmt.Println()
}

func combo(op string, regA int, regB int, regC int) int {
	intOp, _ := strconv.Atoi(op)
	if intOp >= 0 && intOp <= 3 {
		return intOp
	}
	switch intOp {
	case 4:
		return regA
	case 5:
		return regB
	case 6:
		return regC
	default:
		fmt.Println("Invalid combo operand ", intOp)
		os.Exit(1)
	}
	return -1
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
