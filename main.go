package main

import (
	"context"
	"cryptotrade/cmd"
	"cryptotrade/handlers"
	"time"
)

func main() {
	n := handlers.NewReceiverHandler(nil, "1min")
	n.CallRepeater(context.TODO(), "BTCUSDT")
	time.Sleep(1000*time.Second)
	cmd.Execute()
}
