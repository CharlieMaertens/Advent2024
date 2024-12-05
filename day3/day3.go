package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func digit(a string) bool {
	_, err := strconv.Atoi(a)
	return err == nil

}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = "test-" + filename
	}
	data, err := os.ReadFile(filename)
	check(err)
	sum := 0
	lenDo := len("do()")
	lenDont := len("don't()")
	do := true
	l := len(data)
	for j := 0; j < l; j++ {
		var x string
		var y string

		yOff := 0

		if j+lenDo < l && string(data[j:j+lenDo]) == "do()" {
			do = true
		} else if j+lenDont < l && string(data[j:j+lenDont]) == "don't()" {
			do = false
		} else if j+4 < l && string(data[j:j+4]) == "mul(" {
			if j+5 < l && string(data[j+5]) == "," {
				yOff = 6
				x = string(data[j+4])
			} else if j+6 < l && string(data[j+6]) == "," {
				x = string(data[j+4 : j+6])
				yOff = 7
			} else if j+7 < l && string(data[j+7]) == "," {
				yOff = 8
				x = string(data[j+4 : j+7])
			}

			if j+yOff+1 < l && string(data[j+yOff+1]) == ")" {
				y = string(data[j+yOff])
			} else if j+yOff+2 < l && string(data[j+yOff+2]) == ")" {
				y = string(data[j+yOff : j+yOff+2])
			} else if j+yOff+3 < l && string(data[j+yOff+3]) == ")" {
				y = string(data[j+yOff : j+yOff+3])
			}

			if digit(x) && digit(y) {
				xInt, err := strconv.Atoi(x)
				check(err)
				yInt, err := strconv.Atoi(y)
				check(err)
				if do {
					sum += (xInt * yInt)
				}

			}
		}
	}
	fmt.Println(sum)
}
