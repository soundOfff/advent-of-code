package main

import (
	readinput "advent-of-code/day_2/read_input"
	"advent-of-code/day_2/rules"
	"flag"
	"fmt"
)

var (
	inputFileName = flag.String("input-file", "input.txt", "this will be used as input")
	dampener      = flag.Bool("dampener", false, "do we have Dampener")
)

func main() {
	flag.Parse()
	inputData := readinput.ReadFile(*inputFileName)

	fmt.Println("first: ", stepOne(&inputData))
	fmt.Println("second: ", stepTwo(&inputData))
}

func stepOne(inputData *[][]int) int {
	c := 0
	for _, data := range *inputData {
		if validate(data) {
			c++
		}
	}

	return c
}

func stepTwo(inputData *[][]int) int {
	c := 0
	for _, data := range *inputData {
		if validate(data) || validateByRemoving(data) {
			c++
		}
	}

	return c
}

func validateByRemoving(data []int) bool {
	if !*dampener {
		return false
	}
	for i := 0; i < len(data); i++ {
		data1 := append([]int{}, data[:i]...)
		data1 = append(data1, data[i+1:]...)
		if validate(data1) {
			return true
		}
	}
	return false
}

func validate(data []int) bool {
	desc := rules.Decreasing(data)
	incr := rules.Increasing(data)
	max := rules.MaxDiff(data, 3)
	min := rules.MinDiff(data, 1)
	return (desc || incr) && max && min
}
