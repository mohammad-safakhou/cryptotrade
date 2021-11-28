package cmd

import (
	"context"
	"cryptotrade/domain/backend/core/usecase"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(restCmd)
}

var restCmd = &cobra.Command{
	Use:   "rest-server",
	Short: "Starting rest server",
	Run: func(cmd *cobra.Command, args []string) {
		//api.StartRestServer()
		r := usecase.NewReceiverHandler()
		err := r.DataReceiver(context.TODO())
		if err != nil {
			log.Println(err)
		}
	},
}
