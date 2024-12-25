// Jim Lawless
// https://github.com/jimlawless/aoc2024
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	//"slices"
    "strconv"
)

type Expr struct {
    isFunction bool
    lterm string
    rterm string
    op   string
    value int
}

func main() {
	var patGrammar = "[a-zA-Z0-9]+"
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
    accum:=0
    expressions:=make(map[string]Expr)   
	for scanner.Scan() {
		line := scanner.Text()
		tokens := reGrammar.FindAllString(line, -1)
		if len(tokens) == 0 {
			break
		}
        tmp,_:=strconv.Atoi(tokens[1])
        expressions[tokens[0]]=Expr { false,"","","",tmp }
    }
    for scanner.Scan() {
		line := scanner.Text()
		tokens := reGrammar.FindAllString(line, -1)
		if len(tokens) == 0 {
            continue
		}
        expressions[tokens[3]]=Expr {true,tokens[0],tokens[2],tokens[1],0}        
    }
    var zList []string    
    for i:=0;;i++ {
        suffix:=""
        if i<10 {
            suffix = "0"
        } 
        zName:="z"+suffix+strconv.Itoa(i)
        if _, ok:= expressions[zName]; ok {
            zList=append(zList,zName)
        } else {
            break
        }
    }
    for i:=len(zList)-1;i>=0;i-- {
        val:=eval(zList[i],expressions)
        accum=(accum<<1)|val
    }        
    fmt.Println(accum)
}

func eval(sym string,exprs map[string]Expr) int {
    e:=exprs[sym]
    if e.isFunction {
        lval:=eval(e.lterm,exprs)
        rval:=eval(e.rterm,exprs)
        switch e.op {
            case "AND":
                return lval & rval
            case "XOR":
                return lval ^ rval
            case "OR":
                return lval | rval
        }
    } else {
        return e.value
    }
    return 0
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}
