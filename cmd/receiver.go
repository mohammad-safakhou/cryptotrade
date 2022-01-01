package cmd

import (
	"context"
	"cryptotrade/handlers"
	"cryptotrade/pkg"
	"cryptotrade/repository"
	"cryptotrade/utils"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

func init() {
	rootCmd.AddCommand(receiverCmd)
}

var receiverCmd = &cobra.Command{
	Use:   "receiver",
	Short: "Starting candle receiver",
	Run: func(cmd *cobra.Command, args []string) {
		dbPostgres, err := utils.PostgresConnection("localhost", "5432", "root", "root", "cryptotrade", "disable")
		if err != nil {
			panic(err)
		}

		candleKafka, err := pkg.KafkaConnection("127.0.0.1", "9092", "candles", 0)
		if err != nil {
			panic(err)
		}
		candleKafkaHandler := pkg.NewKafkaHandler(candleKafka)
		candlesRepository := repository.NewCandlesRepository(dbPostgres)

		receiverHandler := handlers.NewReceiverHandler(nil, "1min", 30, candlesRepository, candleKafkaHandler)

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)

		go func(receiverHandler handlers.Receiver) {
			receiverHandler.CallRepeater(context.Background(), "BTCUSDT")
		}(receiverHandler)
		<-quit
	},
}
