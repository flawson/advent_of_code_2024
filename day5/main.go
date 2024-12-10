package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type PageRules map[int][]int
type PagePrintouts [][]int

var OrderRulesRegex = regexp.MustCompile(`^(\d+)\|(\d+)`)

func readInput(filename string) (PageRules, PagePrintouts) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil
	}

	m := make(PageRules)
	pp := make(PagePrintouts, 0)

	for _, l := range bytes.Split(data, []byte("\n")) {
		matches := OrderRulesRegex.FindStringSubmatch(string(l))
		if matches != nil {
			m1, _ := strconv.Atoi(matches[1])
			m2, _ := strconv.Atoi(matches[2])

			_, ok := m[m1]
			if ok {
				m[m1] = append(m[m1], m2)
			} else {
				m[m1] = []int{m2}
			}
		}

		if strings.Contains(string(l), ",") {
			printout := []int{}
			pages := strings.Split(string(l), ",")

			for _, page := range pages {
				v, _ := strconv.Atoi(page)
				printout = append(printout, v)
			}

			pp = append(pp, printout)
		}

	}

	return m, pp
}

func part1(pRules PageRules, printouts PagePrintouts) {
	midSum := 0
	for _, printout := range printouts {
		revPages := make([]int, 0, len(printout))
		for i := len(printout) - 1; i >= 0; i-- {
			revPages = append(revPages, printout[i])
		}
		fails := false
		fmt.Printf("%v\n", revPages)
		for idx, page := range revPages {
			for i := 0; i < idx; i++ {
				if slices.Contains(pRules[revPages[i]], page) {
					fails = true
				}
			}
		}
		if !fails {
			mid := len(printout) / 2
			fmt.Printf("%v: %d\n", printout, mid)
			midSum += printout[mid]
		}
	}
	fmt.Printf("%v\n", midSum)
}

func main() {
	pRules, printouts := readInput("input.txt")
	part1(pRules, printouts)
}
