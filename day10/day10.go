package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumTrailsP1(m []string, x int, y int, endsFound map[vec2]bool) int {
	sum := 0
	current, err := strconv.Atoi(string(m[y][x]))
	check(err)

	if current == 9 && !endsFound[vec2{x, y}] {
		endsFound[vec2{x, y}] = true
		return 1
	}
	if x+1 < len(m[0]) {
		right, err := strconv.Atoi(string(m[y][x+1]))
		check(err)
		if right-current == 1 {
			sum += sumTrailsP1(m, x+1, y, endsFound)
		}
	}
	if x-1 >= 0 {
		left, err := strconv.Atoi(string(m[y][x-1]))
		check(err)
		if left-current == 1 {
			sum += sumTrailsP1(m, x-1, y, endsFound)
		}
	}
	if y+1 < len(m) {
		up, err := strconv.Atoi(string(m[y+1][x]))
		check(err)
		if up-current == 1 {
			sum += sumTrailsP1(m, x, y+1, endsFound)
		}
	}
	if y-1 >= 0 {
		down, err := strconv.Atoi(string(m[y-1][x]))
		check(err)
		if down-current == 1 {
			sum += sumTrailsP1(m, x, y-1, endsFound)
		}
	}

	return sum
}

func sumTrailsP2(m []string, x int, y int) int {
	sum := 0
	current, err := strconv.Atoi(string(m[y][x]))
	check(err)

	if current == 9 {
		return 1
	}
	if x+1 < len(m[0]) {
		right, err := strconv.Atoi(string(m[y][x+1]))
		check(err)
		if right-current == 1 {
			sum += sumTrailsP2(m, x+1, y)
		}
	}
	if x-1 >= 0 {
		left, err := strconv.Atoi(string(m[y][x-1]))
		check(err)
		if left-current == 1 {
			sum += sumTrailsP2(m, x-1, y)
		}
	}
	if y+1 < len(m) {
		up, err := strconv.Atoi(string(m[y+1][x]))
		check(err)
		if up-current == 1 {
			sum += sumTrailsP2(m, x, y+1)
		}
	}
	if y-1 >= 0 {
		down, err := strconv.Atoi(string(m[y-1][x]))
		check(err)
		if down-current == 1 {
			sum += sumTrailsP2(m, x, y-1)
		}
	}

	return sum
}

type vec2 struct {
	x int
	y int
}

func part1(m []string) int {
	sum := 0
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if m[y][x] == '0' {
				endsFound := make(map[vec2]bool)
				sum += sumTrailsP1(m, x, y, endsFound)
			}
		}
	}
	return sum
}

func part2(m []string) int {
	sum := 0
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if m[y][x] == '0' {
				sum += sumTrailsP2(m, x, y)
			}
		}
	}
	return sum
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1] + ".txt"
	}
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	tMap := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tMap = append(tMap, scanner.Text())
	}

	fmt.Println(part1(tMap))
	fmt.Println(part2(tMap))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
