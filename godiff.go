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

var status = 0

func printRange(prefix string, start, d int, lines []string) {
	status = 1
	for i := 0; i < d; i++ {
		fmt.Printf("%s%s\n", prefix, lines[start+i])
	}
}

func printCmd(c string, i0, i1, j0, j1 int) {
	switch {
	case i0 == i1 && j0 == j1:
		fmt.Printf("%d%s%d\n", i0, c, j0)
	case i0 == i1:
		fmt.Printf("%d%s%d,%d\n", i0, c, j0, j1)
	case j0 == j1:
		fmt.Printf("%d,%d%s%d\n", i0, i1, c, j0)
	default:
		fmt.Printf("%d,%d%s%d,%d\n", i0, i1, c, j0, j1)
	}
}

func printDiff(tbl []int, x []string, y []string) {
	printChange := func(i, j, di, dj int) {
		switch {
		case di == 0 && dj == 0:
		case di == 0:
			printCmd("a", i, i, j+1, j+dj)
			printRange("> ", j, dj, y)
		case dj == 0:
			printCmd("d", i+1, i+di, j, j)
			printRange("< ", i, di, x)
		default:
			printCmd("c", i+1, i+di, j+1, j+dj)
			printRange("< ", i, di, x)
			fmt.Println("---")
			printRange("> ", j, dj, y)
		}
	}

	i, j := len(x), len(y)
	var di, dj int
	for {
		switch {
		case i > 0 && j > 0 && x[i-1] == y[j-1]:
			defer printChange(i, j, di, dj)
			di, dj = 0, 0
			i--
			j--
		case j > 0 && (i == 0 || tbl[(j-1)*len(x)+i] > tbl[j*len(x)+i-1]):
			dj++
			j--
		case i > 0 && (j == 0 || tbl[(j-1)*len(x)+i] <= tbl[j*len(x)+i-1]):
			di++
			i--
		default:
			printChange(i, j, di, dj)
			return
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("wrong # of arguments")
	}
	x := readLines(os.Args[1])
	y := readLines(os.Args[2])
	tbl := lcs(x, y)
	printDiff(tbl, x, y)
	os.Exit(status)
}
