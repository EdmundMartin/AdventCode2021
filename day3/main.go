package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Popularity struct {
	Zero int
	One int
}

func (p Popularity) String() string {
	return fmt.Sprintf("Zero Count: %d, One Count: %d", p.Zero, p.One)
}

func (p *Popularity) Increment(i string) {
	if i == "0" {
		p.Zero++
	} else {
		p.One++
	}
}


func NewContainer(count int) []*Popularity {
	var container []*Popularity
	for i := 0; i < count; i++ {
		container = append(container, &Popularity{
			Zero: 0,
			One:  0,
		})
	}
	return container
}

func readInputFile(filename string) []string {
	var binaryNumbers []string
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		binaryNumbers = append(binaryNumbers, scanner.Text())
	}
	return binaryNumbers
}

func updatePopularity(larities []*Popularity, binaryNumber string) {
	for idx, val := range strings.Split(binaryNumber, "") {
		larities[idx].Increment(val)
	}
}

func calculate(results []*Popularity) (string, string) {
	gamma := bytes.Buffer{}
	epsilon := bytes.Buffer{}
	for _, pop := range results {
		if pop.Zero > pop.One {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		} else {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		}
	}
	return gamma.String(), epsilon.String()
}

type CommonFilter interface {
	Filter(*Popularity) string
}

type MFilter struct {
}

func (m MFilter) Filter(pop *Popularity) string {
	var majority string
	if pop.One >= pop.Zero {
		majority = "1"
	} else {
		majority = "0"
	}
	return majority
}

type LFilter struct {
}

func (l LFilter) Filter(pop *Popularity) string {
	var minority string
	if pop.One < pop.Zero {
		minority = "1"
	} else {
		minority = "0"
	}
	return minority
}

func filterMajority(values []string, results []*Popularity, filter CommonFilter) []string {
	for idx, pop := range results {
		var newValues []string
		majority := filter.Filter(pop)
		for _, val := range values {
			str := strings.Split(val, "")
			if str[idx] == majority {
				newValues = append(newValues, val)
			}
		}
		values = newValues
		if len(values) == 1 {
			return values
		}
	}
	return values
}

func part1(holder []*Popularity) int64 {
	gamma, epsilon := calculate(holder)
	res, _ := strconv.ParseInt(gamma, 2, 64)
	second, _ := strconv.ParseInt(epsilon, 2, 64)
	return res * second
}

func main() {
	holder := NewContainer(12)
	results := readInputFile("/Users/edmundmartin/go/src/github.com/EdmundMartin/AdventCode2021/day3/test.txt")
	for _, res := range results {
		updatePopularity(holder, res)
	}

	fmt.Printf("Part 1 Answer: %d\n", part1(holder))
	majority := filterMajority(results, holder, MFilter{})
	fmt.Println(majority)
	minority := filterMajority(results, holder, LFilter{})
	fmt.Println(minority)
}