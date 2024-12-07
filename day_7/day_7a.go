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

var patGrammar = "[\\d]+"
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
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		tokens := reGrammar.FindAllString(line, -1)
		result, _ := strconv.Atoi(tokens[0])
		r1, _ := strconv.Atoi(tokens[1])
		if addMul(result, r1, 2, tokens) {
			sum = sum + result
		}
	}
	fmt.Println(sum)
}

func addMul(result int, r1 int, depth int, tokens []string) bool {
	if depth >= len(tokens) {
		return false
	}
	r2, _ := strconv.Atoi(tokens[depth])
	if (r1 * r2) == result {
		return true
	} else {
		b := addMul(result, r1*r2, depth+1, tokens)
		if b {
			return true
		}
	}
	if (r1 + r2) == result {
		return true
	} else {
		b := addMul(result, r1+r2, depth+1, tokens)
		if b {
			return true
		}
	}
	return false
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
