package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var myScore int64 = 0

	for fileScanner.Scan() {
		plays := strings.Split(fileScanner.Text(), " ")

		opponentPlay := playName(plays[0])
		mePlay := playName(nextPlay(opponentPlay, plays[1]))
		winner := ""

		// We always get a score for whatever we are chosing
		myScore += getScore(mePlay)

		if isWinningPlay(mePlay, opponentPlay) {
			winner = "me"
			myScore += 6
		} else if isWinningPlay(opponentPlay, mePlay) {
			winner = "opponent"
		} else {
			winner = ""
			myScore += 3
		}

		fmt.Printf("Play: %v %v - winner %v. Score: %v\n", mePlay, opponentPlay, winner, myScore)

	}

	fmt.Printf("My score: %v\n", myScore)

	readFile.Close()
}

func getScore(play string) int64 {
	if play == "rock" {
		return 1
	} else if play == "paper" {
		return 2
	} else if play == "scissors" {
		return 3
	}

	return 0
}

func nextPlay(opponent string, outcomeCode string) string {
	outcome := ""
	if outcomeCode == "X" {
		outcome = "lose"
	} else if outcomeCode == "Y" {
		outcome = "draw"
	} else {
		outcome = "win"
	}

	const ROCK = "A"
	const PAPER = "B"
	const SCISSORS = "C"
	if opponent == "rock" {
		if outcome == "lose" {
			return SCISSORS
		} else if outcome == "win" {
			return PAPER
		} else {
			return ROCK
		}
	} else if opponent == "paper" {
		if outcome == "lose" {
			return ROCK
		} else if outcome == "win" {
			return SCISSORS
		} else {
			return PAPER
		}
	} else if opponent == "scissors" {
		if outcome == "lose" {
			return PAPER
		} else if outcome == "win" {
			return ROCK
		} else {
			return SCISSORS
		}
	}

	return ""
}

func isWinningPlay(p1 string, p2 string) bool {
	if p1 == p2 {
		return false
	} else if p1 == "rock" && p2 != "paper" {
		return true
	} else if p1 == "paper" && p2 != "scissors" {
		return true
	} else if p1 == "scissors" && p2 != "rock" {
		return true
	}

	return false
}

func playName(play string) string {
	var ROCK = [2]string{"A", "X"}
	var PAPER = [2]string{"B", "Y"}
	var SCISSORS = [2]string{"C", "Z"}

	if play == ROCK[0] || play == ROCK[1] {
		return "rock"
	} else if play == PAPER[0] || play == PAPER[1] {
		return "paper"
	} else if play == SCISSORS[0] || play == SCISSORS[1] {
		return "scissors"
	}

	return ""
}
