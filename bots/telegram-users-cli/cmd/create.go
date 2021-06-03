package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mbtamuli/crypto/model"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the user",
	Long:  "Create the user",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if username, err = cmd.Flags().GetString("username"); err != nil {
			return fmt.Errorf("username needs to be set: %s", err.Error())
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		user := &model.User{
			ID:       uuid.New(),
			Username: username,
		}

		if err := store.CreateUser(user); err != nil {
			panic(err)
		}
	},
}

func init() {
	createCmd.Flags().StringP(
		"username",
		"u",
		"",
		"username of the user to create")

	createCmd.MarkFlagRequired("username")
	rootCmd.AddCommand(createCmd)
}
