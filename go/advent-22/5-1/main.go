package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	stackHeight = 8
)

type procedure struct {
	quantity int
	start    int
	finish   int
}

func buildStacks(stacks [][]rune, line string) [][]rune {
	firstLine := false
	if len(stacks) == 0 {
		firstLine = true
	}

	i := 0
	stack := 0
	for line != "" {
		if (i+1)%4 == 0 && i != 0 {
			line = line[1:]
			i++
			continue
		} else if line[:3] == "   " {
			if firstLine {
				stacks = append(stacks, []rune{})
			} else {
				stack++
			}
			line = line[3:]
			i = i + 3
			continue
		} else {
			if firstLine {
				stacks = append(stacks, []rune{rune(line[1])})
			} else {
				stacks[stack] = append(stacks[stack], rune(line[1]))
				stack++
			}
			line = line[3:]
			i = i + 3
		}
	}

	return stacks
}

func appendProcedure(procedures []procedure, p string) ([]procedure, error) {
	split := strings.Split(p, " ")
	q, err := strconv.Atoi(split[1])
	if err != nil {
		return procedures, err
	}
	newP := procedure{quantity: q}
}

func parseInput(path string) ([][]rune, []procedure, error) {
	file, err := os.Open(path)
	if err != nil {
		return [][]rune{}, []procedure{}, err
	}
	defer file.Close()

	r := bufio.NewReader(file)

	stacks := [][]rune{}
	procedures := []procedure{}
	for i := 0; i >= 0; i++ {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return [][]rune{}, []procedure{}, err
		}
		line = strings.TrimSuffix(line, "\n")

		if i < stackHeight {
			stacks = buildStacks(stacks, line)
		}

		if strings.HasPrefix(line, "move") {
			procedures = appendProcedure(procedures, line)
		}
	}

	return stacks, []procedure{}, nil
}

func main() {
	stacks, _, _ := parseInput("./input.txt")
	fmt.Println(stacks)
	fmt.Println("finished.")
}
