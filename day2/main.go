package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Position struct {
	Depth int
	Horizontal int
	Aim int
}

type Movement struct {
	Horizontal int
	Depth int
	MoveType string
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


func readInputFile(filename string) []Movement {
	var movements []Movement
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result := scanner.Text()
		values := strings.Split(result, " ")
		direction := values[0]
		distance, _ := strconv.Atoi(values[1])
		switch direction {
		case "forward":
			movements = append(movements, Movement{
				Horizontal: distance,
				Depth:      0,
				MoveType: "forward",
			})
		case "down":
			movements = append(movements, Movement{
				Horizontal: 0,
				Depth: distance,
				MoveType: "down",
			})
		case "up":
			movements = append(movements, Movement{
				Horizontal: 0,
				Depth:      -distance,
				MoveType: "up",
			})
		}
	}
	return movements
}

func part1(changes []Movement) int {
	pos := Position{Depth: 0, Horizontal: 0, Aim: 0}
	for _, result := range changes {
		pos.Horizontal += result.Horizontal
		pos.Depth += result.Depth
	}
	return pos.Horizontal * pos.Depth
}

func part2(changes []Movement) int {
	pos := Position{
		Depth:      0,
		Horizontal: 0,
		Aim:        0,
	}
	for _, result := range changes {
		if result.MoveType == "down" {
			pos.Aim += abs(result.Depth)
		}
		if result.MoveType == "up" {
			pos.Aim -= abs(result.Depth)
		}
		if result.MoveType == "forward" {
			pos.Horizontal += result.Horizontal
			pos.Depth += result.Horizontal * pos.Aim
		}
	}
	return pos.Horizontal * pos.Depth
}

func main() {

	results := readInputFile("/Users/edmundmartin/go/src/github.com/EdmundMartin/AdventCode2021/day2/inputs.txt")

	fmt.Printf("Part 1 result: %d\n", part1(results))
	fmt.Printf("Part 2 result: %d\n", part2(results))
}