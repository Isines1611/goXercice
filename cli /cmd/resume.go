package cmd

import (
	"goxercice/helper"

	"github.com/spf13/cobra"
)

var ResumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resumes progress and enters CLI mode",
	Run: func(cmd *cobra.Command, args []string) {
		helper.CLIMode()
	},
}
