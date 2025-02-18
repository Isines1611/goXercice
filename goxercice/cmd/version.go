package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Provides the current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("goXercice Version 1.0")
	},
}
