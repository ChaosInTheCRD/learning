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
			resp = append(resp, "lose")
		case "Y":
			resp = append(resp, "draw")
		case "Z":
			resp = append(resp, "win")
		default:
			return []string{}, []string{}, fmt.Errorf("failed to decrypt character %s in input", l[1])
		}
	}

	return opp, resp, nil
}

func calcScore(opp, resp string) int {
	switch opp {
	case "rock":
		if resp == "win" {
			fmt.Println("picking paper to beat rock, meaning a score of", 6+paper, "for this round.")
			return 6 + paper
		} else {
			fmt.Println("picking scissors to lose to rock, meaning a score of", scissors, "for this round.")
			return scissors
		}
	case "paper":
		if resp == "win" {
			fmt.Println("picking scissors to beat paper, meaning a score of", 6+scissors, "for this round.")
			return 6 + scissors
		} else {
			fmt.Println("picking rock to lose to paper, meaning a score of", rock, "for this round.")
			return rock
		}
	case "scissors":
		if resp == "win" {
			fmt.Println("picking rock to beat scissors, meaning a score of", 6+rock, "for this round.")
			return 6 + rock
		} else {
			fmt.Println("picking paper to lose to scissors, meaning a score of", paper, "for this round.")
			return paper
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
		fmt.Println("you need to", resp[i], "this round.")
		if resp[i] == "draw" {
			fmt.Println("adding 3 to total, plus", getBonus(opp[i]), "as a bonus for the", opp[i], ".")
			total = total + 3 + getBonus(opp[i])
			continue
		} else {
			total = total + calcScore(opp[i], resp[i])
		}
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
