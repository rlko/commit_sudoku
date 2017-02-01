// Sudoku

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

type coord struct {
	i	int
	j	int
}

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
	fmt.Print("\n")
}

func line_has_duplication(arg string) bool {
	var s string

	for _, chr := range arg {
		if unicode.IsNumber(chr) {
			s = string(chr)
			if strings.Count(arg, s) > 1 {
				return true
			}
		}
	}
	return false
}

func col_has_duplication(args []string, col int) bool {
	var i, y int
	for i = 0; i < 9; i++ {
		if unicode.IsNumber(rune(args[i][col])) {
			for y = 0; y < 9; y++ {
				if i != y {
					if args[i][col] == args[y][col] {
						return true
					}
				}
			}
		}
	}
	return false
}

func cols_have_duplication(args []string) bool {
	var i int

	for i = 0; i < 9; i++ {
		if col_has_duplication(args, i) {
			return true
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

func box_has_duplication(args[]string, i int, j int) bool {
	return (line_has_duplication(extract_box(args, i, j)))
}

func boxes_have_duplication(args []string) bool {
	var i, j int

	i = 0
	for i < 9 {
		j = 0
		for j < 9 {
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
	if counter < 17 {
		return false
	}
	return true
}

func validate_grid(args []string) {
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

func digit_is_valid(args []string, chars []byte, cd coord, d byte) bool {
	chars[cd.j] = d + '0'
	args[cd.i] = string(chars)
	if line_has_duplication(args[cd.i]) || col_has_duplication(args, cd.j) || boxes_have_duplication(args) {
		chars[cd.j] = '.'
		args[cd.i] = string(chars)
		return false
	}
	return true
}

func try_digit(args []string, chars []byte, cd coord) bool {
	var d int

	for d = 1; d <= 9; d++ {
		if digit_is_valid(args, chars, cd, byte(d)) {
			if resolve(args) {
				return true
			}
		}
	}
	return false
}

func resolve(args []string) bool{
	var chars	[]byte
	var cd		coord

	for cd.i = 0; cd.i < 9; cd.i++ {
		chars = []byte(args[cd.i])
		for cd.j = 0; cd.j < 9; cd.j++ {
			if args[cd.i][cd.j] == '.' {
				if !try_digit(args, chars, cd) {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	var args []string

	args = os.Args[1:]
	if len(args) != 9 {
		exit_error("Error: invalid grid")
	}

	validate_grid(args)
	resolve(args)
	print_grid(args)
}
