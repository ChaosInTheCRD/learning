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
	N = 100
)

func initSlice() []int {
	mySlice := make([]int, N)
	for i := 0; i < N; i++ {
		mySlice[i] = 0
	}

	return mySlice
}

func fillSlice(s []int, f int, t int) []int {
	for i := f; i <= t; i++ {
		s[i-1] = 1
	}

	return s
}

func parseInput(input string) ([][]int, error) {
	f, err := os.Open(input)
	if err != nil {
		return [][]int{}, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	scheds := [][]int{}
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return [][]int{}, err
		}

		ass := strings.Split(strings.TrimSuffix(line, "\n"), ",")
		for _, a := range ass {
			slice := initSlice()
			ft := strings.Split(a, "-")
			f, err := strconv.Atoi(ft[0])
			if err != nil {
				return [][]int{}, err
			}
			t, err := strconv.Atoi(ft[1])
			if err != nil {
				return [][]int{}, err
			}
			slice = fillSlice(slice, f, t)
			scheds = append(scheds, slice)
		}
	}

	return scheds, nil
}

func isEncased(sub, targ []int) bool {
	for i, n := range sub {
		if n == 0 {
			continue
		} else if n == 1 && targ[i] == 1 {
			continue
		} else {
			return false
		}
	}

	return true
}

func translateBack(scheds [][]int) [][]int {
	in := false
	outs := [][]int{}
	for _, sched := range scheds {
		out := []int{}
		for j := range sched {
			if sched[j] == 1 && !in {
				fmt.Println(j + 1)
				out = append(out, j+1)
				in = true
			} else if sched[j] == 0 && in {
				fmt.Println(j)
				out = append(out, j)
				in = false
			}
		}
		outs = append(outs, out)
	}
	return outs
}

func main() {
	scheds, err := parseInput("./input.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	for i := 0; i < len(scheds); i = i + 2 {
		if isEncased(scheds[i], scheds[i+1]) {
			total = total + 1
		} else if isEncased(scheds[i+1], scheds[i]) {
			total = total + 1
		}
		continue
	}

	fmt.Println(total)
	fmt.Println("finished.")
}
