package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the binance-bot",
	Long:  "Start the binance-bot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running the start command")
		fmt.Println(s.ApiKey)
		fmt.Println(s.SecretKey)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
