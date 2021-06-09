package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ApiKey    string `split_words:"true"`
	SecretKey string `split_words:"true"`
}

var s Config

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

// initConfig reads ENV variables if set.
func initConfig() {
	err := envconfig.Process("bbot", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
}
