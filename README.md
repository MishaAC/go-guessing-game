# Go Guessing Game CLI

A simple **number guessing game built with Go**, designed as a command-line interface (CLI).

This project was created after completing the **basic fundamentals section of a Go bootcamp**, with the goal of practicing core Go concepts through a small but complete game.

---

## ✨ Features

* Interactive terminal menu
* Three difficulty levels (Easy, Medium, Hard)
* Randomly generated secret number
* Limited number of attempts per game
* Clear feedback for each guess (too high, too low, correct)
* Robust input handling using `bufio.Reader`

---

## 🧠 Concepts Applied

This project focuses on **core Go fundamentals**, including:

* Packages and project structure
* Structs and methods
* Constructors (`New` pattern)
* Enums using `iota`
* Game state management
* Separation of concerns (UI vs business logic)
* CLI input handling with `bufio.Reader`

---

## 📁 Project Structure

```
go-guessing-game/
├── main.go            # CLI entry point and UI
├── game/              # Core game logic
│   └── game.go
└── README.md
```

---

## 🚀 Getting Started

### Prerequisites

* Go 1.22 or higher

### Run the application

```bash
go run main.go
```

---

## 🖥️ CLI Preview

```
====================================
         Guessing Game (Go)
====================================
1) Easy     (1 - 20,  10 attempts)
2) Medium   (1 - 50,  7 attempts)
3) Hard     (1 - 100, 5 attempts)
4) Exit

Choose an option:
```

---

## 📝 Notes

* The game logic is fully isolated from the CLI interface.
* Difficulty settings are handled in the UI layer (`main`).
* The project prioritizes clarity, simplicity, and learning.
* Possible improvements include adding tests, score tracking, or replay options.

---

## 👤 Author

**Cristhian**

Built as part of a Go learning journey 🚀
