package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MishaAC/go-guessing-game/game"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	for {
		printMenu()

		option, err := readInt()
		if err != nil {
			fmt.Println("Invalid option. Please enter a number between 1 and 4.")
			continue
		}

		switch option {
		case 1:
			play(1, 20, 10)
		case 2:
			play(1, 50, 7)
		case 3:
			play(1, 100, 5)
		case 4:
			fmt.Println("Goodbye! 👋")
			return
		default:
			fmt.Println("Option must be between 1 and 4.")
		}
	}
}

func printMenu() {
	fmt.Println("====================================")
	fmt.Println("         Guessing Game (Go)         ")
	fmt.Println("====================================")
	fmt.Println("1) Easy     (1 - 20,  10 attempts)")
	fmt.Println("2) Medium   (1 - 50,  7 attempts)")
	fmt.Println("3) Hard     (1 - 100, 5 attempts)")
	fmt.Println("4) Exit")
	fmt.Println()
	fmt.Print("Choose an option: ")
}

func play(min, max, attempts int) {
	g := game.New(min, max, attempts)

	fmt.Println()
	fmt.Printf("I'm thinking of a number between %d and %d.\n", min, max)

	for {
		fmt.Printf("Attempts left: %d\n", g.AttemptsLeft())
		fmt.Print("Enter your guess: ")

		number, err := readInt()
		if err != nil {
			fmt.Printf("Invalid input. Please enter a number between %d and %d.\n", min, max)
			continue
		}

		result := g.Guess(number)

		switch result {
		case game.TooLow:
			fmt.Println("Too low ⬇️")
		case game.TooHigh:
			fmt.Println("Too high ⬆️")
		case game.Correct:
			fmt.Println("Correct! 🎉 You won!")
			return
		case game.GameOver:
			fmt.Println("Game over 💀 No attempts left.")
			return
		}

		fmt.Println()
	}
}

func readInt() (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}
