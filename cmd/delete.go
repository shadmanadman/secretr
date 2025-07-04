package cmd

import (
	"fmt"
	"secretr/internal"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a stored secret",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.DeleteSecret(args[0])
		if err != nil {
			fmt.Println("❌", err)
		} else {
			fmt.Println("🗑️ Secret deleted:", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
