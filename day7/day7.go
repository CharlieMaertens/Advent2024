package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func fixEquation(testVal int, cur int, vals []int) float64 {
	if cur == testVal && len(vals) == 0 {
		return float64(testVal)
	}
	if cur > testVal || len(vals) == 0 {
		return 0
	}
	concatenationInt, err := strconv.Atoi(strconv.Itoa(cur) + strconv.Itoa(vals[0]))
	check(err)
	return math.Max(fixEquation(testVal, cur*vals[0], vals[1:]),
		math.Max(fixEquation(testVal, cur+vals[0], vals[1:]),
			fixEquation(testVal, concatenationInt, vals[1:])))

}

func test(line string) int {
	equation := strings.Split(line, ": ")
	testVal, err := strconv.Atoi(equation[0])
	check(err)
	stringVals := strings.Split(equation[1], " ")
	values := []int{}

	for _, strVal := range stringVals {
		value, err := strconv.Atoi(strVal)
		check(err)
		values = append(values, value)
	}
	return int(fixEquation(testVal, 0, values))
}

func main() {
	start := time.Now()
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1] + ".txt"
	}
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += test(scanner.Text())

	}
	fmt.Println(sum)
	fmt.Println(time.Since(start))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
