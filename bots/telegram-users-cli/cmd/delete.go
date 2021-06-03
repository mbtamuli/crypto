package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var allFlag bool

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the user",
	Long:  "Delete the user",
	Args: func(cmd *cobra.Command, args []string) error {
		var err error
		if allFlag, err = cmd.Flags().GetBool("all"); err != nil {
			return nil
		}
		if username, err = cmd.Flags().GetString("username"); err != nil {
			return fmt.Errorf("username needs to be set: %s", err.Error())
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if allFlag {
			if err := store.DeleteUsers(); err != nil {
				panic(err)
			}
			return
		}
		existing_user, err := store.UserByUsername(username)
		if err != nil {
			fmt.Print(err)
			fmt.Printf("username %s does not exist", username)
			os.Exit(1)
		}
		fmt.Println(existing_user)
		if err := store.DeleteUser(existing_user.ID); err != nil {
			panic(err)
		}
	},
}

func init() {
	deleteCmd.Flags().StringP(
		"username",
		"u",
		"",
		"username of the user to delete")
	deleteCmd.Flags().BoolP(
		"all",
		"a",
		false,
		"delete all users")

	rootCmd.AddCommand(deleteCmd)
}
