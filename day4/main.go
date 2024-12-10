package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	DirVert Direction = iota
	DirHor
	DirVertBackwards
	DirHorBackwards
	DirUpDiagLeft
	DirUpDiagRight
	DirDownDiagLeft
	DirDownDiagRight
)

type SearchMap [][]byte

func readInput(filename string) (SearchMap, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	m := make(SearchMap, 0)
	for _, l := range bytes.Split(data, []byte("\n")) {
		m = append(m, l)
	}

	return m, nil
}

func getKeyMap(i int, j int, dir Direction) string {
	switch dir {
	case DirVert:
		return fmt.Sprintf("%d,%d", i+1, j)
	case DirHor:
		return fmt.Sprintf("%d,%d", i, j+1)
	case DirVertBackwards:
		return fmt.Sprintf("%d,%d", i-1, j)
	case DirHorBackwards:
		return fmt.Sprintf("%d,%d", i, j-1)
	case DirUpDiagLeft:
		return fmt.Sprintf("%d,%d", i-1, j-1)
	case DirUpDiagRight:
		return fmt.Sprintf("%d,%d", i-1, j+1)
	case DirDownDiagLeft:
		return fmt.Sprintf("%d,%d", i+1, j-1)
	case DirDownDiagRight:
		return fmt.Sprintf("%d,%d", i+1, j+1)
	}
	return ""
}

func part1(input SearchMap) int {
	m := make(map[string]string)
	count := 0
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1; j++ {
			if input[i][j] == 'X' {
				m[fmt.Sprintf("%d,%d", i, j)] = "X"
			}
			if input[i][j] == 'M' {
				m[fmt.Sprintf("%d,%d", i, j)] = "M"
			}
			if input[i][j] == 'A' {
				m[fmt.Sprintf("%d,%d", i, j)] = "A"
			}
			if input[i][j] == 'S' {
				m[fmt.Sprintf("%d,%d", i, j)] = "S"
			}
		}
	}

	for xKey := range m {
		if m[xKey] == "X" {
			xFrags := strings.Split(xKey, ",")
			x1, _ := strconv.Atoi(xFrags[0])
			x2, _ := strconv.Atoi(xFrags[1])

			for _, dir := range []Direction{
				DirVert,
				DirHor,
				DirVertBackwards,
				DirHorBackwards,
				DirUpDiagLeft,
				DirUpDiagRight,
				DirDownDiagLeft,
				DirDownDiagRight,
			} {
				mKey := getKeyMap(x1, x2, dir)
				if m[mKey] == "M" {
					mFrags := strings.Split(mKey, ",")
					m1, _ := strconv.Atoi(mFrags[0])
					m2, _ := strconv.Atoi(mFrags[1])

					aKey := getKeyMap(m1, m2, dir)
					if m[aKey] == "A" {
						aFrags := strings.Split(aKey, ",")
						a1, _ := strconv.Atoi(aFrags[0])
						a2, _ := strconv.Atoi(aFrags[1])

						sKey := getKeyMap(a1, a2, dir)
						if m[sKey] == "S" {
							count++
						}
					}
				}
			}
		}
	}

	fmt.Printf("%v\n", count)
	return count
}

func part2(input SearchMap) int {
	m := make(map[string]string)
	count := 0
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1; j++ {
			if input[i][j] == 'M' {
				m[fmt.Sprintf("%d,%d", i, j)] = "M"
			}
			if input[i][j] == 'A' {
				m[fmt.Sprintf("%d,%d", i, j)] = "A"
			}
			if input[i][j] == 'S' {
				m[fmt.Sprintf("%d,%d", i, j)] = "S"
			}
		}
	}

	for mKey := range m {
		if m[mKey] == "M" {
			mFrags := strings.Split(mKey, ",")
			m1, _ := strconv.Atoi(mFrags[0])
			m2, _ := strconv.Atoi(mFrags[1])

			for _, dir := range []Direction{
				DirUpDiagLeft,
				DirUpDiagRight,
				DirDownDiagLeft,
				DirDownDiagRight,
			} {
				aKey := getKeyMap(m1, m2, dir)
				if m[aKey] == "A" {
					aFrags := strings.Split(mKey, ",")
					a1, _ := strconv.Atoi(aFrags[0])
					a2, _ := strconv.Atoi(aFrags[1])

					sKey := getKeyMap(a1, a2, dir)
					if m[sKey] == "S" {
						count++
					}
				}
			}
		}
	}

	fmt.Printf("%v\n", count)
	return count
}

func main() {
	sm, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(sm)
	part2(sm)
	//fmt.Printf("%v", sm)
}
