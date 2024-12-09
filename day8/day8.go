package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type vec2 struct {
	x int
	y int
}

func part1() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1] + ".txt"
	}
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	grid := [][]rune{}
	antennas := make(map[rune][]vec2)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '.' {
				continue
			}

			freq := grid[i][j]
			_, in := antennas[freq]
			if !in {
				antennas[freq] = []vec2{}
			}

			antennas[freq] = append(antennas[freq], vec2{i, j})
		}
	}

	antinodes := make(map[vec2]bool)
	for _, locations := range antennas {
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				// fmt.Println(i, j)
				xOff := locations[i].x - locations[j].x
				yOff := locations[i].y - locations[j].y
				new1 := vec2{locations[i].x + xOff, locations[i].y + yOff}
				new2 := vec2{locations[j].x - xOff, locations[j].y - yOff}
				if new1.x >= 0 && new1.x < len(grid) && new1.y >= 0 && new1.y < len(grid[0]) {
					antinodes[new1] = true
				}
				if new2.x >= 0 && new2.x < len(grid) && new2.y >= 0 && new2.y < len(grid[0]) {
					antinodes[new2] = true
				}
			}
		}
	}

	for antinode := range antinodes {
		// if grid[antinode.x][antinode.y] == '.' {
		grid[antinode.x][antinode.y] = '#'
		// }
	}
	// printGrid(grid)

	fmt.Println(len(antinodes))
}

func part2() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1] + ".txt"
	}
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	grid := [][]rune{}
	antennas := make(map[rune][]vec2)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '.' {
				continue
			}

			freq := grid[i][j]
			_, in := antennas[freq]
			if !in {
				antennas[freq] = []vec2{}
			}

			antennas[freq] = append(antennas[freq], vec2{i, j})
		}
	}

	antinodes := make(map[vec2]bool)
	for _, locations := range antennas {
		for i := 0; i < len(locations); i++ {
			for j := 0; j < len(locations); j++ {
				if i == j {
					continue
				}
				// fmt.Println(i, j)
				xOff := locations[i].x - locations[j].x
				yOff := locations[i].y - locations[j].y
				new1 := vec2{locations[i].x + xOff, locations[i].y + yOff}
				new2 := vec2{locations[i].x - xOff, locations[i].y - yOff}
				for new1.x >= 0 && new1.x < len(grid) && new1.y >= 0 && new1.y < len(grid[0]) {
					antinodes[new1] = true
					new1 = vec2{new1.x + xOff, new1.y + yOff}
				}
				for new2.x >= 0 && new2.x < len(grid) && new2.y >= 0 && new2.y < len(grid[0]) {
					antinodes[new2] = true
					new2 = vec2{new2.x - xOff, new2.y - yOff}
				}
			}
		}
	}

	for antinode := range antinodes {
		if grid[antinode.x][antinode.y] == '.' {
			grid[antinode.x][antinode.y] = '#'
		}
	}
	printGrid(grid)

	fmt.Println(len(antinodes))
}

func main() {
	part1()
	part2()
}

func printGrid(grid [][]rune) {
	d := ""
	for _, line := range grid {
		d += string(line) + "\n"
	}
	a := strconv.Itoa(rand.Intn(100))
	err := os.WriteFile("./tmp/test-"+a+".txt", []byte(d), 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
