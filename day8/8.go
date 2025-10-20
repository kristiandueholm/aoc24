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

	antinodes := make(map[Pos]struct{})

	for _, positions := range runeMap {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				antennaA := positions[i]
				antennaB := positions[j]
				rowDist, colDist := getPosDistance(antennaA, antennaB)
				possibleAntinodes := []Pos{
					{Row: antennaA.Row + rowDist, Col: antennaA.Col + colDist},
					{Row: antennaA.Row - rowDist, Col: antennaA.Col - colDist},
					{Row: antennaB.Row + rowDist, Col: antennaB.Col + colDist},
					{Row: antennaB.Row - rowDist, Col: antennaB.Col - colDist},
				}
				for _, possibleAntinode := range possibleAntinodes {
					if possibleAntinode == antennaA || possibleAntinode == antennaB {
						continue
					}
					if possibleAntinode.Col < 0 || possibleAntinode.Col > maxCol {
						continue
					}
					if possibleAntinode.Row < 0 || possibleAntinode.Row > maxRow {
						continue
					}
					antinodes[possibleAntinode] = struct{}{}
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}
