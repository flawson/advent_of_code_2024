package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Unset = iota
	Increasing
	Decreasing
)

func parseWorker(in chan string, out chan []int) {
	for l := range in {
		vals := strings.Split(l, " ")
		valNums := make([]int, len(vals))

		for i := 0; i < len(vals); i++ {
			n, _ := strconv.Atoi(vals[i])
			valNums[i] = n
		}

		out <- valNums
	}
	close(out)
}

func evalWorker(in chan []int, res chan bool) {
	for e := range in {
		state := Unset
		valid := true
		dampened := false

		for i := 0; i < len(e)-1; i++ {
			prevState := state

			diff := e[i] - e[i+1]
			switch {
			case diff > 0:
				state = Increasing
				if diff > 3 {
					valid = false
				}
			case diff < 0:
				state = Decreasing
				if diff < -3 {
					valid = false
				}
			default:
				valid = false
			}

			if prevState != Unset && prevState != state {
				valid = false
			}

			if !valid {
				if !dampened {
					e = append(e[:i], e[i+1:]...)
					valid = true
				} else {
					break
				}
			}
		}

		res <- valid
	}
	close(res)
}

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	inputs := make(chan string)
	outputs := make(chan []int)
	entryRes := make(chan bool)
	validEntries := 0
	totalEntries := 0

	go parseWorker(inputs, outputs)
	go evalWorker(outputs, entryRes)

	go func() {
		for v := range entryRes {
			if v {
				validEntries++
			}
		}
	}()

	b := bufio.NewScanner(r)
	for b.Scan() {
		t := b.Text()
		inputs <- t
		totalEntries++
	}

	fmt.Printf("Valid entries: %d\n", validEntries)
	fmt.Printf("Total entries: %d\n", totalEntries)
}
