package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func fillNInt(x int, n int, arr *[]int) {
	for range n {
		*arr = append(*arr, x)
	}
}

func getChecksum(arr []int) int {
	sum := 0
	for i, id := range arr {
		sum += i * id
	}
	return sum
}

func getDenseArray(input io.ReadCloser) []int {
	reader := bufio.NewReader(input)
	dense := make([]int, 0)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if r == '\n' {
			continue
		}
		dense = append(dense, int(r-'0'))
	}
	return dense
}

func part1(dense []int) {
	result := make([]int, 0)
	i, j := 0, len(dense)-1
	for i < j {
		if i%2 == 0 {
			fillNInt(i/2, dense[i], &result)
			i++
		} else {
			// Fill up from the back
			if dense[i] > dense[j] {
				// fill, move j, decrease val at i, keep i
				fillNInt(j/2, dense[j], &result)
				dense[i] -= dense[j]
				j -= 2
			} else if dense[i] < dense[j] {
				// fill, move i, decrease val at j, keep j
				fillNInt(j/2, dense[i], &result)
				dense[j] -= dense[i]
				i++
			} else {
				// just fill and move both
				fillNInt(j/2, dense[j], &result)
				i++
				j -= 2
			}
		}
	}
	fmt.Println(getChecksum(result))
}

type ItemType int

const (
	File ItemType = iota
	Gap
)

type DiskLocation struct {
	Size  int
	Type  ItemType
	Moved bool
	Index int
}

func part2(dense []int) {
	locations := make([]DiskLocation, len(dense))

	for i := range dense {
		if i%2 == 0 {
			locations[i] = DiskLocation{Size: dense[i], Type: File, Index: i / 2}
		} else {
			locations[i] = DiskLocation{Size: dense[i], Type: Gap, Index: i / 2}
		}
	}

	// j on items from end moving down

	for j := len(dense) - 1; j > 0; j-- {
		if locations[j].Moved {
			continue
		}
		if locations[j].Type == Gap {
			continue
		}
		// i on gaps from start to j
		for i := 1; i < j; i++ {
			if locations[i].Type != Gap {
				continue
			}
			if locations[i].Size == locations[j].Size {
				locations[i] = locations[j]
				locations[i].Moved = true
				locations[j].Type = Gap
				break
			}
			// Need to make a gap for remaining size
			if locations[i].Size > locations[j].Size {
				remainingSize := locations[i].Size - locations[j].Size
				locations[i] = locations[j]
				locations[i].Moved = true
				locations[j].Type = Gap
				newGap := DiskLocation{Size: remainingSize, Type: Gap}
				newGapIndex := i + 1
				locations = append(locations, DiskLocation{})
				copy(locations[newGapIndex+1:], locations[newGapIndex:])
				locations[newGapIndex] = newGap
				break
			}
		}
		//j -= 1
	}
	result := make([]int, 0)
	for _, l := range locations {
		if l.Type == Gap {
			fillNInt(0, l.Size, &result)
		} else {
			fillNInt(l.Index, l.Size, &result)
		}
	}

	fmt.Println(result[:50])

	fmt.Println(getChecksum(result))
}

func passCopy(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	dense := getDenseArray(input)
	part1(passCopy(dense))
	part2(passCopy(dense))
}
