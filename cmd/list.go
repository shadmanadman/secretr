package cmd

import (
	"fmt"
	"secretr/internal"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all stored secrets",
	Run: func(cmd *cobra.Command, args []string) {
		names, err := internal.ListSecrets()
		if err != nil {
			fmt.Println("❌", err)
			return
		}
		if len(names) == 0 {
			fmt.Println("📂 No secrets found.")
			return
		}
		fmt.Println("🔐 Stored secrets:")
		for _, name := range names {
			fmt.Println("•", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
