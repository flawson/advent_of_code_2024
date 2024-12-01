package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

var lineRegex = regexp.MustCompile(`^(\d+)\s*(\d+)\n$`)

func readInput(reader io.Reader) ([]int, []int, error) {
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	r := bufio.NewReader(reader)

	for {
		s, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, nil, err
		}

		nums := lineRegex.FindStringSubmatch(s)
		if nums == nil || len(nums) != 3 {
			continue
		}

		n1, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, nil, err
		}
		list1 = append(list1, n1)

		n2, err := strconv.Atoi(nums[2])
		if err != nil {
			return nil, nil, err
		}
		list2 = append(list2, n2)
	}

	if len(list1) != len(list2) {
		return nil, nil, fmt.Errorf("len(list1)[%d] != len(list2)[%d]", len(list1), len(list2))
	}

	slices.Sort(list1)
	slices.Sort(list2)

	return list1, list2, nil
}

func part1(l1, l2 []int) {
	dist := 0

	// Assume based on readInput that len(l1) == len(l2)
	for idx := 0; idx < len(l1); idx++ {
		d := l1[idx] - l2[idx]
		if d < 0 {
			d = d * -1
		}
		dist += d
	}

	fmt.Printf("Part 1, total distance diff: %d\n", dist)
}

func part2(l1, l2 []int) {
	l2Freq := make(map[int]int)
	for _, v := range l2 {
		if _, ok := l2Freq[v]; !ok {
			l2Freq[v] = 1
		} else {
			l2Freq[v]++
		}
	}

	simScore := 0

	for _, v := range l1 {
		if count, ok := l2Freq[v]; !ok {
			// If not in list, don't add anything
			continue
		} else {
			simScore += (count * v)
		}
	}

	fmt.Printf("Part 2, similarity score: %d\n", simScore)
}

func main() {
	fh, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	l1, l2, err := readInput(fh)
	if err != nil {
		log.Fatal(err.Error())
	}

	part1(l1, l2)
	part2(l1, l2)
}
