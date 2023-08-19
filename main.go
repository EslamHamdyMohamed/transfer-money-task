package main

import (
	"github.com/subosito/gotenv"
	"transfer_money/database"
	"transfer_money/providers"
)

func main() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}
	database.ConnectToDatabase()
	database.MigrateAll()
	providers.Run()
}
