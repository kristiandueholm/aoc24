package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Location struct {
	Col int
	Row int
}

func buildMap(input io.ReadCloser) map[int][]Location {
	reader := bufio.NewReader(input)
	heightMap := make(map[int][]Location, 0)

	nRow, nChar := 0, 0
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if char == '\n' {
			nChar = 0
			nRow++
			continue
		}
		digit := int(char - '0')
		loc := Location{Row: nRow, Col: nChar}
		heightMap[digit] = append(heightMap[digit], loc)
		nChar++
	}
	return heightMap
}

func getNextNeighbours(loc Location, height int, heightMap map[int][]Location, visited map[Location]bool) int {
	if height == 9 && !visited[loc] {
		visited[loc] = true
		return 1
	}

	sum := 0

	nextHeightLocs := heightMap[height+1]
	for _, nextLoc := range nextHeightLocs {
		dx := loc.Col - nextLoc.Col
		dy := loc.Row - nextLoc.Row
		if (dx == 1 && dy == 0) || (dx == -1 && dy == 0) || (dx == 0 && dy == 1) || (dx == 0 && dy == -1) {
			sum += getNextNeighbours(nextLoc, height+1, heightMap, visited)
		}
	}

	return sum
}

func getNextNeighboursP2(loc Location, height int, heightMap map[int][]Location) int {
	if height == 9 {
		return 1
	}

	sum := 0

	nextHeightLocs := heightMap[height+1]
	for _, nextLoc := range nextHeightLocs {
		dx := loc.Col - nextLoc.Col
		dy := loc.Row - nextLoc.Row
		if (dx == 1 && dy == 0) || (dx == -1 && dy == 0) || (dx == 0 && dy == 1) || (dx == 0 && dy == -1) {
			sum += getNextNeighboursP2(nextLoc, height+1, heightMap)
		}
	}

	return sum
}

func part1(heightMap map[int][]Location) int {
	sum := 0
	for _, trailStart := range heightMap[0] {
		visited := make(map[Location]bool, 0)
		sum += getNextNeighbours(trailStart, 0, heightMap, visited)
	}
	return sum
}

func part2(heightMap map[int][]Location) int {
	sum := 0
	for _, trailStart := range heightMap[0] {
		sum += getNextNeighboursP2(trailStart, 0, heightMap)
	}
	return sum
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	heightMap := buildMap(input)
	part1 := part1(heightMap)
	part2 := part2(heightMap)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
