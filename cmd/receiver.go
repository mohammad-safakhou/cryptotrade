package cmd

import (
	"context"
	"cryptotrade/handlers"
	"cryptotrade/models"
	"cryptotrade/utils"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	rootCmd.AddCommand(receiverCmd)
	receiverCmd.Flags().String("server-port", "10000", "Setting http server port")
}

var receiverCmd = &cobra.Command{
	Use:   "receiver",
	Short: "Starting candle receiver",
	Run: func(cmd *cobra.Command, args []string) {
		serverPort, err := cmd.Flags().GetString("server-port")

		e := echo.New()

		// Middleware
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))

		handlers.SharedPostgresDB, err = utils.PostgresConnection("localhost", "5432", "root", "root", "cryptotrade", "disable")
		if err != nil {
			panic(err)
		}

		strategy, err := models.Strategies(models.StrategyWhere.Name.EQ(null.NewString("main", true))).One(context.TODO(), handlers.SharedPostgresDB)
		if err != nil {
			panic(err)
		}

		var strat handlers.Strategy
		err = json.Unmarshal([]byte(strategy.Data.String), &strat)
		if err != nil {
			panic(err)
		}
		strat.MainTimeFrame.Storage = &handlers.Storage{}
		if strat.SubTimeFrames != nil {
			for _, value := range strat.SubTimeFrames {
				value.Storage = &handlers.Storage{}
			}
		}

		spew.Config.Indent = "\t"
		spew.Dump("project starting with strategy: ", strat)

		handlers.SharedObject = &handlers.Object{
			Exit:     false,
			Strategy: &strat,
			Action:   make(chan *handlers.Action, 100),
		}

		go handlers.SharedObject.ActionHandler()

		// Routes
		e.POST("/receiver", func(ctx echo.Context) error {
			log.Println("new signal coming ---------------------------------------------------------------------------------------------------")
			bodyBytes, _ := ioutil.ReadAll(ctx.Request().Body)

			go func(bodyBytes []byte) {
				content := models.Content{Data: null.NewString(string(bodyBytes), true)}
				err := content.Insert(ctx.Request().Context(), handlers.SharedPostgresDB, boil.Infer())
				if err != nil {
				}
			}(bodyBytes)

			var signal handlers.Signals
			err = json.Unmarshal(bodyBytes, &signal)
			if err != nil {
				log.Println(err.Error())
				return err
			}
			signal.ReceivedTime = time.Now()
			handlers.SharedObject.ReceiveSignal(&signal)

			return ctx.JSON(http.StatusOK, "")
		})
		e.GET("/exit", func(ctx echo.Context) error {
			handlers.SharedObject.StopStrategy()
			return ctx.JSON(http.StatusOK, "successfully stopped the bot, you can resume it by calling resume route")
		})
		e.GET("/resume", func(ctx echo.Context) error {
			handlers.SharedObject.ResumeStrategy()
			return ctx.JSON(http.StatusOK, "resumed the bot")
		})
		e.GET("/object/details", func(ctx echo.Context) error {
			return ctx.JSON(http.StatusOK, handlers.SharedObject)
		})

		// Start server
		go func() {
			if err := e.Start(":" + serverPort); err != nil && err != http.ErrServerClosed {
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
	},
}
