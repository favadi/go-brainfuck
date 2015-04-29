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
	data_ptr := 0 // point to current cell, initialized to point to first cell

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
	test_str := string(buf)

	for i := 0; i < len(test_str); i++ {
		switch test_str[i] {
		case '+':
			cells[data_ptr] = cells[data_ptr] + 1
		case '-':
			cells[data_ptr] = cells[data_ptr] - 1
		case '>':
			data_ptr = data_ptr + 1
		case '<':
			data_ptr = data_ptr - 1
		case '.':
			fmt.Printf("%c", cells[data_ptr])
		case ',':
			var in_char byte
			_, err := fmt.Scanf("%c", &in_char)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			cells[data_ptr] = in_char
		case '[':
			// find next ]
			if cells[data_ptr] == 0 {
				for j := i; j < len(test_str); j++ {
					if test_str[j] == ']' {
						i = j
						break
					}
					if j == len(test_str)-1 {
						fmt.Println("could not found matching [")
						os.Exit(1)
					}
				}
			}
		case ']':
			// find previous [
			if cells[data_ptr] != 0 {
				for j := i; j > 0; j-- {
					if test_str[j] == '[' {
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
