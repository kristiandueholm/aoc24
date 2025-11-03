package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Loc struct {
	X int
	Y int
}

type Obstacle int

const (
	None Obstacle = iota
	Wall
	Box
)

func GetMap(obstacles string) map[Loc]Obstacle {
	m := make(map[Loc]Obstacle, 0)
	for i, line := range strings.Split(obstacles, "\n") {
		for j, b := range line {
			switch b {
			case '#':
				m[Loc{j, i}] = Wall
			case 'O':
				m[Loc{j, i}] = Box
			}
		}
	}
	return m
}

func GetRobot(obstacles string) Loc {
	for i, line := range strings.Split(obstacles, "\n") {
		for j, r := range line {
			if r == '@' {
				return Loc{j, i}
			}
		}
	}
	panic("could not get robot")
}

func GetActions(s string) string {
	return strings.ReplaceAll(s, "\n", "")
}

func step(action rune, robot Loc, obstacles map[Loc]Obstacle) Loc {
	seq := make([]Loc, 0)
	var dx, dy int
	switch action {
	case '<':
		dx, dy = -1, 0
	case '>':
		dx, dy = +1, 0
	case '^':
		dx, dy = 0, -1
	case 'v':
		dx, dy = 0, +1
	}

	i := 1
	for {
		next := Loc{robot.X + dx*i, robot.Y + dy*i}
		seq = append(seq, next)
		if obstacles[next] != Box {
			break
		}
		i++
	}

	// Sequence aquired
	// Check if moving is possible

	n := len(seq)
	// not moving at all
	if obstacles[seq[n-1]] == Wall {
		return robot
	}
	if n == 1 {
		return seq[0]
	}

	obstacles[seq[0]] = None
	for _, l := range seq[1:] {
		obstacles[l] = Box
	}
	return seq[0]
}

func CalculateCoordSum(obstacles map[Loc]Obstacle) int {
	sum := 0
	for k, v := range obstacles {
		if v == Box {
			sum += k.X + 100*k.Y
		}
	}
	return sum
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputParts := strings.Split(string(input), "\n\n")
	robot := GetRobot(inputParts[0])
	obstacles := GetMap(inputParts[0])
	actions := GetActions(inputParts[1])
	for _, a := range actions {
		robot = step(a, robot, obstacles)
	}
	part1 := CalculateCoordSum(obstacles)
	fmt.Printf("Part 1: %d\n", part1)
}

