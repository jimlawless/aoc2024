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

	rules := make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := reGrammar.FindAllString(line, -1)
		if len(tokens) == 0 {
			break
		}
		rules[tokens[0]+tokens[1]] = true
	}

	for scanner.Scan() {
		line := scanner.Text()
		tokens := reGrammar.FindAllString(line, -1)
		valid := true
		for i := 0; i < len(tokens); i++ {
			if i > 0 {
				if rules[tokens[i]+tokens[i-1]] {
					valid = false
					break
				}
			}
			if i < (len(tokens) - 1) {
				if rules[tokens[i+1]+tokens[i]] {
					valid = false
					break
				}
			}
		}
		if valid {
			middle, _ := strconv.Atoi(tokens[(len(tokens) / 2)])
			sum = sum + middle
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
