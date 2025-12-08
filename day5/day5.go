package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Interval struct {
	From int
	To   int
}

func IdInInterval(id int, intervals []Interval) bool {
	for _, interval := range intervals {
		if interval.From <= id && id <= interval.To {
			return true
		}
	}
	return false
}

func MergeIntervals(intervals []Interval) []Interval {
	intervalsSorted := make([]Interval, len(intervals))
	copy(intervalsSorted, intervals)
	slices.SortFunc(intervalsSorted, func(a, b Interval) int {
		return a.From - b.From
	})
	merged := make([]Interval, 0)
	for i, interval := range intervalsSorted {
		N := len(merged)
		if i == 0 {
			merged = append(merged, interval)
		} else if interval.From <= merged[N-1].To && interval.To <= merged[N-1].To {
			// Prev wraps current completely
			continue
		} else if interval.From <= merged[N-1].To {
			// Overlapping intervals, merge
			merged[N-1].To = interval.To
		} else {
			// Non-overlapping, append
			merged = append(merged, interval)
		}
	}
	return merged
}

func CountIntervals(intervals []Interval) int {
	sum := 0
	for _, interval := range intervals {
		sum += interval.To - interval.From + 1
	}
	return sum
}

func main() {
	// file, err := os.ReadFile("test/5.test")
	file, err := os.ReadFile("input/5.in")
	if err != nil {
		log.Fatal(err)
	}
	fileString := string(file)
	divide := strings.Split(fileString, "\n\n")
	intervals := make([]Interval, 0)
	for _, line := range strings.Split(divide[0], "\n") {
		startEnd := strings.Split(line, "-")
		from, fromErr := strconv.Atoi(startEnd[0])
		if fromErr != nil {
			log.Fatal(fromErr)
		}
		to, toErr := strconv.Atoi(startEnd[1])
		if toErr != nil {
			log.Fatal(toErr)
		}
		interval := Interval{From: from, To: to}
		intervals = append(intervals, interval)
	}
	part1 := 0
	for _, line := range strings.Split(divide[1], "\n") {
		id, idErr := strconv.Atoi(line)
		if idErr != nil {
			log.Fatal(idErr)
		}
		if IdInInterval(id, intervals) {
			part1++
		}
	}
	mergedIntervals := MergeIntervals(intervals)
	part2 := CountIntervals(mergedIntervals)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
