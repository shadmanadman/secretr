package cmd

import (
	"fmt"
	"os"
	"secretr/internal"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all stored secrets",
	Run: func(cmd *cobra.Command, args []string) {
		secretsPath := internal.SecretsDir()

		files, err := os.ReadDir(secretsPath)
		if err != nil {
			fmt.Println("❌ Error reading secrets:", err)
			return
		}

		if len(files) == 0 {
			fmt.Println("📂 No secrets found.")
			return
		}

		fmt.Println("🔐 Stored secrets:")

		for _, file := range files {
			name := strings.TrimSuffix(file.Name(), ".secret")
			fmt.Println(".", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
