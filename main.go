package main

import (
	proto "github.com/abaeve/sde-service/proto"
	"github.com/abaeve/services-common/config"
	"github.com/abaeve/sde-service/handler"
)

var version string = "1.0.0"

func main() {
	service := config.NewService(version, "sde-service", initialize)

	proto.RegisterTypeQueryHandler(service.Server(), handler.TypeQueryHandler)

	service.Run()
}

func initialize(config *config.Configuration) error {
	return nil
}