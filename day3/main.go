package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	line := ""
	for scanner.Scan() {
		line += scanner.Text()
	}

	part1(line)
	part2(line)
}

func part1(line string) {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(line, -1)

	acc := 0
	for _, match := range matches {
		acc += mulToNum(match)
	}

	fmt.Println(acc)
}

func part2(line string) {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(line, -1)

	acc := 0
	enabled := true
	for _, match := range matches {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else if enabled {
			acc += mulToNum(match)
		}
	}

	fmt.Println(acc)
}

func mulToNum(match string) int {
	split := strings.Split(strings.TrimPrefix(strings.TrimSuffix(match, ")"), "mul("), ",")
	left, _ := strconv.Atoi(split[0])
	right, _ := strconv.Atoi(split[1])

	return left * right
}
