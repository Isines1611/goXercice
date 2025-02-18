package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Isines1611/goXercice/goxercice/helper"
	"github.com/spf13/cobra"
)

var VerifyAllCmd = &cobra.Command{
	Use:   "verifyAll",
	Short: "Verifies all exercices at once",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("❌ Error: Unable to get home directory:", err)
			return
		}

		exercicePath := filepath.Join(homeDir, ".goxercice", "exercices.json")

		exerciceFile, err := os.Open(exercicePath)
		if err != nil {
			return
		}
		defer exerciceFile.Close()

		var data struct {
			Exercices []string `json:"exercices"`
		}

		decoder := json.NewDecoder(exerciceFile)
		if err := decoder.Decode(&data); err != nil {
			return
		}

		count := 0
		config := helper.LoadConfig()

		for _, file := range data.Exercices {
			path := filepath.Join(config.Path, "exercices", file)

			if helper.CorrectExercice(path, false, false) {
				count++
			} else {
				break
			}
		}

		fmt.Println("Result:", fmt.Sprint(count)+"/"+fmt.Sprint(len(data.Exercices)), "\nConsider using `goxercice resume` to continue your learning.")

		// save config
		config = helper.LoadConfig()
		config.Next = count
		config.Hint = 0
		helper.SaveConfig(config)
	},
}
