package main

import (
	"transfer_money/database"
	"transfer_money/providers"
)

func main() {
	database.ConnectToDatabase()
	database.MigrateAll()
	providers.Run()
}
