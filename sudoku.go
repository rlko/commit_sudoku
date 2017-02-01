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
			if ((j + 1) % 3) == 0 && j != 8 {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
		if ((i + 1) % 3) == 0 && i != 8 {
			fmt.Println("---+---+---")
		}
	}
}

func line_has_duplication(arg string) bool {
	var s string

	for _, chr := range arg {
		if unicode.IsNumber(chr) {
			s = fmt.Sprintf("%c", chr)
			if strings.Count(arg, s) > 1 {
				return true
			}
		}
	}
	return false
}

func cols_have_duplication(args []string) bool {
	var i, j, y int

	for i = 0; i < 9; i++ {
		for j = 1; j < 9; j++ {
			if unicode.IsNumber(rune(args[j][i])) {
				for y = 0; y < 9; y++ {
					if j != y {
						if args[j][i] == args[y][i] {
							return true
						}
					}
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
			box[counter] = args[i + y][j + x]
			counter++
		}
	}
	return string(box[:])
}

func boxes_have_duplication(args []string) bool {
	var i, j int
//	var box string

	i = 0
	for i < 8 {
		j = 0
		for j < 8 {
//			box = extract_box(args, i, j)
//			fmt.Println(box)
			if line_has_duplication(extract_box(args, i, j)) {
				return true
			}
			j += 3
		}
		i += 3
	}
	return false
}

func has_minimum_required(args []string) bool {
	var counter int

	for _, arg := range args {
		for _, chr := range arg {
			if unicode.IsNumber(rune(chr)) {
				counter++
			}
		}
	}
	fmt.Print("\n")
	fmt.Print(counter)
	fmt.Println(" digits inside")
	if counter < 17 {
		return false
	}
	return true
}

func validate(args []string) {
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
	if !has_minimum_required(args) {
		exit_error("Error: there is not enough digits in the grid")
	}
	if cols_have_duplication(args) {
		exit_error("Error: there is a duplication in a column")
	}
	if boxes_have_duplication(args) {
		exit_error("Error: there is a duplication in a box")
	}
}

func resolve(args []string) {
// go go go
}

func main() {
	var args []string

	args = os.Args[1:]
	if len(args) != 9 {
		exit_error("Error: invalid grid")
	}

	print_grid(args)
	validate(args)
	resolve(args)
}
