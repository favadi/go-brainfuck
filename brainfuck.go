package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	CELL_NUMBER = 30000
)

func main() {
	var inputFile string
	cells := [CELL_NUMBER]byte{}
	curCell := 0 // point to current cell, initialized to point to first cell

	flag.StringVar(&inputFile, "file", "", "path to program file")
	flag.Parse()

	if len(inputFile) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	program := string(buf)

	for i := 0; i < len(program); i++ {
		switch program[i] {
		case '+':
			cells[curCell] = cells[curCell] + 1
		case '-':
			cells[curCell] = cells[curCell] - 1
		case '>':
			curCell = curCell + 1
		case '<':
			curCell = curCell - 1
		case '.':
			fmt.Printf("%c", cells[curCell])
		case ',':
			var in_char byte
			_, err := fmt.Scanf("%c", &in_char)
			if err != nil {
				log.Fatal(err.Error())
			}
			cells[curCell] = in_char
		case '[':
			// find next ]
			if cells[curCell] == 0 {
				for j := i; j < len(program); j++ {
					if program[j] == ']' {
						i = j
						break
					}
					if j == len(program)-1 {
						log.Fatal("could not found matching [")
					}
				}
			}
		case ']':
			// find previous [
			if cells[curCell] != 0 {
				for j := i; j > 0; j-- {
					if program[j] == '[' {
						i = j
						break
					}
					if j == 0 {
						log.Fatal("could not found matching [")
					}
				}
			}
		}
	}
}
