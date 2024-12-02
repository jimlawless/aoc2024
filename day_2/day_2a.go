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
		var lst []int
		for i := 0; i < len(tokens); i++ {
			tmp, _ := strconv.Atoi(tokens[i])
			lst = append(lst, tmp)
		}
		if lst[0] > lst[1] {
			if checkDescending(lst) {
				sum = sum + 1
			}
		} else {
			if checkAscending(lst) {
				sum = sum + 1
			}
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

func checkDescending(z []int) bool {
	for i := 0; i < (len(z) - 1); i++ {
		diff := z[i] - z[i+1]
		if (diff < 1) || (diff > 3) {
			return false
		}
	}
	return true
}

func checkAscending(z []int) bool {
	for i := 0; i < (len(z) - 1); i++ {
		diff := z[i+1] - z[i]
		if (diff < 1) || (diff > 3) {
			return false
		}
	}
	return true
}
