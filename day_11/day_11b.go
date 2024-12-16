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
var memo map[string]int

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
	memo = make(map[string]int)
	for i := 0; i < len(tokens); i++ {
		face, _ := strconv.Atoi(tokens[i])
		count += countStones(face, 0, 75)
	}
	fmt.Println(count)
}

func countStones(face int, depth int, max int) int {
	stones := 0
	if depth == max {
		return 1
	}

	key := strconv.Itoa(face) + "_" + strconv.Itoa(depth)
	if memo[key] != 0 {
		return memo[key]
	}
	if face == 0 {
		stones = countStones(1, depth+1, max)
		memo[key] = stones
		return stones
	}
	str := strconv.Itoa(face)
	if len(str)%2 == 0 {
		mid := len(str) / 2
		left, _ := strconv.Atoi(str[0:mid])
		right, _ := strconv.Atoi(str[mid:])
		stones = countStones(left, depth+1, max) +
			countStones(right, depth+1, max)
		memo[key] = stones
		return stones
	}
	stones = countStones(face*2024, depth+1, max)
	memo[key] = stones
	return stones
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
