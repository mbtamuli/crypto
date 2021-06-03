package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "binance-bot",
	Short: "A bot that fetches data from Binance and sends information via Telegram",
	Long: `A bot that uses the Binance API to fetch real-time data and sends it
to users registered in the user database via the telegram-users-cli.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix("bbot")
	viper.AutomaticEnv() // read in environment variables that match
}
