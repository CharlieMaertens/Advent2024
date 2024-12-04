package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func report(strReport []string) []int {
	report := []int{}
	for _, strLevel := range strReport {
		level, err := strconv.Atoi(strLevel)
		check(err)
		report = append(report, level)
	}
	return report
}

func getDir(a int, b int) int {
	if a == b {
		return 0
	} else if a < b {
		return 1
	} else {
		return -1
	}
}

func reportSafe(report []int) int {
	dir := getDir(report[0], report[1])
	if dir == 0 {
		return 0
	}

	for i := 0; i < len(report)-1; i++ {
		level := report[i]
		adjacentLevel := report[i+1]

		if getDir(level, adjacentLevel) != dir {
			return 0
		}

		diff := math.Abs(float64(level) - float64(adjacentLevel))
		if diff < 1 || diff > 3 {
			return 0
		}
	}
	return 1
}

func safeReports(file *os.File) int {
	a := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strReport := strings.Split(line, " ")
		report := report(strReport)
		a += reportSafe(report)
	}

	check(scanner.Err())
	return a
}

func solvePart1() int {
	file, err := os.Open("test-input.txt")
	check(err)
	defer file.Close()
	return safeReports(file)
}

func reportSafe2(strReport []string) int {
	report := report(strReport)
	if reportSafe(report) == 1 {
		return 1
	}
	for i := 0; i < len(report); i++ {
		if reportSafe(slices.Concat(report[:i], report[i+1:])) == 1 {
			return 1
		}
	}
	return 0
}

func safeReports2(file *os.File) int {
	a := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Split(line, " ")
		a += reportSafe2(report)
	}

	check(scanner.Err())
	return a
}

func solvePart2() int {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	return safeReports2(file)
}

func main() {
	fmt.Println(solvePart1())
	fmt.Println(solvePart2())
}
