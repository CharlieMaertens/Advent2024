package main

import (
	"bufio"
	"fmt"
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

func middleValue(update []string) int {
	value, err := strconv.Atoi(update[len(update)/2])
	check(err)
	return value
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	beforeMap := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() && scanner.Text() != "" {
		rule := strings.Split(scanner.Text(), "|")
		before := rule[0]
		after := rule[1]

		_, in := beforeMap[after]
		if !in {
			beforeMap[after] = []string{}
		}

		beforeMap[after] = append(beforeMap[after], before)
	}
	part1 := 0
	part2 := 0
	correct := true
	for scanner.Scan() {
		update := strings.Split(scanner.Text(), ",")
		correct = true
		for i := 0; i < len(update) && correct; i++ {
			for j := i + 1; j < len(update) && correct; j++ {
				correct = !slices.Contains(beforeMap[update[i]], update[j])
			}
		}
		if correct {
			part1 += middleValue(update)
		} else {
			slices.SortFunc(update, func(a string, b string) int {
				if slices.Contains(beforeMap[a], b) {
					return -1
				}
				return 1
			})
			part2 += middleValue(update)
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
