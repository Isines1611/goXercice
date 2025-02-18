package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "hint",
	Short: "Provides a hint for the next pending exercise",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("goXercice Version 1.0")
	},
}
