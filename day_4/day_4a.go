// Jim Lawless
// https://github.com/jimlawless/aoc2024
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	handleError(err)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sum := 0

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == 'X' {
				if isXmas(x, y, -1, -1, 0, lines) {
					sum = sum + 1
				}
				if isXmas(x, y, 0, -1, 0, lines) {
					sum = sum + 1
				}
				if isXmas(x, y, 1, -1, 0, lines) {
					sum = sum + 1
				}
				if isXmas(x, y, 1, 0, 0, lines) {
					sum = sum + 1
				}
				if isXmas(x, y, 1, 1, 0, lines) {
					sum = sum + 1
				}
				if isXmas(x, y, 0, 1, 0, lines) {
					sum = sum + 1
				}
				if isXmas(x, y, -1, 1, 0, lines) {
					sum = sum + 1
				}
				if isXmas(x, y, -1, 0, 0, lines) {
					sum = sum + 1
				}
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

func isXmas(x int, y int, xInc int, yInc int, depth int, arr []string) bool {
	word := "MAS"
	if depth >= len(word) {
		return true
	}

	if x+xInc < 0 {
		return false
	}
	if y+yInc < 0 {
		return false
	}

	if x+xInc >= len(arr[y]) {
		return false
	}
	if y+yInc >= len(arr) {
		return false
	}
	if word[depth] == arr[y+yInc][x+xInc] {
		return isXmas(x+xInc, y+yInc, xInc, yInc, depth+1, arr)
	}
	return false
}
