package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func IntifyBank(bank string) []int {
	bankInts := make([]int, len(bank))
	for i, bat := range bank {
		bankInts[i] = int(bat - '0')
	}
	return bankInts
}

func MaxJolts(bank string) int {
	bankInts := IntifyBank(bank)
	max := 0
	// Double pointer loop
	for i := 0; i < len(bankInts)-1; i++ {
		for j := i + 1; j < len(bankInts); j++ {
			bat := bankInts[i]*10 + bankInts[j]
			if bat > max {
				max = bat
			}
		}
	}
	return max
}

func Part2(bank string) int {
	bankInts := IntifyBank(bank)
	prevBest := -1
	battery := 0
	for pos := 12; pos > 0; pos-- {
		bestIdx := prevBest + 1
		for i := bestIdx; i <= len(bankInts)-pos; i++ {
			if bankInts[i] > bankInts[bestIdx] {
				bestIdx = i
			}
		}
		prevBest = bestIdx
		battery = battery*10 + bankInts[bestIdx]
	}
	return battery
}

func main() {
	file, err := os.Open("input/3.in")
	// file, err := os.Open("test/3.test")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	part1 := 0
	part2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		part1 += MaxJolts(line)
		part2 += Part2(line)
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
