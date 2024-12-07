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

func fix(rules map[string][]string, update []string) []string {
	for i := 0; i < len(update); i++ {
		for j := 0; j < len(update); j++ {

			if slices.Contains(rules[update[i]], update[j]) {
				temp := update[i]
				update[i] = update[j]
				update[j] = temp
			}
		}

	}

	return update
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	beforeMap := make(map[string][]string)
	afterMap := make(map[string][]string)
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

		_, in = afterMap[before]
		if !in {
			afterMap[before] = []string{}
		}

		afterMap[after] = append(afterMap[before], after)
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
			betterUpdate := fix(beforeMap, update)
			part2 += middleValue(betterUpdate)
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
