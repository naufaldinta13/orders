package main

import (
	"log"
	"os"

	"github.com/naufaldinta13/orders/config"
	"github.com/naufaldinta13/orders/routes"

	"github.com/joho/godotenv"
	"github.com/oklog/run"
)

var Routine run.Group

func initMySQLConnection() {
	c := &config.DBCOnfig{
		Server:   os.Getenv("MYSQL_SERVER"),
		Username: os.Getenv("MYSQL_USERNAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}

	if e := config.NewDBConnection(c); e != nil {
		log.Fatal(e)
	}
}

func initGrpcConnection() {
	c := &config.GrpcConfig{
		Name:           "rent.orders",
		RegistryServer: os.Getenv("SERVICE_REGISTRY"),
		Server:         os.Getenv("SERVICE_SERVER"),
	}

	if e := config.NewGrpcConnection(c, nil); e != nil {
		log.Fatal(e)
	}
}

func init() {
	godotenv.Load()

	initMySQLConnection()
	initGrpcConnection()
}

func main() {
	route := routes.SetupRoutes()

	go func() {
		route.Run(os.Getenv("REST_SERVER"))
	}()

	Routine.Add(config.Start, config.Shutdown)
	Routine.Run()
}
