package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	RANGE_START  = 1
	RANGE_END    = 100
	MAX_ATTEMPTS = 3
)

func generateRandomNumber(rangeStart, rangeEnd int) int {
	return rand.Intn(rangeEnd-rangeStart+1) + rangeStart
}

func getUserGuess(rangeStart, rangeEnd int) *int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Guess a number between %d and %d: ", rangeStart, rangeEnd)
	guessStr, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading input. Please try again.")
		return nil
	}

	guessStr = strings.TrimSpace(guessStr)
	guess, err := strconv.Atoi(guessStr)
	if err != nil {
		log.Println("Invalid input! Please enter an integer.")
		return nil
	}

	if guess < rangeStart || guess > rangeEnd {
		log.Printf("Please enter a number within the range %d-%d.\n", rangeStart, rangeEnd)
		return nil
	}

	return &guess
}

func processGuess(guess, target int) string {
	if guess < target {
		return "Too low."
	} else if guess > target {
		return "Too high."
	} else {
		return "Correct!"
	}
}

func playGame() {
	log.Println("Welcome to the Guess the Number Game!")
	targetNumber := generateRandomNumber(RANGE_START, RANGE_END)
	attempts := 0

	for {
		guess := getUserGuess(RANGE_START, RANGE_END)
		if guess == nil {
			continue
		}

		attempts++
		response := processGuess(*guess, targetNumber)
		log.Println(response)

		if response == "Correct!" {
			log.Printf("Congratulations! You guessed it in %d attempts.\n", attempts)
			break
		}
		if MAX_ATTEMPTS > 0 && attempts >= MAX_ATTEMPTS {
			log.Printf("Sorry, you've reached the maximum number of attempts. The number was %d.\n", targetNumber)
			break
		}
	}
}

func main() {
	playGame()
}
