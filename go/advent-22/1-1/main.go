package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func mostCalsElf(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, errors.Join(err, fmt.Errorf("failed to read file"))
	}
	defer f.Close()

	r := bufio.NewReader(f)

	out := 0
finish:
	for {
		elfTotal := 0
		for {
			line, err := r.ReadString(10)
			if line == "\n" {
				break
			}
			if err == io.EOF {
				break finish
			} else if err != nil {
				return 0, errors.Join(err, fmt.Errorf("failed to read line of file"))
			}

			i, err := strconv.Atoi(strings.TrimSuffix(line, "\n"))
			if err != nil {
				return 0, errors.Join(err, fmt.Errorf("failed to convert line %s to integer", line))
			}
			elfTotal = elfTotal + i
		}
		if elfTotal > out {
			out = elfTotal
		}
	}
	return out, nil
}

func main() {
	out, err := mostCalsElf("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("most cals for one elf is:", out)
}
