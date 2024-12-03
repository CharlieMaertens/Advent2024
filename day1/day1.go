package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildLists() ([]int, []int) {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	l1 := []int{}
	l2 := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "   ")
		num1, err := strconv.Atoi(nums[0])
		check(err)
		num2, err := strconv.Atoi(nums[1])
		check(err)

		l1 = append(l1, num1)
		l2 = append(l2, num2)
	}

	check(scanner.Err())
	return l1, l2
}

func solvePart1() int {
	l1, l2 := buildLists()

	sort.Ints(l1)
	sort.Ints(l2)

	a := 0.0
	for i := 0; i < len(l1); i++ {
		a += math.Abs(float64(l1[i] - l2[i]))
	}

	return int(a)
}

func occurencesMap(nums []int) map[int]int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	return m
}

func solvePart2() int {
	l1, l2 := buildLists()

	om := occurencesMap(l2)

	a := 0
	for _, num := range l1 {
		a += (num * om[num])
	}

	return a
}

func main() {
	a := solvePart1()
	fmt.Printf("%d\n", a)

	a = solvePart2()
	fmt.Printf("%d\n", a)
}
