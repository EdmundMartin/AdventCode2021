package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func part1(values []int) int {
	prev := math.MaxInt32
	count := 0
	for _, value := range values {
		if value > prev {
			count++
		}
		prev = value
	}
	return count
}

func part2(values []int) int {
	count := 0
	for i := 3; i < len(values); i++ {
		prev := values[i-3]
		current := values[i]
		if current > prev {
			count += 1
		}
	}
	return count
}

func readAsSlice(fileName string) []int {
	f, _ := os.Open(fileName)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var values []int
	for scanner.Scan() {
		result, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("all values should be ints")
		}
		values = append(values, result)
	}
	return values
}

func main() {
	values := readAsSlice("/Users/edmundmartin/go/src/github.com/EdmundMartin/AdventCode2021/day1/input.txt")

	result := part1(values)
	fmt.Printf("Part 1 result: %d\n", result)
	result = part2(values)
	fmt.Printf("Part 2 result: %d\n", result)
}
