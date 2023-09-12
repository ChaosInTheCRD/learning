package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type backpack struct {
	A string
	B string
}

func parseInput(path string) ([]backpack, error) {
	f, err := os.Open(path)
	if err != nil {
		return []backpack{}, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	backpacks := []backpack{}

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return []backpack{}, err
		}

		h := len(line) / 2
		a := line[:h]
		b := line[h:]
		backpacks = append(backpacks, backpack{A: a, B: b})
	}

	return backpacks, nil
}

func findDuplicateTypes(b backpack) int {
	for i := range b.A {
		if strings.Contains(b.B, string(b.A[i])) {
			return priority(rune(b.A[i]))
		}
	}

	return 0
}

func priority(c rune) int {
	if unicode.IsUpper(c) {
		return int(c - 'A' + 27)
	} else {
		return int(c - 'a' + 1)
	}
}

func main() {
	backpacks, err := parseInput("./input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	total := 0
	for _, n := range backpacks {
		total = total + findDuplicateTypes(n)
	}

	fmt.Println(total)
}
