package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	rows := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		row := []int{}
		for _, entry := range split {
			n, _ := strconv.Atoi(entry)
			row = append(row, n)
		}
		rows = append(rows, row)
	}

	part1(rows)
	part2(rows)
}

func part1(rows [][]int) {
	acc := 0

	for _, row := range rows {
		if legal(row) {
			acc += 1
		}
	}

	fmt.Println(acc)
}

func part2(rows [][]int) {
	acc := 0

	for _, row := range rows {
		valid := legal(row)

		if valid {
			acc++
			continue
		}

		for j := 0; j < len(row); j++ {
			slice := []int{}
			slice = append(slice, row[:j]...)
			slice = append(slice, row[j+1:]...)

			if legal(slice) {
				acc++
				break
			}
		}
	}

	fmt.Println(acc)
}

func legal(row []int) bool {
	incr := row[0] < row[1]
	for i := 1; i < len(row); i++ {
		a := row[i-1]
		b := row[i]

		if (incr && a > b) || (!incr && a < b) {
			return false
		}

		diff := int(math.Abs(float64(a - b)))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
