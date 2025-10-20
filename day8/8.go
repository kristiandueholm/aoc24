package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Pos struct {
	Row int
	Col int
}

func getPosDistance(a Pos, b Pos) (int, int) {
	return a.Row - b.Row, a.Col - b.Col
}

func withinBounds(pos Pos, maxRow int, maxCol int) bool {
	return pos.Row >= 0 &&
		pos.Row <= maxRow &&
		pos.Col >= 0 &&
		pos.Col <= maxCol
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	reader := bufio.NewReader(input)

	runeMap := make(map[rune][]Pos)
	row, col := 0, 0
	maxRow, maxCol := 0, 0

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			maxRow = row
			maxCol = col - 1
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if r == '\n' {
			row++
			col = 0
			continue
		}
		if r == '.' {
			col++
			continue
		}

		runeMap[r] = append(runeMap[r], Pos{Row: row, Col: col})

		col++
	}

	antinodesP1 := make(map[Pos]struct{})

	for _, positions := range runeMap {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				antennaA := positions[i]
				antennaB := positions[j]
				rowDist, colDist := getPosDistance(antennaA, antennaB)
				possibleAntinodesP1 := []Pos{
					{Row: antennaA.Row + rowDist, Col: antennaA.Col + colDist},
					{Row: antennaA.Row - rowDist, Col: antennaA.Col - colDist},
					{Row: antennaB.Row + rowDist, Col: antennaB.Col + colDist},
					{Row: antennaB.Row - rowDist, Col: antennaB.Col - colDist},
				}
				for _, possibleAntinode := range possibleAntinodesP1 {
					if possibleAntinode == antennaA || possibleAntinode == antennaB {
						continue
					}
					if !withinBounds(possibleAntinode, maxRow, maxCol) {
						continue
					}
					antinodesP1[possibleAntinode] = struct{}{}
				}
				// possibleAntinodesP2 := make([]Pos, 0)
				// n := 1
				// for {

				// }
			}
		}
	}
	fmt.Printf("Part 1: %d\n", len(antinodesP1))
}
