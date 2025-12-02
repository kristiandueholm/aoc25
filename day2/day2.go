package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CheckSubstringList(l []string) bool {
	for i := range l {
		if i == 0 {
			continue
		}
		if l[i] != l[i-1] {
			return false
		}
	}
	return true
}

func GetAllSubstrings(s string) [][]string {
	n := len(s)
	var fullDivisions []int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			fullDivisions = append(fullDivisions, i)
		}
	}
	l := make([][]string, len(fullDivisions))
	for i, div := range fullDivisions {
		nChunks := n / div
		l[i] = make([]string, nChunks)
		for j := 0; j < nChunks; j++ {
			start := j * div
			end := start + div
			l[i][j] = s[start:end]
		}
	}
	return l
}

func GetIntervalList(from string, to string) []string {
	fromInt, fromErr := strconv.Atoi(from)
	if fromErr != nil {
		log.Fatal(fromErr)
	}
	toInt, toErr := strconv.Atoi(to)
	if toErr != nil {
		log.Fatal(toErr)
	}
	n := toInt - fromInt + 1
	l := make([]string, n)
	for i := fromInt; i <= toInt; i++ {
		l[i-fromInt] = strconv.Itoa(i)
	}
	return l
}

func main() {
	file, err := os.ReadFile("./input/2.in")
	// file, err := os.ReadFile("./test/day2.in")
	if err != nil {
		log.Fatal(err)
	}
	fileContent := strings.TrimSpace(string(file))
	intervals := strings.Split(fileContent, ",")
	part1 := 0
	part2 := 0
	for _, interval := range intervals {
		split := strings.Split(interval, "-")
		from := split[0]
		to := split[1]
		l := GetIntervalList(from, to)
		for _, item := range l {
			n := len(item)
			// if n%2 != 0 {
			// 	continue
			// }
			mid := n / 2
			left := item[:mid]
			right := item[mid:]
			itemInt, convErr := strconv.Atoi(item)
			if convErr != nil {
				log.Fatal(convErr)
			}
			if left == right {
				part1 += itemInt
			}
			allSubstrings := GetAllSubstrings(item)
			for _, innerList := range allSubstrings {
				if CheckSubstringList(innerList) {
					part2 += itemInt
					break
				}
			}
		}
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
