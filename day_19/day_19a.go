// Jim Lawless
// https://github.com/jimlawless/aoc2024
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var patGrammar = "[a-z]+"
var reGrammar *regexp.Regexp

func main() {
	file, err := os.Open(os.Args[1])
	handleError(err)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	reGrammar, _ = regexp.Compile(patGrammar)

	var dynRegex *regexp.Regexp
	var dynPattern string
	scanner := bufio.NewScanner(file)
	dynRegex, _ = regexp.Compile(dynPattern)
	scanner.Scan()

	// Build dynamic regexp
	line := scanner.Text()
	tokens := reGrammar.FindAllString(line, -1)
	dynPattern = "^("
	for i := 0; i < len(tokens); i++ {
		if i > 0 {
			dynPattern += "|"
		}
		dynPattern += tokens[i]
	}
	dynPattern += ")+$"
	fmt.Println("Dynamic regexp is ", dynPattern)
	dynRegex, _ = regexp.Compile(dynPattern)

	// Read past empty line
	scanner.Scan()

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if dynRegex.MatchString(line) {
			count++
		}
	}
	fmt.Println("\n", count)
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
