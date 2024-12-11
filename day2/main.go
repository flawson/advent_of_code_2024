package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	Unset = iota
	Increasing
	Decreasing
)

func checkLine(e []int) bool {
	state := Unset

	for i := 0; i < len(e)-1; i++ {
		prevState := state

		diff := e[i] - e[i+1]
		switch {
		case diff > 0:
			state = Increasing
			if diff > 3 {
				return false
			}
		case diff < 0:
			state = Decreasing
			if diff < -3 {
				return false
			}
		default:
			return false
		}

		if prevState != Unset && prevState != state {
			return false
		}
	}

	return true
}

func evalWorker(in chan []int, res chan bool, dampen bool) {
	for e := range in {
		valid := false

		ok := checkLine(e)
		if ok {
			valid = true
		} else {
			fmt.Printf("Truely Undampened: %v\n", e)
			if dampen {
				for i := 0; i < len(e); i++ {
					newE := slices.Concat(e[:i], e[i+1:])
					fmt.Printf("Dampened %v\n", newE)
					//fmt.Printf("Dampened entry: %v\n", newE)
					newOk := checkLine(newE)

					if newOk {
						valid = true
						break
					}
				}
			}
		}
		res <- valid
	}
}

func readInput(filename string) [][]int {
	inputList := make([][]int, 0)
	data, _ := os.ReadFile(filename)

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		elemList := make([]int, 0)
		l := strings.Split(line, " ")
		for _, v := range l {
			i, _ := strconv.Atoi(v)
			elemList = append(elemList, i)
		}
		inputList = append(inputList, elemList)
	}

	return inputList
}

func part1(inputList [][]int) {
	inputs := make(chan []int)
	entryRes := make(chan bool)
	validEntries := 0
	totalEntries := 0

	go evalWorker(inputs, entryRes, false)

	go func() {
		for _, l := range inputList {
			inputs <- l
		}
	}()

	for i := 0; i < len(inputList); i++ {
		v := <-entryRes
		totalEntries++
		if v {
			validEntries++
		}
	}

	fmt.Printf("Valid entries: %d\n", validEntries)
	fmt.Printf("Total entries: %d\n", totalEntries)
}

func part2(inputList [][]int) {
	inputs := make(chan []int)
	entryRes := make(chan bool)
	validEntries := 0
	totalEntries := 0

	go evalWorker(inputs, entryRes, true)

	go func() {
		for _, l := range inputList {
			inputs <- l
		}
	}()

	for i := 0; i < len(inputList); i++ {
		v := <-entryRes

		if v {
			validEntries++
		}
		totalEntries++
	}

	fmt.Printf("Valid entries: %d\n", validEntries)
	fmt.Printf("Total entries: %d\n", totalEntries)
}

func main() {
	inputList := readInput("input.txt")
	part1(inputList)
	part2(inputList)
}
