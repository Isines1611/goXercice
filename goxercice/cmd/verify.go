package cmd

import (
	"path/filepath"

	"github.com/Isines1611/goXercice/goxercice/helper"
	"github.com/spf13/cobra"
)

var VerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the pending exercice",
	Run: func(cmd *cobra.Command, args []string) {
		config := helper.LoadConfig()
		currExerciceName, _ := helper.GetNextExercise(-1)
		currExercicePath := filepath.Join(config.Path, "exercices", currExerciceName)

		helper.CorrectExercice(currExercicePath, true, true)

	},
}
