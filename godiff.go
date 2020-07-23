package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readLines(f string) []string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func longest(x []string, y []string) []string {
	if len(x) > len(y) {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lcs(x []string, y []string) []int {
	tbl := make([]int, (len(x)+1)*(len(y)+1))

	var _lcs func(i, j int) int
	_lcs = func(i, j int) int {
		if i == 0 || j == 0 {
			return 0
		}
		s := tbl[j*len(x)+i]
		if s != 0 {
			return s
		}

		if x[i-1] == y[j-1] {
			s = _lcs(i-1, j-1) + 1
		} else {
			s = max(_lcs(i, j-1), _lcs(i-1, j))
		}
		tbl[j*len(x)+i] = s
		return s

	}
	_lcs(len(x), len(y))
	return tbl
}

func printDiff(tbl []int, x []string, y []string) {
	var _printDiff func(i, j int)
	_printDiff = func(i, j int) {
		switch {
		case i > 0 && j > 0 && x[i-1] == y[j-1]:
			_printDiff(i-1, j-1)
			fmt.Println(" ", x[i-1])
		case j > 0 && (i == 0 || tbl[(j-1)*len(x)+i] >= tbl[j*len(x)+i-1]):
			_printDiff(i, j-1)
			fmt.Println("+", y[j-1])
		case i > 0 && (j == 0 || tbl[(j-1)*len(x)+i] < tbl[j*len(x)+i-1]):
			_printDiff(i-1, j)
			fmt.Println("-", x[i-1])
		}
	}
	_printDiff(len(x), len(y))
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("wrong # of arguments")
	}
	x := readLines(os.Args[1])
	y := readLines(os.Args[2])
	tbl := lcs(x, y)
	printDiff(tbl, x, y)
}
