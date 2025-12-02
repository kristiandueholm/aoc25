package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func SignChanged(before int, after int) bool {
	if before < 0 && after > 0 {
		return true
	}
	if before > 0 && after < 0 {
		return true
	}
	return false
}

func PyMod(d int, m int) int {
	d %= m
	if d < 0 {
		d += m
	}
	return d
}

func main() {
	file, err := os.Open("./input/1.in")
	// file, err := os.Open("./test/day1.in")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	dial := 50
	endZero := 0
	signChanges := 0
	divisions := 0
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		n, convErr := strconv.Atoi(line[1:])
		if convErr != nil {
			log.Fatal(convErr)
		}
		var sum int
		if dir == 'L' {
			sum = dial - n
		} else {
			sum = dial + n

		}
		if SignChanged(dial, sum) {
			signChanges++
		}
		dial = PyMod(sum, 100)
		if dial == 0 {
			endZero++
		}
		divisions += (AbsInt(sum) - 1) / 100
	}
	log.Printf("Part 1: %d \n", endZero)
	log.Printf("Part 2: %d \n", endZero+signChanges+divisions)
}
