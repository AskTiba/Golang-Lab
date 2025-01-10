package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Constants for game choices
const (
	rock     = "rock"
	paper    = "paper"
	scissors = "scissors"
)

// GetComputerChoice generates a random choice for the computer.
func GetComputerChoice() string {
	choices := []string{rock, paper, scissors}
	rand.Seed(time.Now().UnixNano())
	return choices[rand.Intn(len(choices))]
}

// DetermineWinner evaluates the winner based on player and computer choices.
func DetermineWinner(playerChoice, computerChoice string) string {
	if playerChoice == computerChoice {
		return "It's a tie!"
	}

	switch playerChoice {
	case rock:
		if computerChoice == scissors {
			return "You win! Rock beats Scissors."
		}
	case paper:
		if computerChoice == rock {
			return "You win! Paper beats Rock."
		}
	case scissors:
		if computerChoice == paper {
			return "You win! Scissors beats Paper."
		}
	}

	return "You lose! Better luck next time."
}

// PlayRound handles the logic for a single round of the game.
func PlayRound(playerChoice string) string {
	playerChoice = strings.ToLower(strings.TrimSpace(playerChoice))
	if !isValidChoice(playerChoice) {
		return "Invalid choice. Please choose 'rock', 'paper', or 'scissors'."
	}

	computerChoice := GetComputerChoice()
	fmt.Printf("Player chose: %s\n", playerChoice)
	fmt.Printf("Computer chose: %s\n", computerChoice)

	return DetermineWinner(playerChoice, computerChoice)
}

// isValidChoice checks if the player's choice is valid.
func isValidChoice(choice string) bool {
	return choice == rock || choice == paper || choice == scissors
}

// Main function to run the game.
func main() {
	fmt.Println("Welcome to Rock, Paper, Scissors!")
	fmt.Println("Type 'rock', 'paper', or 'scissors' to play. Type 'exit' to quit.")

	for {
		// Get player's choice
		var playerChoice string
		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&playerChoice)

		// Exit condition
		if strings.ToLower(strings.TrimSpace(playerChoice)) == "exit" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

		// Play a round and show the result
		result := PlayRound(playerChoice)
		fmt.Println(result)
	}
}
