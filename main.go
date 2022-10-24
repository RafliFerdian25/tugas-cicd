package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/databases"
	"belajar-go-echo/routes"
)

func main() {
	config.InitConfig()
	db, err := databases.ConnectDB()
	if err != nil {
		panic(err)
	}
	err = databases.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := routes.New(db)

	apiPort := config.Cfg.APIPort
	// apiPort := os.Getenv("APIPort")
	app.Logger.Fatal(app.Start(apiPort))
	// app.Start(apiPort)
}
