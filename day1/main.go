package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "   ")

		l, _ := strconv.Atoi(nums[0])
		left = append(left, l)

		r, _ := strconv.Atoi(nums[1])
		right = append(right, r)
	}

	part1(left, right)
	part2(left, right)
}

func part1(left []int, right []int) {
	slices.Sort(left)
	slices.Sort(right)

	acc := 0

	for i := range left {
		diff := int(math.Abs(float64(left[i] - right[i])))
		acc += diff
	}

	fmt.Println(acc)
}

func part2(left []int, right []int) {
	dict := make(map[int]int)

	for _, n := range right {
		dict[n] += 1
	}

	acc := 0
	for _, n := range left {
		acc += n * dict[n]
	}

	fmt.Println(acc)
}
