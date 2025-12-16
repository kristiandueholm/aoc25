package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func runify(line string) []rune {
	runes := make([]rune, 0)
	for _, c := range line {
		if c != ' ' {
			runes = append(runes, c)
		}
	}
	return runes
}

func CalcResult(columns []string, op rune) int {
	sum := 0
	for _, column := range columns {
		converted, err := strconv.Atoi(column)
		if err != nil {
			log.Fatal(err)
		}
		if sum == 0 {
			sum += converted
			continue
		}
		switch op {
		case '+':
			sum += converted
		case '*':
			sum *= converted
		default:
			panic("Yikes")
		}

	}

	return sum
}

func Stuff(lines []string) int {
	maxWidth := 0
	runeLines := make([][]rune, len(lines))

	for i, line := range lines {
		runeLine := []rune(line)
		runeLines[i] = runeLine
		if len(runeLine) > maxWidth {
			maxWidth = len(runeLine)
		}
	}

	sum := 0
	curOp := '*'
	var builder strings.Builder
	columns := make([]string, 0)
	for j := 0; j < maxWidth; j++ {
		if j < len(runeLines[len(lines)-1]) && runeLines[len(lines)-1][j] != ' ' {
			sum += CalcResult(columns, curOp)
			curOp = runeLines[len(lines)-1][j]
			columns = make([]string, 0)
		}
		builder.Reset()
		for i := 0; i < len(lines)-1; i++ {
			// Just build your strings
			if runeLines[i][j] != ' ' {
				_, writeErr := builder.WriteRune(runeLines[i][j])
				if writeErr != nil {
					log.Fatal(writeErr)
				}
			}
		}
		if builder.String() != "" {
			columns = append(columns, builder.String())
		}
	}
	sum += CalcResult(columns, curOp)
	return sum
}

func main() {
	contents, err := os.ReadFile("./input/6.in")
	// contents, err := os.ReadFile("./test/6.in")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(contents)), "\n")
	ops := runify(lines[len(lines)-1])
	columns := make([]int, len(lines[0]))

	for i, column := range strings.Fields(lines[0]) {
		columnInt, convErr := strconv.Atoi(column)
		if convErr != nil {
			log.Fatal(convErr)
		}
		columns[i] = columnInt
	}

	for _, line := range lines[1 : len(lines)-1] {
		for i, column := range strings.Fields(line) {
			columnInt, convErr := strconv.Atoi(column)
			if convErr != nil {
				log.Fatal(convErr)
			}
			switch ops[i] {
			case '+':
				columns[i] += columnInt
			case '*':
				columns[i] *= columnInt
			default:
				panic("Yikes")
			}
		}
	}
	part1 := 0
	for _, column := range columns {
		part1 += column
	}
	part2 := Stuff(lines)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
