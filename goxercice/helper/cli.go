package helper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func CLIMode() {
	ClearScreen()
	config := LoadConfig()
	currExerciceName, _ := GetNextExercise(-1)
	currExercicePath := filepath.Join(config.Path, "exercices", currExerciceName)

	fmt.Println("\n🎯 GoXercice CLI Mode")

	doneChan := make(chan bool)
	nextChan := make(chan bool)

	go startWatch(currExercicePath, doneChan, nextChan)

	reader := bufio.NewReader(os.Stdin)

	for {
		currExerciceName, _ := GetNextExercise(-1)
		currExercicePath := filepath.Join(config.Path, "exercices", currExerciceName)

		fmt.Println("📌 Current Exercise:", currExerciceName)
		fmt.Println("🔹 Commands: (v) Verify | (h) Hint | (q) Quit")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "v":
			if CorrectExercice(currExercicePath, true, true) {
				ClearScreen()

				if next := CheckForNext(); next != "" {
					go startWatch(next, doneChan, nextChan) // watch the new file
				}
			}
		case "verify":
			if CorrectExercice(currExercicePath, true, true) {
				ClearScreen()

				if next := CheckForNext(); next != "" {
					go startWatch(next, doneChan, nextChan) // watch the new file
				}
			}
		case "n":
			ClearScreen()
		case "next":
			ClearScreen()
		case "h":
			hint, err := GetHint()

			if err != nil {
				fmt.Println("❌ Errors getting the hint:", err)
			} else {
				fmt.Println(hint)
			}
		case "hint":
			hint, err := GetHint()

			if err != nil {
				fmt.Println("❌ Errors getting the hint:", err)
			} else {
				fmt.Println(hint)
			}
		case "q":
			fmt.Println("👋 Exiting GoXercice. See you next time!")
			return
		case "quit":
			fmt.Println("👋 Exiting GoXercice. See you next time!")
			return
		default:
			fmt.Println("❌ Unknown command! Try: (v) Verify | (h) Hint | (q) Quit")
		}

		select {
		case <-nextChan:
			if next := CheckForNext(); next != "" {
				go startWatch(next, nextChan, nextChan) // watch the new file
			} else {
				return
			}
		default:
		}
	}

}

func CheckForNext() string {
	config := LoadConfig()
	currExerciceName, _ := GetNextExercise(-1)

	if config.Next == countExercices() {
		fmt.Println("🎉 All exercices completed!")
		return ""
	}

	fmt.Println("➡️ Next Exercise:", currExerciceName)
	return filepath.Join(config.Path, "exercices", currExerciceName)
}

// Clear the terminal to keep everything clean
func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[H\033[2J")
	}
}

func countExercices() int {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Error: Unable to get home directory:", err)
		return -1
	}

	exercicePath := filepath.Join(homeDir, ".goxercice", "exercices.json")

	exerciceFile, err := os.Open(exercicePath)
	if err != nil {
		return -1
	}
	defer exerciceFile.Close()

	var data struct {
		Exercices []string `json:"exercices"`
	}

	decoder := json.NewDecoder(exerciceFile)
	if err := decoder.Decode(&data); err != nil {
		return -1
	}

	return len(data.Exercices)
}
