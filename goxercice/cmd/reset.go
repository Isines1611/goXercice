package cmd

import (
	"fmt"

	"github.com/Isines1611/goXercice/goxercice/helper"
	"github.com/spf13/cobra"
)

var ResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets all progress (confirmation required)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("⚠️ Are you sure you want to reset all progress? (yes/no): ")
		var response string
		fmt.Scanln(&response)

		if response == "yes" || response == "y" {
			helper.ResetFiles()
			fmt.Println("🔄 Progress reset! Run 'goxercice init' to start again.")
		} else {
			fmt.Println("❌ Reset cancelled.")
		}
	},
}
