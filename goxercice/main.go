package main

import (
	"fmt"
	"os"

	"github.com/Isines1611/goXercice/goxercice/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goxercice",
	Short: "GoXercice is an interactive Go learning tool",
	Long:  "GoXercice helps you learn Go by completing exercices and verifying your solutions.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to GoXercice! Type 'goxercice help' for available commands")
		cmd.Help()
	},
}

func main() {
	// Register subcommands
	rootCmd.AddCommand(cmd.InitCmd, cmd.VerifyCmd, cmd.VerifyAllCmd, cmd.HintCmd, cmd.ResetCmd, cmd.ResumeCmd, cmd.VersionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
