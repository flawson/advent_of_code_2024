package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type GuardState byte

const (
	StateUp    GuardState = '^'
	StateDown  GuardState = 'v'
	StateLeft  GuardState = '<'
	StateRight GuardState = '>'

	FreeSpace   = '.'
	Obstruction = '#'
	Moved       = 'X'
)

func IsGuard(chk GuardState) (bool, GuardState) {
	switch chk {
	case StateUp:
		return true, StateUp
	case StateDown:
		return true, StateDown
	case StateLeft:
		return true, StateLeft
	case StateRight:
		return true, StateRight
	default:
		return false, 0
	}
}

type LayoutMap [][]byte
type GuardPos struct {
	Row         int
	Col         int
	Orientation GuardState
}

func (lm LayoutMap) String() string {
	rows := []string{}
	for row := range lm {
		rows = append(rows, string(lm[row]))
	}
	return strings.Join(rows, "\n")
}

func (lm LayoutMap) MoveGuard(gp GuardPos) GuardPos {
	switch gp.Orientation {
	case StateUp:
		if gp.Row == 0 {
			lm[gp.Row][gp.Col] = Moved
			return GuardPos{gp.Row, gp.Col, 0}
		}
		if lm[gp.Row-1][gp.Col] != Obstruction {
			lm[gp.Row-1][gp.Col] = byte(StateUp)
			lm[gp.Row][gp.Col] = Moved
			return GuardPos{gp.Row - 1, gp.Col, StateUp}
		} else {
			return GuardPos{gp.Row, gp.Col, StateRight}
		}
	case StateDown:
		if gp.Row == len(lm)-1 {
			lm[gp.Row][gp.Col] = Moved
			return GuardPos{gp.Row, gp.Col, 0}
		}
		if lm[gp.Row+1][gp.Col] != Obstruction {
			lm[gp.Row+1][gp.Col] = byte(StateDown)
			lm[gp.Row][gp.Col] = Moved
			return GuardPos{gp.Row + 1, gp.Col, StateDown}
		} else {
			return GuardPos{gp.Row, gp.Col, StateLeft}
		}
	case StateLeft:
		if gp.Col == 0 {
			lm[gp.Row][gp.Col] = Moved
			return GuardPos{gp.Row, gp.Col, 0}
		}
		if lm[gp.Row][gp.Col-1] != Obstruction {
			lm[gp.Row][gp.Col-1] = byte(StateLeft)
			lm[gp.Row][gp.Col] = Moved
			return GuardPos{gp.Row, gp.Col - 1, StateLeft}
		} else {
			return GuardPos{gp.Row, gp.Col, StateUp}
		}
	case StateRight:
		if gp.Col == len(lm)-1 {
			lm[gp.Row][gp.Col] = Moved
			return GuardPos{gp.Row, gp.Col, 0}
		}
		if lm[gp.Row][gp.Col+1] != Obstruction {
			lm[gp.Row][gp.Col+1] = byte(StateRight)
			lm[gp.Row][gp.Col] = Moved
			return GuardPos{gp.Row, gp.Col + 1, StateRight}
		} else {
			return GuardPos{gp.Row, gp.Col, StateDown}
		}
	}
	return GuardPos{gp.Row, gp.Col, 0}
}

func (lm LayoutMap) FindMoveCount() int {
	count := 0

	for _, row := range lm {
		for _, val := range row {
			if val == Moved {
				count++
			}
		}
	}

	return count
}

func readInput(r io.Reader) (LayoutMap, GuardPos) {
	lm := LayoutMap{}
	var gp GuardPos

	buffer := bufio.NewScanner(r)
	row := 0
	for buffer.Scan() {
		line := buffer.Bytes()
		lm = append(lm, make([]byte, len(line)))

		for col, val := range line {
			lm[row][col] = val
			if ok, pos := IsGuard(GuardState(val)); ok {
				gp = GuardPos{row, col, pos}
			}
		}
		row++
	}

	return lm, gp
}

func main() {
	filename := "input.txt"
	fh, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}

	lm, gp := readInput(fh)

	for {
		gp = lm.MoveGuard(gp)
		if gp.Orientation == 0 {
			break
		}
	}
	fmt.Printf("%v\n", lm)
	fmt.Printf("%v\n", gp)
	fmt.Printf("%v\n", lm.FindMoveCount())
}
