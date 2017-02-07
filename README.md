# commit_sudoku - a sudoku generator and solver in Go
###### Wrote this to learn Go.
###### The solving part is originally a school (pool) project that has to be written in C.

#### Build
* `go build sudoku.go`


#### Usage 
* To resolve a grid:
	* `sudoku [-r] [-mode=file|piscine] [input]`
	* Examples:
		* `sudoku.go sample`
		* `sudoku.go -r sample`
		* `sudoku.go -mode file sample`
		* `sudoku.go -mode piscine "9...7...." "2...9..53" ".6..124.." "84...1.9." "5.....8.." ".31..4..." "..37..68." ".9..5.741" "47......."`
		* `sudoku -r -c | sudoku`


* To generate a grid:
 	* `sudoku -c [-r] [-diff=easy|normal|hard]`
 	* Examples:
 		* `sudoku.go -c`
 		* `sudoku.go -c -r`
 		* `sudoku.go -c -r -diff easy`
   		* `sudoku.go -c -diff hard`

