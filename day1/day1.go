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
	myfile, err := os.Open("input.txt")
	check(err)
	defer myfile.Close()

	lst1 := []int{}
	lst2 := []int{}

	scanner := bufio.NewScanner(myfile)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		num1, err := strconv.Atoi(numbers[0])
		check(err)
		num2, err := strconv.Atoi(numbers[1])
		check(err)

		lst1 = append(lst1, num1)
		lst2 = append(lst2, num2)
	}

	check(scanner.Err())
	return lst1, lst2
}

func solve() float64 {
	lst1, lst2 := buildLists()

	sort.Ints(lst1)
	sort.Ints(lst2)

	answer := 0.0
	for i := 0; i < len(lst1); i++ {
		answer += math.Abs(float64(lst1[i] - lst2[i]))
	}

	return answer
}

func main() {
	answer := solve()
	fmt.Printf("%f\n", answer)
}
