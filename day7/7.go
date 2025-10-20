package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(line string) (int, []int) {
	splits := strings.Split(strings.TrimSpace(line), " ")
	if len(splits) < 1 {
		panic("Found line without hope")
	}

	result, err := strconv.Atoi(strings.TrimRight(splits[0], ":"))
	if err != nil {
		log.Fatal(err)
	}

	paramsStrings := splits[1:]

	params := make([]int, len(paramsStrings))
	for i, s := range paramsStrings {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		params[i] = n
	}

	return result, params
}

func intConcat(a int, b int) int {
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)
	res, err := strconv.Atoi(strA + strB)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func searchp1(node int, acc []int, expected int) bool {
	if len(acc) == 0 {
		return node == expected
	}
	return searchp1(node+acc[0], acc[1:], expected) || searchp1(node*acc[0], acc[1:], expected)
}
func searchp2(node int, acc []int, expected int) bool {
	if len(acc) == 0 {
		return node == expected
	}
	plus := searchp2(node+acc[0], acc[1:], expected)
	times := searchp2(node*acc[0], acc[1:], expected)
	concat := searchp2(intConcat(node, acc[0]), acc[1:], expected)
	return plus || times || concat
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1 := 0
	part2 := 0

	for scanner.Scan() {
		result, params := parse(scanner.Text())
		if searchp1(params[0], params[1:], result) {
			part1 += result
		}
		if searchp2(params[0], params[1:], result) {
			part2 += result
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
