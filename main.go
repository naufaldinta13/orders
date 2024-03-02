package main

import (
	"os"

	"github.com/naufaldinta13/orders/src"

	"github.com/env-io/factory"
	"github.com/env-io/factory/grpc"
	"github.com/env-io/factory/rest"
	"github.com/env-io/factory/sql"
	"github.com/joho/godotenv"
)

func init() {
	// load env variable file if exists
	godotenv.Load()

	// application config
	factory.AppConfig = &factory.Config{
		AppName:    "testing.order",
		IsDev:      os.Getenv("DEBUG_MODE") == "true",
		AppVersion: os.Getenv("APP_VERSION"),
		AppService: os.Getenv("APP_SERVICE"),
	}

	// new logger instances
	factory.NewLogger(factory.AppConfig.AppName)

	// new mysql instances
	initMysqlConnection()

	// initial rest server
	initRestServer()

	// initial grpc server
	initGrpcServer()
}

func main() {
	factory.Routine.Add(rest.Start, rest.Shutdown)
	factory.Routine.Add(grpc.Start, grpc.Shutdown)

	factory.Logger.Sugar().Error(factory.Routine.Run())
}

func initMysqlConnection() {
	c := &sql.Config{
		Server:     os.Getenv("MYSQL_SERVER"),
		Username:   os.Getenv("MYSQL_USERNAME"),
		Password:   os.Getenv("MSQYL_PASSWORD"),
		Database:   os.Getenv("MYSQL_DATABASE"),
		Datasource: os.Getenv("MYSQL_DATASOURCE"),
	}

	if e := sql.NewConnection(c, nil); e != nil {
		factory.Logger.Error(e.Error())
	}
}

func initRestServer() {
	c := &rest.Config{
		Server:    os.Getenv("REST_SERVER"),
		IsDev:     factory.AppConfig.IsDev,
		JwtSecret: os.Getenv("REST_JWT"),
	}

	rest.NewServer(c, src.RegisterRestHandler)
}

func initGrpcServer() {
	c := &grpc.Config{
		Name:           factory.AppConfig.AppName,
		RegistryServer: os.Getenv("SERVICE_REGISTRY"),
		Server:         os.Getenv("SERVICE_SERVER"),
	}

	if e := grpc.NewService(c, nil); e != nil {
		factory.Logger.Panic(e.Error())
	}
}
