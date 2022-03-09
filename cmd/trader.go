package cmd

import (
	"context"
	"cryptotrade/handlers"
	"cryptotrade/pkg"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

func init() {
	rootCmd.AddCommand(traderCmd)
}

var traderCmd = &cobra.Command{
	Use:   "trader",
	Short: "Starting candle receiver",
	Run: func(cmd *cobra.Command, args []string) {
		candleKafka, err := pkg.KafkaConnection("127.0.0.1", "9092", "candles", 0)
		if err != nil {
			panic(err)
		}
		candleKafkaHandler := pkg.NewKafkaHandler(candleKafka)

		traderHandler := handlers.NewTraderHandler(candleKafkaHandler)

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)

		go func(traderHandler handlers.Trader) {
			err := traderHandler.Handler(context.Background())
			if err != nil {
				panic(err)
			}
		}(traderHandler)
		<-quit
	},
}
