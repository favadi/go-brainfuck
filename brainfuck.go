package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	CELL_NUMBER = 30000
)

func main() {
	cells := [CELL_NUMBER]byte{}
	curCell := 0 // point to current cell, initialized to point to first cell

	inputFile := flag.String("file", "", "path to program file")
	flag.Parse()

	if *inputFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	buf, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
				fmt.Println(err)
				os.Exit(1)
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
						fmt.Println("could not found matching [")
						os.Exit(1)
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
						fmt.Println("could not found matching [")
						os.Exit(1)
					}
				}
			}
		}
	}
}
