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

func parseInput(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	groups := [][]string{}

	group := []string{}
	for {
		if len(group) == 3 {
			fmt.Println(group)
			groups = append(groups, group)
			group = []string{}
		}

		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return [][]string{}, err
		}

		group = append(group, line)
	}

	return groups, nil
}

func findDuplicateTypes(g []string) int {
	for _, i := range g[0] {
		if strings.Contains(g[1], string(i)) && strings.Contains(g[2], string(i)) {
			return priority(i)
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
	groups, err := parseInput("./input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	total := 0
	for i := range groups {
		total = total + findDuplicateTypes(groups[i])
	}

	fmt.Println(total)
}
