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
	scanner.Scan()
	line := scanner.Text()
	tokens := reGrammar.FindAllString(line, -1)
	count := 0
	for i := 0; i < len(tokens); i++ {
		face, _ := strconv.Atoi(tokens[i])
		count += countStones(face, 0, 25)
	}
	fmt.Println(count)
}

func countStones(face int, depth int, max int) int {
	if depth == max {
		return 1
	}
	if face == 0 {
		return countStones(1, depth+1, max)
	}
	str := strconv.Itoa(face)
	if len(str)%2 == 0 {
		mid := len(str) / 2
		left, _ := strconv.Atoi(str[0:mid])
		right, _ := strconv.Atoi(str[mid:])
		return countStones(left, depth+1, max) +
			countStones(right, depth+1, max)
	}
	return countStones(face*2024, depth+1, max)
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
