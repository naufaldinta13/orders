package config

import (
	"time"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
)

type GrpcConfig struct {
	Name           string
	Server         string
	RegistryServer string
}

var Service micro.Service

func NewGrpcConnection(c *GrpcConfig, cb func(micro.Service)) (e error) {
	Service = micro.NewService(
		micro.Name(c.Name),
		micro.Registry(registry.NewRegistry(registry.Addrs(c.RegistryServer))),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Address(c.Server),
	)

	Service.Init()
	Service.Server().Init(server.Wait(nil))

	if cb != nil {
		cb(Service)
	}

	return
}

func Start() error {
	return Service.Run()
}

func Shutdown(e error) {
	Service.Server().Stop()
}
