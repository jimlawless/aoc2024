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

var mulGrammar = "mul[(][\\d]+[,][\\d]+[)]"
var mulRegexp *regexp.Regexp
var numGrammar = "[\\d]+"
var numRegexp *regexp.Regexp

func main() {
	file, err := os.Open(os.Args[1])
	handleError(err)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	numRegexp, _ = regexp.Compile(numGrammar)
	mulRegexp, _ = regexp.Compile(mulGrammar)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		tokens := mulRegexp.FindAllString(line, -1)
		for i := 0; i < len(tokens); i++ {
			t2 := numRegexp.FindAllString(tokens[i], -1)
			m1, _ := strconv.Atoi(t2[0])
			m2, _ := strconv.Atoi(t2[1])
			sum += m1 * m2
		}
	}
	fmt.Println(sum)
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
