package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	Long:  "List all users",
	Run: func(cmd *cobra.Command, args []string) {
		users, err := store.Users()
		if err != nil {
			log.Fatal(err)
		}

		for _, user := range users {
			fmt.Println(user.Username)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
