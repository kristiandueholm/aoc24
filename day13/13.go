package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Loc struct {
	X int
	Y int
}

type Action struct {
	Dir  Loc
	Cost int
}

func ScanParagraphs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 2, bytes.TrimSpace(data[:i]), nil
	}
	if atEOF && len(data) > 0 {
		return len(data), bytes.TrimSpace(data), nil
	}
	return 0, nil, nil
}

func GetLoc(line string) Loc {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(line, -1)
	x, xErr := strconv.Atoi(matches[0])
	if xErr != nil {
		log.Fatal(xErr)
	}
	y, yErr := strconv.Atoi(matches[1])
	if yErr != nil {
		log.Fatal(yErr)
	}
	return Loc{X: x, Y: y}
}

func SolveEquations(a1 Action, a2 Action, t Loc, lim int) int {
	detA := a1.Dir.X*a2.Dir.Y - a2.Dir.X*a1.Dir.Y
	if detA == 0 {
		return 0
	}
	numA := t.X*a2.Dir.Y - t.Y*a2.Dir.X
	numB := t.Y*a1.Dir.X - t.X*a1.Dir.Y

	if numA%detA != 0 || numB%detA != 0 {
		return 0 // not an integer solution
	}
	a1Count := numA / detA
	a2Count := numB / detA

	if a1Count < 0 || a2Count < 0 || a1Count > lim || a2Count > lim {
		return 0
	}

	return a1Count*a1.Cost + a2Count*a2.Cost
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Split(ScanParagraphs)
	part1, part2 := 0, 0
	for scanner.Scan() {
		machine := strings.Split(scanner.Text(), "\n")
		buttonA := Action{Dir: GetLoc(machine[0]), Cost: 3}
		buttonB := Action{Dir: GetLoc(machine[1]), Cost: 1}
		prize1 := GetLoc(machine[2])
		prize2 := Loc{X: prize1.X + 10000000000000, Y: prize1.Y + 10000000000000}
		part1 += SolveEquations(buttonA, buttonB, prize1, 100)
		part2 += SolveEquations(buttonA, buttonB, prize2, 999999999999999999)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
