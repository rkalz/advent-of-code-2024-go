package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	acc := 0

	// horizontal
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			var sb strings.Builder
			for k := 0; k < 4; k++ {
				sb.WriteByte(get(lines, i, j+k))
			}
			if sb.String() == "XMAS" || sb.String() == "SAMX" {
				acc++
			}
		}
	}

	// vertical
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			var sb strings.Builder
			for k := 0; k < 4; k++ {
				sb.WriteByte(get(lines, i+k, j))
			}
			if sb.String() == "XMAS" || sb.String() == "SAMX" {
				acc++
			}
		}
	}

	// diagonal \
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			var sb strings.Builder
			for k := 0; k < 4; k++ {
				sb.WriteByte(get(lines, i+k, j+k))
			}
			if sb.String() == "XMAS" || sb.String() == "SAMX" {
				acc++
			}
		}
	}

	// diagonal /
	for i := 0; i < len(lines); i++ {
		for j := len(lines[i]) - 1; j >= 0; j-- {
			var sb strings.Builder
			for k := 0; k < 4; k++ {
				sb.WriteByte(get(lines, i-k, j-k))
			}
			if sb.String() == "XMAS" || sb.String() == "SAMX" {
				acc++
			}
		}
	}

	fmt.Println(acc)
}

func part2(lines []string) {
	acc := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			var left strings.Builder
			var right strings.Builder
			for k := 0; k < 3; k++ {
				left.WriteByte(get(lines, i+k, j+k))
				right.WriteByte(get(lines, i+k, j+2-k))
			}
			if left.String() == "SAM" || left.String() == "MAS" {
				if right.String() == "SAM" || right.String() == "MAS" {
					acc++
				}
			}
		}

	}

	fmt.Println(acc)
}

func get(array []string, i int, j int) byte {
	if i < 0 || i >= len(array) {
		return 0
	}
	if j < 0 || j >= len(array[i]) {
		return 0
	}
	return array[i][j]
}
