package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func isXmas(word string) int {
	if word == "MAS" {
		return 1
	} else {
		return 0
	}
}

func xmasChecks(x int, y int, cw []string) int {
	numXMAS := 0
	xPlus := x+3 < len(cw[0])
	yPlus := y+3 < len(cw)
	xMinus := x-3 >= 0
	yMinus := y-3 >= 0

	if xPlus {
		numXMAS += isXmas(string(cw[x+1][y]) + string(cw[x+2][y]) + string(cw[x+3][y]))
	}
	if yPlus {
		numXMAS += isXmas(string(cw[x][y+1]) + string(cw[x][y+2]) + string(cw[x][y+3]))
	}
	if xMinus {
		numXMAS += isXmas(string(cw[x-1][y]) + string(cw[x-2][y]) + string(cw[x-3][y]))
	}
	if yMinus {
		numXMAS += isXmas(string(cw[x][y-1]) + string(cw[x][y-2]) + string(cw[x][y-3]))
	}
	if xPlus && yPlus {
		numXMAS += isXmas(string(cw[x+1][y+1]) + string(cw[x+2][y+2]) + string(cw[x+3][y+3]))
	}
	if xPlus && yMinus {
		numXMAS += isXmas(string(cw[x+1][y-1]) + string(cw[x+2][y-2]) + string(cw[x+3][y-3]))
	}
	if xMinus && yPlus {
		numXMAS += isXmas(string(cw[x-1][y+1]) + string(cw[x-2][y+2]) + string(cw[x-3][y+3]))
	}
	if xMinus && yMinus {
		numXMAS += isXmas(string(cw[x-1][y-1]) + string(cw[x-2][y-2]) + string(cw[x-3][y-3]))
	}
	return numXMAS

}

func part1(cw []string) int {
	sum := 0
	for i := 0; i < len(cw); i++ {
		for j := 0; j < len(cw[0]); j++ {
			if string(cw[j][i]) == "X" {
				sum += xmasChecks(j, i, cw)
			}
		}
	}
	return sum
}

func isMasX(word string) bool {
	if word == "MS" || word == "SM" {
		return true
	} else {
		return false
	}
}

func masXCheck(x int, y int, cw []string) int {
	xPlus := x+1 < len(cw)
	yPlus := y+1 < len(cw[0])
	xMinus := x-1 >= 0
	yMinus := y-1 >= 0

	if xPlus && xMinus && yPlus && yMinus {
		if isMasX(string(cw[x+1][y+1])+string(cw[x-1][y-1])) &&
			isMasX(string(cw[x+1][y-1])+string(cw[x-1][y+1])) {
			return 1
		}
	}
	return 0

}

func part2(cw []string) int {
	sum := 0
	for i := 0; i < len(cw); i++ {
		for j := 0; j < len(cw[0]); j++ {
			if string(cw[i][j]) == "A" {
				sum += masXCheck(i, j, cw)
			}
		}
	}
	return sum
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	cw := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cw = append(cw, scanner.Text())
	}

	fmt.Println(part1(cw))
	fmt.Println(part2(cw))
}
