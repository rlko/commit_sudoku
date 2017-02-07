// Sudoku

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"bytes"
	"math/rand"
	"strconv"
	"time"
)

type t_coord struct {
	i int
	j int
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: sudoku [-r] [-mode=input_mode] [input]\n")
	fmt.Fprintf(os.Stderr, "       sudoku -c [-r] [-diff=difficulty_mode]\n")
	flag.PrintDefaults()
}

func print_grid(grid []string, raw bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", grid[i][j])
			if !raw && ((j+1)%3) == 0 && j != 8 {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
		if !raw && ((i+1)%3) == 0 && i != 8 {
			fmt.Println("---+---+---")
		}
	}
}

func has_minimum_required(grid []string) bool {
	var counter int

	for _, line := range grid {
		for _, chr := range line {
			if unicode.IsDigit(rune(chr)) {
				counter++
			}
		}
	}
	return counter > 16
}

func permut_digits(a, b int, grid []string) []string {
	var i int

	if a == b {
		return grid
	}
	for i = 0; i < 9; i++ {
		grid[i] = strings.Replace(grid[i], strconv.Itoa(a), "0", -1)
	}
	for i = 0; i < 9; i++ {
		grid[i] = strings.Replace(grid[i], strconv.Itoa(b), strconv.Itoa(a), -1)
	}
	for i = 0; i < 9; i++ {
		grid[i] = strings.Replace(grid[i], "0", strconv.Itoa(b), -1)
	}
	return grid
}

func permut_lines(a, b int, grid []string) {
	var tmp string

	if a == b {
		return
	}
	tmp = grid[a]
	grid[a] = grid[b]
	grid[b] = tmp
}

func cols_to_lines(grid []string) []string {
	var i, j, k int
	var ngrid []string
	var buffer bytes.Buffer

	for i = 0; i < 9; i++ {
		for k = 0; k < 9; k++ {
			for j = 0; j < 9; j++ {
				if i == j {
					buffer.WriteByte(grid[k][i])
					break
				}
			}
		}
		ngrid = append(ngrid, buffer.String())
		buffer.Reset()
	}
	return ngrid
}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func shuffle(grid []string) {
	var i int
	var rdm int

	rdm = random(1, 3)
	for i = 0; i < rdm; i++ {
		permut_lines(random(0, 3), random(0, 3), grid)
		permut_lines(random(3, 6), random(3, 6), grid)
		permut_lines(random(6, 9), random(6, 9), grid)
	}
	rdm = random(1, 3)
	grid = cols_to_lines(grid)
	for i = 0; i < rdm; i++ {
		permut_lines(random(0, 3), random(0, 3), grid)
		permut_lines(random(3, 6), random(3, 6), grid)
		permut_lines(random(6, 9), random(6, 9), grid)
	}
	rdm = random(1, 3)
	for i = 0; i < rdm; i++ {
		permut_digits(random(1, 10), random(1, 10), grid)
		permut_digits(random(1, 10), random(1, 10), grid)
		permut_digits(random(1, 10), random(1, 10), grid)
		permut_digits(random(1, 10), random(1, 10), grid)
	}
}

func unfill(grid []string, diff string) bool {
	var chars []byte
	var i, j int
	var counter, max int

	if diff == "easy" {
		max = random(40, 50)
	} else if diff == "hard" {
		max = random(60, 64)
	} else {
		max = random(50, 60)
	}
	for {
		i = random(0, 9)
		j = random(0, 9)
		chars = []byte(grid[i])
		if grid[i][j] != '.' {
			if counter == max {
				break
			}
			chars[j] = '.'
			grid[i] = string(chars)
			counter++
		}
	}
	if has_minimum_required(grid) {
		return true
	}
	return false
}

func commit_sudoku(create, raw bool, diff, mode string) bool {
	if create {
		var ngrid = []string{"892546371", "367218594", "514793268", "641357982", "985421736", "723689415", "159872643", "238964157", "476135829"}

		if mode != "" {
			fmt.Fprintf(os.Stderr, "Warning: -mode flag does nothing when generating grid\n")
		}
		if len(flag.Args()) > 0 {
			fmt.Fprintf(os.Stderr, "Warning: Ignoring all arguments\n")
		}
		shuffle(ngrid)
		if unfill(ngrid, diff) {
			print_grid(ngrid, raw)
			return true
		}
		fmt.Println("Error: Unexpected error")
		os.Exit(1)
	}
	return false
}

