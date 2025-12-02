package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetEvenSubstrings(s string) []string {
	n := len(s)
	var l []string
	for i := 0; i < n; i++ {
		for j := i + 2; j < n; j += 2 {
			l = append(l, s[i:j])
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
	if err != nil {
		log.Fatal(err)
	}
	fileContent := strings.TrimSpace(string(file))
	intervals := strings.Split(fileContent, ",")
	sum := 0
	for _, interval := range intervals {
		split := strings.Split(interval, "-")
		from := split[0]
		to := split[1]
		l := GetIntervalList(from, to)
		for _, item := range l {
			substrings := GetEvenSubstrings(item)
			for _, substring := range substrings {
				mid := len(substring) / 2
				left := substring[:mid]
				right := substring[mid:]
				if left == right {
					conv, convErr := strconv.Atoi(substring)
					if convErr != nil {
						log.Fatal(convErr)
					}
					sum += conv
				}
			}
		}
	}
	fmt.Printf("Part 1: %d", sum)
}
