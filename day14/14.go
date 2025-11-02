package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
	X  int
	Y  int
	VX int
	VY int
}

type Loc struct {
	X int
	Y int
}

func GetRobot(line string) Robot {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(line, -1)
	matchesInt := make([]int, len(matches))
	for i, match := range matches {
		m, err := strconv.Atoi(match)
		if err != nil {
			log.Fatal(err)
		}
		matchesInt[i] = m
	}
	return Robot{
		X:  matchesInt[0],
		Y:  matchesInt[1],
		VX: matchesInt[2],
		VY: matchesInt[3],
	}
}

func CalcMovement(r Robot, time int, width int, height int) (int, int) {
	x := (r.X + r.VX*time) % width
	y := (r.Y + r.VY*time) % height
	if x < 0 {
		x += width
	}
	if y < 0 {
		y += height
	}
	return x, y
}

func CalcQuadrants(robotCount map[Loc]int, width int, height int) int {
	quadrants := []int{0, 0, 0, 0}
	for i := range width {
		if i == width/2 {
			continue
		}
		for j := range height {
			if j == height/2 {
				continue
			}
			loc := Loc{X: i, Y: j}
			if i < width/2 && j < height/2 {
				// upper left q0
				quadrants[0] += robotCount[loc]

			} else if i > width/2 && j < height/2 {
				// upper right q1
				quadrants[1] += robotCount[loc]

			} else if i < width/2 && j > height/2 {
				// lower left q2
				quadrants[2] += robotCount[loc]

			} else {
				// lower right q3
				quadrants[3] += robotCount[loc]
			}
		}
	}
	prod := 1
	for _, q := range quadrants {
		prod *= q
	}
	return prod
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	robotCount := make(map[Loc]int, 0)
	width, height, time := 101, 103, 100
	for scanner.Scan() {
		robot := GetRobot(scanner.Text())
		newX, newY := CalcMovement(robot, time, width, height)
		robotLoc := Loc{X: newX, Y: newY}
		robotCount[robotLoc]++
	}
	// testRobot := Robot{
	// 	X:  2,
	// 	Y:  4,
	// 	VX: 2,
	// 	VY: -3,
	// }
	// fmt.Println(CalcMovement(testRobot, 5, 11, 7))
	part1 := CalcQuadrants(robotCount, width, height)
	fmt.Println(part1)
}
