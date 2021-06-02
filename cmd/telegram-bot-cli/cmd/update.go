package cmd

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the user",
	Long:  "Update the user",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if username, err = cmd.Flags().GetString("username"); err != nil {
			return fmt.Errorf("username needs to be set: %s", err.Error())
		}
		if chat_id, err = cmd.Flags().GetString("chat_id"); err != nil {
			return fmt.Errorf("chat_id needs to be set: %s", err.Error())
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		existing_user, err := store.UserByUsername(username)
		if err != nil {
			fmt.Printf("username %s does not exist", username)
		}
		existing_user.ChatID = sql.NullString{String: chat_id, Valid: true}
		if err := store.UpdateUser(&existing_user); err != nil {
			panic(err)
		}
	},
}

func init() {
	updateCmd.Flags().StringP(
		"username",
		"u",
		"",
		"username of the user to update")
	updateCmd.Flags().StringP(
		"chat_id",
		"c",
		"",
		"chat_id of the user")

	updateCmd.MarkFlagRequired("username")
	updateCmd.MarkFlagRequired("chat_id")
	rootCmd.AddCommand(updateCmd)
}
