package cmd

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	rootCmd.AddCommand(receiverCmd)
}

var receiverCmd = &cobra.Command{
	Use:   "receiver",
	Short: "Starting candle receiver",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		// Middleware
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))

		// Routes
		e.POST("/receiver", func(ctx echo.Context) error {
			type receiverModel struct {
				Text string `json:"text"`
			}
			req := new(receiverModel)
			if err := ctx.Bind(req); err != nil {
				return ctx.JSON(http.StatusBadRequest, err)
			}

			fmt.Println(req)
			return ctx.JSON(http.StatusOK, "")
		})

		// Start server
		go func() {
			if err := e.Start(":443"); err != nil && err != http.ErrServerClosed {
				log.Println(err)
				log.Fatal("shutting down the server")
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}

		//dbPostgres, err := utils.PostgresConnection("localhost", "5432", "root", "root", "cryptotrade", "disable")
		//if err != nil {
		//	panic(err)
		//}
		//
		//candleKafka, err := pkg.KafkaConnection("127.0.0.1", "9092", "candles", 0)
		//if err != nil {
		//	panic(err)
		//}
		//candleKafkaHandler := pkg.NewKafkaHandler(candleKafka)
		//candlesRepository := repository.NewCandlesRepository(dbPostgres)
		//
		//receiverHandler := handlers.NewReceiverHandler(nil, "1min", 30, candlesRepository, candleKafkaHandler)
		//
		//quit := make(chan os.Signal, 1)
		//signal.Notify(quit, os.Interrupt)
		//
		//go func(receiverHandler handlers.Receiver) {
		//	receiverHandler.CallRepeater(context.Background(), "BTCUSDT")
		//}(receiverHandler)
		//<-quit
	},
}
