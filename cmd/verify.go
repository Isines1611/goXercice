package cmd

import (
	"goxercice/helper"
	"path/filepath"

	"github.com/spf13/cobra"
)

var VerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the pending exercice",
	Run: func(cmd *cobra.Command, args []string) {
		config := helper.LoadConfig()
		currExerciceName, _ := helper.GetNextExercise()
		currExercicePath := filepath.Join(config.Path, "exercices", currExerciceName)

		helper.CorrectExercice(currExercicePath, true, true)

	},
}
