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
	for i := 1; i <= len(x); i++ {
		for j := 1; j <= len(y); j++ {
			if x[i-1] == y[j-1] {
				tbl[j*len(x)+i] = tbl[(j-1)*len(x)+i-1] + 1
			} else {
				tbl[j*len(x)+i] = max(tbl[(j-1)*len(x)+i], tbl[j*len(x)+i-1])
			}
		}
	}
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
