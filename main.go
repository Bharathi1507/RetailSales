package main

import (
	"RetailSales/database"
	"RetailSales/routers"
	"RetailSales/services"
	"RetailSales/utils"
)

func main() {
	//connects to mysql
	database.SqlConnect()

	//create the table
	utils.CreateTables()

	// runs the cronjob for csv data process
	go services.RunCronJob()

	//starts the server
	routers.NewServer().Start()
}
