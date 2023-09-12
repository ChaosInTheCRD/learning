package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3
)

func transformInput(input string) ([]string, []string, error) {
	f, err := os.Open(input)
	if err != nil {
		return []string{}, []string{}, errors.Join(err, fmt.Errorf("failed to open file"))
	}
	defer f.Close()

	r := bufio.NewReader(f)

	opp := []string{}
	resp := []string{}

	for {
		line, err := r.ReadString(10)
		if line == "\n" {
			continue
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return []string{}, []string{}, errors.Join(err, fmt.Errorf("failed to read line of file"))
		}

		line = strings.TrimSuffix(line, "\n")
		l := strings.Split(line, " ")
		switch l[0] {
		case "A":
			opp = append(opp, "rock")
		case "B":
			opp = append(opp, "paper")
		case "C":
			opp = append(opp, "scissors")
		default:
			return []string{}, []string{}, fmt.Errorf("failed to decrypt character %s in input", l[0])
		}

		switch l[1] {
		case "X":
			resp = append(resp, "rock")
		case "Y":
			resp = append(resp, "paper")
		case "Z":
			resp = append(resp, "scissors")
		default:
			return []string{}, []string{}, fmt.Errorf("failed to decrypt character %s in input", l[1])
		}
	}

	return opp, resp, nil
}

func calcScore(opp, resp string) int {
	switch resp {
	case "rock":
		if opp == "scissors" {
			return 6 + getBonus(resp)
		} else {
			return getBonus(resp)
		}
	case "paper":
		if opp == "rock" {
			return 6 + getBonus(resp)
		} else {
			return getBonus(resp)
		}
	case "scissors":
		if opp == "paper" {
			return 6 + getBonus(resp)
		} else {
			return getBonus(resp)
		}
	default:
		return 0
	}
}

func getBonus(s string) int {
	switch s {
	case "rock":
		return rock
	case "paper":
		return paper
	case "scissors":
		return scissors
	default:
		return 0
	}
}

func totalScore(opp, resp []string) (int, error) {
	if len(opp) != len(resp) {
		return 0, fmt.Errorf("Provided opponent input is of length %d, but provided response input is of length %d", len(opp), len(resp))
	}

	total := 0

	for i := range resp {
		fmt.Println(resp[i], "vs", opp[i])
		if resp[i] == opp[i] {
			fmt.Println(3 + getBonus(resp[i]))
			total = total + 3 + getBonus(resp[i])
			continue
		}
		fmt.Println(calcScore(opp[i], resp[i]))
		total = total + calcScore(opp[i], resp[i])
	}

	return total, nil
}

func main() {
	opp, resp, err := transformInput("./input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("opp:", opp)
	fmt.Println("resp:", resp)

	score, err := totalScore(opp, resp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Total Score:", score)
}
