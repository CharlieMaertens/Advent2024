package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func part1(diskMap string) int {
	start := time.Now()
	blocks := []int{}
	index := 0
	for i := 0; i < len(diskMap)-1; i += 2 {
		fileLen, err := strconv.Atoi(string(diskMap[i]))
		check(err)
		freeLen, err := strconv.Atoi(string(diskMap[i+1]))
		check(err)

		for range fileLen {
			blocks = append(blocks, index)
		}
		for range freeLen {
			blocks = append(blocks, -1)
		}
		index++
	}

	if len(diskMap)%2 != 0 {
		fileLen, err := strconv.Atoi(string(diskMap[len(diskMap)-1]))
		check(err)
		for range fileLen {
			blocks = append(blocks, index)
		}
	}

	backIndex := len(blocks) - 1
	for i := 0; i <= backIndex; i++ {
		for blocks[backIndex] == -1 {
			backIndex--
		}
		if blocks[i] == -1 {
			blocks[i] = blocks[backIndex]
			blocks[backIndex] = -1
			backIndex--
		}
	}

	checkSum := 0
	for i := 0; i <= backIndex; i++ {
		checkSum += (i * blocks[i])
	}

	fmt.Println(time.Since(start))
	return checkSum
}

type block = struct {
	length     int
	value      int
	startIndex int
}

func part2(diskMap string) int {
	start := time.Now()
	blocks := []block{}
	index := 0
	final := []int{}

	for i := 0; i < len(diskMap)-1; i += 2 {
		fileLen, err := strconv.Atoi(string(diskMap[i]))
		check(err)
		freeLen, err := strconv.Atoi(string(diskMap[i+1]))
		check(err)
		if fileLen > 0 {
			blocks = append(blocks, block{length: fileLen, value: index, startIndex: len(final)})
		}

		for range fileLen {
			final = append(final, index)
		}

		if freeLen > 0 {
			blocks = append(blocks, block{length: freeLen, value: -1, startIndex: len(final)})
		}

		for range freeLen {
			final = append(final, -1)
		}
		index++
	}

	if len(diskMap)%2 != 0 {
		fileLen, err := strconv.Atoi(string(diskMap[len(diskMap)-1]))
		check(err)
		if fileLen > 0 {
			blocks = append(blocks, block{length: fileLen, value: index, startIndex: len(final)})
		}
		for range fileLen {
			final = append(final, index)
		}
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i].value == -1 {
			continue
		}
		for j := 0; j < len(blocks); j++ {
			if blocks[i].value == blocks[j].value {
				break
			}

			if blocks[j].value != -1 {
				continue
			}

			if blocks[i].length > blocks[j].length {
				continue
			}

			for range blocks[i].length {
				final[blocks[j].startIndex] = blocks[i].value
				final[blocks[i].startIndex] = -1

				blocks[j].startIndex++
				blocks[i].startIndex++
			}
			blocks[j].length -= blocks[i].length
			break
		}

	}

	checkSum := 0
	for i := 0; i < len(final); i++ {
		if final[i] == -1 {
			continue
		}
		checkSum += (i * final[i])
	}

	fmt.Println(time.Since(start))
	return checkSum
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1] + ".txt"
	}
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	diskMap := ""
	for scanner.Scan() {
		diskMap = scanner.Text()
	}
	fmt.Println(part1(diskMap))
	fmt.Println(part2(diskMap))

}

func printBlocks(blocks []int) {
	for _, value := range blocks {
		if value == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(value)
		}
	}
	fmt.Println()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
