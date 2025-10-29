package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Plot struct {
	Row   int
	Col   int
	Plant rune
}

func isWithinBounds(loc [2]int, maxRow, maxCol int) bool {
	return loc[0] >= 0 && loc[0] <= maxRow &&
		loc[1] >= 0 && loc[1] <= maxCol
}

func getNeighbors(plot Plot, grid [][]rune) []Plot {
	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1
	neighborLocs := [][2]int{
		{plot.Row - 1, plot.Col},
		{plot.Row + 1, plot.Col},
		{plot.Row, plot.Col - 1},
		{plot.Row, plot.Col + 1},
	}

	neighbors := make([]Plot, 0)

	for _, loc := range neighborLocs {
		if isWithinBounds(loc, maxRow, maxCol) {
			row := loc[0]
			col := loc[1]
			plant := grid[row][col]
			neighborPlot := Plot{Row: row, Col: col, Plant: plant}
			neighbors = append(neighbors, neighborPlot)
		}
	}
	return neighbors
}

func countCorners(plot Plot, grid [][]rune) int {
	topDiff := plot.Row == 0 || plot.Plant != grid[plot.Row-1][plot.Col]
	botDiff := plot.Row == len(grid)-1 || plot.Plant != grid[plot.Row+1][plot.Col]
	leftDiff := plot.Col == 0 || plot.Plant != grid[plot.Row][plot.Col-1]
	rightDiff := plot.Col == len(grid[0])-1 || plot.Plant != grid[plot.Row][plot.Col+1]
	topLeftDiagDiff := plot.Row != 0 && plot.Col != 0 && grid[plot.Row-1][plot.Col-1] != plot.Plant
	topRightDiagDiff := plot.Row != 0 && plot.Col != len(grid[0])-1 && grid[plot.Row-1][plot.Col+1] != plot.Plant
	botLeftDiagDiff := plot.Row != len(grid)-1 && plot.Col != 0 && grid[plot.Row+1][plot.Col-1] != plot.Plant
	botRightDiagDiff := plot.Row != len(grid)-1 && plot.Col != len(grid[0])-1 && grid[plot.Row+1][plot.Col+1] != plot.Plant

	corners := 0
	// Top left: Convex (outer) OR Concave (inner)
	if (topDiff && leftDiff) || (!topDiff && !leftDiff && topLeftDiagDiff) {
		corners++
	}
	// Top right: Convex (outer) OR Concave (inner)
	if (topDiff && rightDiff) || (!topDiff && !rightDiff && topRightDiagDiff) {
		corners++
	}
	// Bot left: Convex (outer) OR Concave (inner)
	if (botDiff && leftDiff) || (!botDiff && !leftDiff && botLeftDiagDiff) {
		corners++
	}
	// Bot right: Convex (outer) OR Concave (inner)
	if (botDiff && rightDiff) || (!botDiff && !rightDiff && botRightDiagDiff) {
		corners++
	}
	return corners
}

func searchRegion(plot Plot, grid [][]rune, visitedGrid [][]bool) (int, int, int) {
	if visitedGrid[plot.Row][plot.Col] {
		return 0, 0, 0
	}

	visitedGrid[plot.Row][plot.Col] = true

	area := 1
	perimeter := 0
	corners := countCorners(plot, grid)

	for _, neighbor := range getNeighbors(plot, grid) {
		if neighbor.Plant != plot.Plant {
			perimeter++
		} else if !visitedGrid[neighbor.Row][neighbor.Col] && neighbor.Plant == plot.Plant {
			a, p, c := searchRegion(neighbor, grid, visitedGrid)
			area += a
			perimeter += p
			corners += c
		}
	}

	if plot.Row == 0 || plot.Row == len(grid)-1 {
		perimeter++
	}
	if plot.Col == 0 || plot.Col == len(grid[0])-1 {
		perimeter++
	}

	return area, perimeter, corners
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	grid := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}

	part1 := 0
	part2 := 0

	for i := range grid {
		for j := range grid[0] {
			if !visited[i][j] {
				plot := Plot{Row: i, Col: j, Plant: grid[i][j]}
				area, perimeter, corners := searchRegion(plot, grid, visited)
				part1 += area * perimeter
				part2 += area * corners
			}
		}
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)

}
