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
	var disk []int
	scanner.Scan()
	line := scanner.Text()
	id := 0
	for i := 0; i < len(line); i++ {
		if (i % 2) == 0 {
			for j := 0; j < int(line[i]-'0'); j++ {
				disk = append(disk, id)
			}
			id++
		} else {
			for j := 0; j < int(line[i]-'0'); j++ {
				disk = append(disk, -1)
			}
		}
	}
	var free int
	for free = 0; free < len(disk); free++ {
		if disk[free] == -1 {
			break
		}
	}
	for last := len(disk) - 1; last >= 0; last-- {
		if last < free {
			break
		}
		if last == -1 {
			continue
		}
		disk[free] = disk[last]
		disk[last] = -1
		for ; free < len(disk); free++ {
			if disk[free] == -1 {
				break
			}
		}
	}
	checksum := 0
	fmt.Println(len(disk))
	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			checksum = checksum + (i * disk[i])
		}
	}
	fmt.Println("\n", checksum)
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
