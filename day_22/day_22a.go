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



func main() {
    var patGrammar = "[\\d]+"
    var reGrammar *regexp.Regexp	
    file, err := os.Open(os.Args[1])
	handleError(err)
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	reGrammar, _ = regexp.Compile(patGrammar)
    sum:=0
	for scanner.Scan() {
		line := scanner.Text()
		tokens := reGrammar.FindAllString(line, -1)
		if len(tokens) == 0 {
			continue
		}
        secret,_:=strconv.Atoi(tokens[0])
        for i:=0; i<2000; i++ {
            secret^=(secret*64)
            secret%=16777216
            secret^=(secret/32)
            secret%=16777216
            secret^=secret*2048
            secret%=16777216                              
        }
        sum+=secret
        fmt.Println(tokens[0],": " ,secret)                      
    }
	fmt.Println(sum)
}


func handleError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
