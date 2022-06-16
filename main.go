package main

import (
	"context"
	"cryptotrade/handlers"
	"fmt"
	kucoin "github.com/Kucoin/kucoin-futures-go-sdk"
)

//go:generate sqlboiler --wipe --no-tests psql -o models

func main() {
	s := kucoin.NewApiService(
		kucoin.ApiBaseURIOption("https://api-sandbox-futures.kucoin.com"),
		kucoin.ApiKeyOption("62aba38329c69200011e7f5d"),
		kucoin.ApiSecretOption("e1fbb51d-b338-4427-8ebf-7bdba7da8c6f"),
		kucoin.ApiPassPhraseOption("220618"),
		kucoin.ApiKeyVersionOption("2"),
	)

	trader := handlers.NewTraderHandler(s)
	receiver := handlers.NewReceiverHandler(s, trader)

	err := receiver.Handler(context.TODO(), "Short 1m")
	if err != nil {
		fmt.Println(err.Error())
	}
	//cmd.Execute()
}
