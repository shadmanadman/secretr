package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"secretr/internal"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a stored secret",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		secretPath := filepath.Join(internal.SecretsDir(), args[0]+".secret")

		if _, err := os.Stat(secretPath); os.IsNotExist(err) {
			fmt.Println("❌ Secret not found.")
			return
		}

		err := os.Remove(secretPath)
		if err != nil {
			fmt.Println("❌ Failed to delete secret:", err)
		} else {
			fmt.Println("🗑️ Secret deleted:", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
