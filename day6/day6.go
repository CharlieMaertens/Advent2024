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

type vec2 struct {
	x int
	y int
}

func inbounds(lab []string, pos vec2) bool {
	h := len(lab)
	w := len(lab[0])
	return pos.x >= 0 && pos.x < w && pos.y >= 0 && pos.y < h
}

func getInitPos(lab []string) vec2 {
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

func hitObstacle(lab []string, pos vec2) bool {
	return inbounds(lab, pos) && lab[pos.y][pos.x] == '#'
}

func part1(lab []string) int {
	steps := make(map[string]bool)
	pos := getInitPos(lab)
	dir := vec2{0, -1}
	nextPos := vec2{pos.x + dir.x, pos.y + dir.y}
	for inbounds(lab, pos) {
		nextPos = vec2{pos.x + dir.x, pos.y + dir.y}
		//save distinct steps
		steps["("+string(pos.x)+","+string(pos.y)+")"] = true
		if hitObstacle(lab, nextPos) {
			tmp := dir.x
			dir.x = -dir.y
			dir.y = tmp
		}
		nextPos = vec2{pos.x + dir.x, pos.y + dir.y}

		pos = nextPos
	}

	return len(steps)
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	lab := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lab = append(lab, scanner.Text())
	}

	fmt.Println(part1(lab))
}
