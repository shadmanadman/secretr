package cmd

import (
	"bufio"
	"fmt"
	"os"
	"secretr/internal"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "Retrieve a stored secret",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter passphrase: ")
		pass, _ := reader.ReadString('\n')

		secret, err := internal.RetrieveSecret(args[0], pass)
		if err != nil {
			fmt.Println("❌ Error:", err)
		} else {
			fmt.Println("🔐 Secret:", secret)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
