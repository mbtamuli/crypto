package cmd

import (
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the binance-bot",
	Long:  "Start the binance-bot",
	Run: func(cmd *cobra.Command, args []string) {
		binance.UseTestnet = true
		// _ = binance.NewClient(s.ApiKey, s.SecretKey)
		wsDepthHandler := func(event *binance.WsMarketStatEvent) {
			spew.Dump(event)
		}
		errHandler := func(err error) {
			fmt.Println(err)
		}
		doneC, stopC, err := binance.WsMarketStatServe("BTCUSDT", wsDepthHandler, errHandler)
		if err != nil {
			fmt.Println(err)
			return
		}
		// use stopC to exit
		go func() {
			time.Sleep(5 * time.Second)
			stopC <- struct{}{}
		}()
		// remove this if you do not want to be blocked here
		<-doneC
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
