// Sudoku

package main

import (
	"fmt"
	"os"
	//	"reflect"
	//	"flag"
	"regexp"
	"strings"
	"unicode"
)

func exit_error(str string) {
	fmt.Println(str)
	os.Exit(1)
}

func print_grid(args []string) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", args[i][j])
			if (j+1)%3 == 0 && j != 8 {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
		if (i+1)%3 == 0 && i != 8 {
			fmt.Println("---+---+---")
		}
	}
}

func line_has_duplication(arg string) bool {
	for _, chr := range arg {
		if unicode.IsNumber(chr) {
			s := fmt.Sprintf("%c", chr)
			if strings.Count(arg, s) > 1 {
				return true
			}
		}
	}
	return false
}

func cols_have_duplication(args []string) bool {
	var i, j int
	var c byte

	for i = 0; i < 8; i++ {
		c = args[0][i]
		if unicode.IsNumber(rune(c)) {
			for j = 1; j < 8; j++ {
				if args[j][i] == c {
					return true
				}
			}
		}
	}
	return false
}

func extract_box(args []string, i int, j int) string {
	var box [9]byte
	var counter int

	counter = 0
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			box[counter] = args[i+y][j+x]
			counter++
		}
	}
	return string(box[:])
}

func boxes_have_duplication(args []string) bool {
	var i, j int
	var box string

	i = 0
	for i < 8 {
		j = 0
		for j < 8 {
			box = extract_box(args, i, j)
			//			fmt.Println(box)
			if line_has_duplication(box) {
				return true
			}
			j += 3
		}
		i += 3
	}
	return false
}

func main() {
	var args []string

	args = os.Args[1:]
	if len(args) != 9 {
		exit_error("Error: invalid grid")
	}

	print_grid(args)
	var match bool
	for _, arg := range args {
		match, _ = regexp.MatchString("^[1-9\\.]+$", arg)
		if match == false {
			exit_error("Error: invalid character")
		}
		if len(arg) != 9 {
			exit_error("Error: arg has less than 9 characters")
		}
		if line_has_duplication(arg) {
			exit_error("Error: there is a duplication in a line")
		}
	}
	if cols_have_duplication(args) {
		exit_error("Error: there is a duplication in a column")
	}
	if boxes_have_duplication(args) {
		exit_error("Error: there is a duplication in a box")
	}
}
