package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	stones := make([]int, 0)
	for scanner.Scan() {
		stone, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		stones = append(stones, stone)
	}

	for i, stone := range stones {
		stoneString := strconv.Itoa(stone)
		nDigits := len(stoneString)
		if stone == 0 {
			stones[i] = 1
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
			// mutating length of the array I am iterating over??
		}
	}
}
