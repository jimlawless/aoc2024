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

var patGrammar = "(Button[\\s][A|B]|Prize|[\\d]+)"
var reGrammar *regexp.Regexp

var minimum int

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
	minimum := -1
	ax := 0
	ay := 0
	bx := 0
	by := 0
	xVal := 0
	yVal := 0

	for scanner.Scan() {

		line := scanner.Text()
		words := reGrammar.FindAllString(line, -1)
		if len(words) == 0 {
			continue
		}
		switch words[0] {
		case "Button A":
			ax, _ = strconv.Atoi(words[1])
			ay, _ = strconv.Atoi(words[2])
			continue
		case "Button B":
			bx, _ = strconv.Atoi(words[1])
			by, _ = strconv.Atoi(words[2])
			continue
		case "Prize":
			xVal, _ = strconv.Atoi(words[1])
			yVal, _ = strconv.Atoi(words[2])
		}
		minimum = -1
		for a := 0; a <= 100; a++ {
			for b := 0; b <= 100; b++ {
				xDistance := ax*a + bx*b
				yDistance := ay*a + by*b
				if (xDistance) > xVal {
					continue
				}
				if (yDistance) > yVal {
					continue
				}
				if xDistance == xVal && yDistance == yVal {
					tokens := a*3 + b
					if minimum == -1 {
						minimum = tokens
					} else {
						if tokens < minimum {
							minimum = tokens
						}
					}
				}
			}
		}
		if minimum > 0 {
			sum += minimum
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
