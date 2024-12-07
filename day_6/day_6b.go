// Jim Lawless
// https://github.com/jimlawless/aoc2024
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	offsets := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	direction := 0

	file, err := os.Open(os.Args[1])
	handleError(err)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	sum := 0
	var lines []string
	x_org := 0
	y_org := 0
	for scanner.Scan() {
		line := scanner.Text()
		pos := strings.Index(line, "^")
		if pos >= 0 {
			y_org = len(lines)
			x_org = pos
		}
		lines = append(lines, line)
	}
	max_count := len(lines)*len(lines[0]) + 1
	for k := 0; k < len(lines); k++ {
		for j := 0; j < len(lines[k]); j++ {
			if lines[k][j] != '.' && lines[k][j] != 'O' {
				continue
			}
			direction = 0
			counter := 0
			x := x_org
			y := y_org
			for {
				counter = counter + 1
				if counter > max_count {
					sum = sum + 1
					break
				}
				tmpX := x + offsets[direction][0]
				tmpY := y + offsets[direction][1]
				if (tmpX >= len(lines[0])) || (tmpX < 0) {
					break
				}
				if (tmpY >= len(lines)) || (tmpY < 0) {
					break
				}
				if tmpY == k && tmpX == j {
					direction = (direction + 1) % len(offsets)
					continue
				}
				if (lines[tmpY][tmpX] == '.') || (lines[tmpY][tmpX] == '^') {
					x = tmpX
					y = tmpY
					continue
				}
				direction = (direction + 1) % len(offsets)
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
