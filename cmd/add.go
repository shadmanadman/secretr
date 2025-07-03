package cmd

import (
	"bufio"
	"fmt"
	"os"
	"secretr/internal"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a new secret",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter secret: ")
		secret, _ := reader.ReadString('\n')
		fmt.Print("Enter passphrase: ")
		pass, _ := reader.ReadString('\n')

		err := internal.StoreSecret(args[0], secret, pass)
		if err != nil {
			fmt.Println("❌ Error:", err)
		} else {
			fmt.Println("✅ Secret stored.")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
