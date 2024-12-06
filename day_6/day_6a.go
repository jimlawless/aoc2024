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

	sum := 1
	var lines []string
	x := 0
	y := 0
	taken := make(map[int]bool)

	for scanner.Scan() {
		line := scanner.Text()
		pos := strings.Index(line, "^")
		if pos >= 0 {
			y = len(lines)
			x = pos
		}
		lines = append(lines, line)
	}
	taken[y*1000+x] = true
	for {
		tmpX := x + offsets[direction][0]
		tmpY := y + offsets[direction][1]
		if (tmpX >= len(lines[0])) || (tmpX < 0) {
			if (lines[y][x] == '.') || (lines[y][x] == '^') {
				if !taken[y*1000+x] {
					sum = sum + 1
				}
				taken[y*1000+x] = true
			}
			break
		}
		if (tmpY >= len(lines)) || (tmpY < 0) {
			if (lines[y][x] == '.') || (lines[y][x] == '^') {
				if !taken[y*1000+x] {
					sum = sum + 1
				}
				taken[y*1000+x] = true
			}
			break
		}
		if (lines[tmpY][tmpX] == '.') || (lines[tmpY][tmpX] == '^') {
			if !taken[y*1000+x] {
				sum = sum + 1
			}
			taken[y*1000+x] = true
			x = tmpX
			y = tmpY
			continue
		}
		direction = (direction + 1) % len(offsets)
	}
	fmt.Println(sum)
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
