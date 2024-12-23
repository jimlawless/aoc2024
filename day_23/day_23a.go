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
)

func main() {
	var patGrammar = "[a-z]+"
	var reGrammar *regexp.Regexp

	file, err := os.Open(os.Args[1])
	handleError(err)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	ndx := make(map[string][]string)
	exists := make(map[string]bool)
	groupExists := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	reGrammar, _ = regexp.Compile(patGrammar)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		tokens := reGrammar.FindAllString(line, -1)
		if len(tokens) == 0 {
			continue
		}
		if ndx[tokens[0]] == nil {
			ndx[tokens[0]] = []string{tokens[1]}
		} else {
			ndx[tokens[0]] = append(ndx[tokens[0]], tokens[1])
		}
		exists[tokens[0]+tokens[1]] = true
		exists[tokens[1]+tokens[0]] = true
	}
	var filtered []string
	for key, value := range ndx {
		first := key
		nodelist := value
		for j := 0; j < len(nodelist); j++ {
			second := nodelist[j]
			secondlist := ndx[second]
			for k := 0; k < len(secondlist); k++ {
				third := secondlist[k]
				if exists[first+third] || exists[third+first] {
					lst := []string{first, second, third}
					slices.Sort(lst)
					group := lst[0] + "," + lst[1] + "," + lst[2]
					if !groupExists[group] {
						groupExists[group] = true
						filtered = append(filtered, group)
						if first[0] == 't' || second[0] == 't' || third[0] == 't' {
							sum++
						}
					}
				}
			}
		}
	}
	slices.Sort(filtered)
	for i := 0; i < len(filtered); i++ {
		fmt.Println(filtered[i])
	}
	fmt.Println("\n", sum)
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
