// Jim Lawless
// https://github.com/jimlawless/aoc2024
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
    "slices"
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
    var left[]int
    var right[]int
	for scanner.Scan() {
		line := scanner.Text()
		tokens := reGrammar.FindAllString(line, -1)
        tmp,_:=strconv.Atoi(tokens[0])
        left=append(left,tmp)
        tmp,_=strconv.Atoi(tokens[1])
        right=append(right,tmp)
	}
    slices.Sort(left)
    slices.Sort(right)
    for i:=0;i<len(left);i++ {
        if left[i]>right[i] {
            sum+=left[i]-right[i]
        } else {
            sum+=right[i]-left[i]
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
