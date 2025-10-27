package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getCount(counts map[int]int) int {
	count := 0
	for _, v := range counts {
		count += v
	}
	return count
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	stones := make(map[int]int, 0)
	for scanner.Scan() {
		stone, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		stones[stone]++
	}

	iterations := 75
	var part1 int

	for i := range iterations {
		if i == 25 {
			part1 = getCount(stones)
		}
		temp := make(map[int]int, 0)
		for stone, count := range stones {
			stoneString := strconv.Itoa(stone)
			nDigits := len(stoneString)
			if stone == 0 {
				temp[1] += count
			} else if nDigits%2 == 0 && nDigits > 0 {
				mid := nDigits / 2
				left, errL := strconv.Atoi(stoneString[:mid])
				if errL != nil {
					log.Fatal(errL)
				}
				right, errR := strconv.Atoi(stoneString[mid:])
				if errR != nil {
					log.Fatal(errR)
				}
				temp[left] += count
				temp[right] += count
			} else {
				temp[stone*2024] += count
			}
		}
		stones = temp
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", getCount(stones))
}
