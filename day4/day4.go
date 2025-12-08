package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Item int

const (
	Empty Item = iota
	Paper
)

func GetDimensions(file io.ReadCloser) (int, int) {
	scanner := bufio.NewScanner(file)
	rows, cols := 0, 0
	for rows = 0; scanner.Scan(); rows++ {
		cols = len(scanner.Text())
	}
	return rows, cols
}

func GetItem(r rune) Item {
	switch r {
	case '@':
		return Paper
	case '.':
		return Empty
	default:
		panic("Bad item")
	}
}

func CountPaper(positions [][]Item, row int, col int) int {
	maxRow := len(positions) - 1
	maxCol := len(positions[0]) - 1

	dr := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dc := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	neighbors := 0

	for i := range dr {
		r := row + dr[i]
		c := col + dc[i]
		if r >= 0 && c >= 0 && r <= maxRow && c <= maxCol {
			if positions[r][c] == Paper {
				neighbors++
			}
		}
	}
	return neighbors
}

func main() {
	// file, err := os.Open("test/4.test")
	file, err := os.Open("input/4.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	rows, cols := GetDimensions(file)
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	positions := make([][]Item, rows)
	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		positions[row] = make([]Item, cols)
		for col := range line {
			positions[row][col] = GetItem(rune(line[col]))
		}
	}
	part1 := 0
	for i, row := range positions {
		for j := range row {
			neighbors := CountPaper(positions, i, j)
			if neighbors < 4 && positions[i][j] == Paper {
				part1++
			}
		}
	}
	part2 := 0
	prevRemoved := -1
	for prevRemoved != 0 {
		prevRemoved = 0
		for i, row := range positions {
			for j := range row {
				neighbors := CountPaper(positions, i, j)
				if neighbors < 4 && positions[i][j] == Paper {
					part2++
					prevRemoved++
					positions[i][j] = Empty
				}
			}
		}
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
