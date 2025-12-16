package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("./test/7.in")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	nSplits := 0
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		lastSplit := 0
		for j, symbol := range line {
			above := rune(lines[i][j])
			// TODO: This does not work, I need to infer beams
			if above == '|' && symbol == '^' {
				dist := j - lastSplit
				if dist > 1 {
					nSplits += 2
				} else {
					nSplits += 1
				}
			}
		}
	}
	part1 := nSplits
	fmt.Printf("Part 1: %d\n", part1)
}
