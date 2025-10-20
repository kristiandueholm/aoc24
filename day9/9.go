package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func getDotString(input io.ReadCloser) string {
	reader := bufio.NewReader(input)
	n := 0
	var builder strings.Builder
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		digit := int(b - '0')
		var c string
		if n%2 == 0 {
			c = strconv.Itoa(n)
		} else {
			c = "."
		}
		s := strings.Repeat(c, digit)
		_, err = builder.WriteString(s)
		if err != nil {
			log.Fatal(err)
		}
		n++
	}
	return builder.String()
}

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

func part1(input io.ReadCloser) {
	reader := bufio.NewReader(input)
	dense := make([]int, 0)
	result := make([]int, 0)
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

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	part1(input)
}
