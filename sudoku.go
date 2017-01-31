// Sudoku

package main

import (
	"os"
	"fmt"
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

func has_duplication_l(arg string) bool {
	for _, chr := range arg {
		if unicode.IsNumber(chr) {
			s  := fmt.Sprintf("%c", chr)
			if strings.Count(arg, s) > 1 {
				return true
			}
		}
	}
	return false
}

func has_duplication_c(args []string) bool {
	var i, j	int
	var c		byte

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


func main() {

	var args		[]string

	args = os.Args[1:]
	if len(args) != 9 {
		exit_error("Error: invalid grid")
	}

	var match bool
	for _, arg := range args {
		match, _ = regexp.MatchString("^[1-9\\.]+$", arg)
		if match == false {
			exit_error("Error: invalid character")
		}
		if len(arg) != 9 {
			exit_error("Error: arg has less than 9 characters")
		}
		if has_duplication_l(arg) {
			exit_error("Error: there is a duplication in a line")
		}
	}
	if has_duplication_c(args) {
			exit_error("Error: there is a duplication in a column")
	}
}