func get_grid(mode string) []string {
	var grid []string

	if mode == "piscine" {
		grid = flag.Args()
		if len(grid) != 9 {
			fmt.Println("Error: invalid grid")
			return nil
		}
	} else {
		var l int = len(flag.Args())
		var path string

		if l < 1 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				grid = append(grid, scanner.Text())
			}
			return grid
		}
		if l > 1 {
			fmt.Fprintf(os.Stderr, "Warning: Ignoring all arguments but the first one\n")
		}
		path = flag.Args()[0]
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			grid = append(grid, scanner.Text())
		}
		file.Close()
		if len(grid) != 9 {
			fmt.Println("Error: invalid grid")
			return nil
		}
	}
	return grid
}

func line_has_duplication(line string) bool {
	var s string

	for _, chr := range line {
		if unicode.IsDigit(chr) {
			s = string(chr)
			if strings.Count(line, s) > 1 {
				return true
			}
		}
	}
	return false
}

func col_has_duplication(grid []string, col int) bool {
	var i, y int

	for i = 0; i < 9; i++ {
		if unicode.IsDigit(rune(grid[i][col])) {
			for y = 0; y < 9; y++ {
				if i != y {
					if grid[i][col] == grid[y][col] {
						return true
					}
				}
			}
		}
	}
	return false
}

func cols_have_duplication(grid []string) bool {
	var i int

	for i = 0; i < 9; i++ {
		if col_has_duplication(grid, i) {
			return true
		}
	}
	return false
}

func extract_box(grid []string, i, j int) string {
	var box [9]byte
	var counter int

	counter = 0
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			box[counter] = grid[i+x][j+y]
			counter++
		}
	}
	return string(box[:])
}

func box_has_duplication(grid []string, i, j int) bool {
	return (line_has_duplication(extract_box(grid, i, j)))
}

func boxes_have_duplication(grid []string) bool {
	var i, j int

	i = 0
	for i < 9 {
		j = 0
		for j < 9 {
			if line_has_duplication(extract_box(grid, i, j)) {
				return true
			}
			j += 3
		}
		i += 3
	}
	return false
}

func validate_grid(grid []string) bool {
	var match bool

	for _, line := range grid {
		match, _ = regexp.MatchString("^[1-9\\.]+$", line)
		if match == false {
			fmt.Println("Error: invalid character")
			return false
		}
		if len(line) != 9 {
			fmt.Println("Error: there is not 9 characters in a line")
			return false
		}
		if line_has_duplication(line) {
			fmt.Println("Error: there is a duplication in a line")
			return false
		}
	}
	if !has_minimum_required(grid) {
		fmt.Println("Error: there is not enough digits in the grid")
		return false
	}
	if cols_have_duplication(grid) {
		fmt.Println("Error: there is a duplication in a column")
		return false
	}
	if boxes_have_duplication(grid) {
		fmt.Println("Error: there is a duplication in a box")
		return false
	}
	return true
}

func digit_is_valid(grid []string, chars []byte, cd t_coord, d byte) bool {
	chars[cd.j] = d + '0'
	grid[cd.i] = string(chars)
	return !(line_has_duplication(grid[cd.i]) || col_has_duplication(grid, cd.j) || boxes_have_duplication(grid))
}

func try_digits(grid []string, chars []byte, cd t_coord) bool {
	var d int

	for d = 1; d <= 9; d++ {
		if digit_is_valid(grid, chars, cd, byte(d)) {
			if resolve(grid) {
				return true
			}
		}
	}
	return false
}

func resolve(grid []string) bool {
	var chars []byte
	var cd t_coord

	for cd.i = 0; cd.i < 9; cd.i++ {
		chars = []byte(grid[cd.i])
		for cd.j = 0; cd.j < 9; cd.j++ {
			if grid[cd.i][cd.j] == '.' {
				if !try_digits(grid, chars, cd) {
					chars[cd.j] = '.'
					grid[cd.i] = string(chars)
					return false
				}
			}
		}
	}
	return true
}

func solved(grid []string) bool {
	for _, arg := range grid {
		match, _ := regexp.MatchString("^[1-9]+$", arg)
		if match == false {
			return false
		}
	}
	return true
}

func main() {
	var mode_flag = flag.String("mode", "", "File or piscine input mode")
	var raw_flag = flag.Bool("r", false, "To print raw ouput")
	var create_flag = flag.Bool("c", false, "Generate a grid")
	var diff_flag = flag.String("diff", "normal", "Easy or normal or hard")

	flag.Usage = usage
	flag.Parse()
	if commit_sudoku(*create_flag, *raw_flag, *diff_flag, *mode_flag) {
		return
	}

	var grid []string

	grid = get_grid(*mode_flag)
	if grid == nil {
		return
	}
	if !validate_grid(grid) {
		return
	}
	resolve(grid)
	if solved(grid) {
		print_grid(grid, *raw_flag)
	} else {
		fmt.Println("Error")
	}
}
