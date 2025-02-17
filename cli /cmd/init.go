package cmd

import (
	"fmt"
	"goxercice/helper"
	"os"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize goxercice and download exercises",
	Long:  `Sets up goxercice in the user's home directory and downloads exercises. If a path is provided, exercises will be stored there. Otherwise, they are stored in the current directory.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📂 Initializing goxercice...")

		var exercisesPath string

		if len(args) > 0 {
			exercisesPath = args[0]
			if _, err := os.Stat(exercisesPath); os.IsNotExist(err) {
				fmt.Println("❌ Invalid path:", err)
				return
			}
		} else {
			exercisesPath, _ = os.Getwd() // use current directory
		}

		if !helper.Initialize(exercisesPath) {
			return
		}

		// ensures the directory exists and clones or updates the exercises repository
		if err := os.MkdirAll(exercisesPath, 0755); err != nil {
			return
		}

		fmt.Println("📂 Storing exercises in:", exercisesPath)

		// Create the directory if it doesn’t exist
		if _, err := os.Stat(exercisesPath); os.IsNotExist(err) {
			os.MkdirAll(exercisesPath, os.ModePerm)
		}

		fmt.Println("⬇️ Cloning exercises repository...")

		// USE GITHUB API

		helper.DownloadFiles(exercisesPath)

		fmt.Println("✅ Initialization complete! Start with your first exercise.")

		fmt.Print("▶️ Enter to start goXercice...")
		var response string = ""
		fmt.Scanln(&response)

		helper.CLIMode()
	},
}
