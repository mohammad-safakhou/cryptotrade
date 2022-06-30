package main

import (
	"cryptotrade/cmd"
)

//go:generate sqlboiler --wipe --no-tests psql -o models

func main() {
	//dbPostgres, err := utils.PostgresConnection("localhost", "5432", "root", "root", "cryptotrade", "disable")
	//if err != nil {
	//	panic(err)
	//}
	//strategy, err := models.Strategies(models.StrategyWhere.Name.EQ(null.NewString("main", true))).One(context.TODO(), dbPostgres)
	//if err != nil {
	//	panic(err)
	//}
	//
	//var strat handlers.Strategy
	//err = json.Unmarshal([]byte(strategy.Data.String), &strat)
	//if err != nil {
	//	panic(err)
	//}
	//handlers.SharedObject = &handlers.Object{
	//	Exit:     false,
	//	Strategy: &strat,
	//	Action:   make(chan *handlers.Action, 100),
	//}
	//handlers.SharedObject.ClosePosition()
	//handlers.SharedObject.OpenPosition("buy")
	//position := handlers.SharedObject.GetOpenPosition()
	//fmt.Println(position)
	cmd.Execute()
}
