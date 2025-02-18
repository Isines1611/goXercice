package cmd

import (
	"fmt"

	"github.com/Isines1611/goXercice/goxercice/helper"
	"github.com/spf13/cobra"
)

var HintCmd = &cobra.Command{
	Use:   "hint",
	Short: "Provides a hint for the next pending exercise",
	Run: func(cmd *cobra.Command, args []string) {
		hint, err := helper.GetHint()

		if err != nil {
			fmt.Println("❌ Errors getting the hint:", err)
		} else {
			fmt.Println(hint)
		}
	},
}
