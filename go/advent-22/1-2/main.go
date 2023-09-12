package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func generateSlice(file string) ([]int, error) {
	f, err := os.Open(file)
	if err != nil {
		return []int{}, errors.Join(err, fmt.Errorf("failed to read file"))
	}
	defer f.Close()

	r := bufio.NewReader(f)

	out := []int{}
	for {
		line, err := r.ReadString(10)
		if line == "\n" {
			out = append(out, 0)
			continue
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return []int{}, errors.Join(err, fmt.Errorf("failed to read line of file"))
		}

		i, err := strconv.Atoi(strings.TrimSuffix(line, "\n"))
		if err != nil {
			return []int{}, errors.Join(err, fmt.Errorf("failed to convert line %s to integer", line))
		}

		out = append(out, i)
	}
	return out, nil
}

func findCalsPerElf(cals []int) []int {
	out := []int{}
	c := 0
	for _, cal := range cals {
		if cal == 0 {
			out = append(out, c)
			c = 0
			continue
		}

		c = c + cal
	}

	out = sortElvesByCal(out)

	return out
}

func sortElvesByCal(elves []int) []int {
	sort.Ints(elves)
	return elves
}

func main() {
	cals, err := generateSlice("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	calsPerElf := findCalsPerElf(cals)
	if err != nil {
		fmt.Println(err)
		return
	}

	topThreeTotal := 0
	for i := 1; i <= 3; i++ {
		topThreeTotal = topThreeTotal + calsPerElf[len(calsPerElf)-i]
	}

	fmt.Println(topThreeTotal)
}
