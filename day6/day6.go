package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type vec2 struct {
	x int
	y int
}

type dVec struct {
	pos vec2
	dir vec2
}

func inbounds(lab [][]rune, pos vec2) bool {
	h := len(lab)
	w := len(lab[0])
	return pos.x >= 0 && pos.x < w && pos.y >= 0 && pos.y < h
}

func getInitPos(lab [][]rune) vec2 {
	pos := vec2{0, 0}
	for y := 0; y < len(lab); y++ {
		for x := 0; x < len(lab[0]); x++ {
			if lab[y][x] == '^' {
				pos = vec2{x, y}
			}
		}
	}
	return pos
}

func hitObstacle(lab [][]rune, pos vec2) bool {
	return inbounds(lab, pos) && lab[pos.y][pos.x] == '#'
}

func printLab(lab [][]rune) {
	d := ""
	for _, line := range lab {
		d += string(line) + "\n"
	}
	a := strconv.Itoa(rand.Intn(100))
	err := os.WriteFile("./tmp/test-"+a+".txt", []byte(d), 0644)
	check(err)
}

func part2(steps map[vec2]bool) int {
	stepLabs := getLab()
	loops := 0
	for key := range steps {
		if stepLabs[key.y][key.x] == '^' {
			continue
		}
		lab := getLab()
		lab[key.y][key.x] = '#'
		stepLabs[key.y][key.x] = 'X'
		loopy := isLoop(lab)
		if loopy == 1 {
			stepLabs[key.y][key.x] = 'O'
			loops++
		}
	}

	return loops
}

func isLoop(lab [][]rune) int {
	steps := make(map[dVec]bool)
	stepLabs := getLab()
	pos := getInitPos(lab)
	dir := vec2{0, -1}
	var nextPos vec2
	for inbounds(lab, pos) {
		nextPos = vec2{pos.x + dir.x, pos.y + dir.y}
		stepLabs[pos.y][pos.x] = 'X'
		_, in := steps[dVec{pos, dir}]
		if in {
			return 1
		}
		if hitObstacle(lab, nextPos) {
			tmp := dir.x
			dir.x = -dir.y
			dir.y = tmp
		} else {
			steps[dVec{pos, dir}] = true
			pos = nextPos
		}
	}

	return 0
}

func part1(lab [][]rune) map[vec2]bool {
	stepLabs := getLab()
	steps := make(map[vec2]bool)
	pos := getInitPos(lab)
	dir := vec2{0, -1}
	var nextPos vec2
	for inbounds(lab, pos) {
		nextPos = vec2{pos.x + dir.x, pos.y + dir.y}
		stepLabs[pos.y][pos.x] = 'X'

		if hitObstacle(lab, nextPos) {
			tmp := dir.x
			dir.x = -dir.y
			dir.y = tmp
		} else {
			steps[pos] = true
			pos = nextPos
		}
	}
	printLab(stepLabs)

	return steps
}

func getLab() [][]rune {
	filename := "input.txt"
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	lab := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		lab = append(lab, line)
	}
	return lab
}

func main() {
	start := time.Now()
	lab := getLab()

	steps := part1(lab)
	loops := part2(steps)

	fmt.Println(len(steps))
	fmt.Println(loops)
	fmt.Println(time.Since(start))
}
