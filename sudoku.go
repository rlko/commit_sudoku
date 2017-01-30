// Sudoku

package main

import (
	"os"
	"fmt"
//	"reflect"
//	"flag"
	"regexp"
)

func exit_error(str string) {
	fmt.Println(str)
	os.Exit(1)
}

func main() {

	var args		[]string
	var length		int

	args = os.Args[1:]
	length = len(args)
	if length < 1 {
		exit_error("Error: No args input :(");
	}

	var match bool
	for i := 0; i < length; i++ {
		match, _ = regexp.MatchString("^[1-9\\.]+$", args[i])
		if match == false {
			exit_error("Error: Invalid character");
		}
		if len(args[i]) != 9 {
			exit_error("Error")
		}
	}
}
