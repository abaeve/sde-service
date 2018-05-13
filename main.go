package main

import (
	"fmt"
	"github.com/abaeve/sde-service/handler"
	proto "github.com/abaeve/sde-service/proto"
	"github.com/chremoas/services-common/config"
	"github.com/micro/go-micro"
	"gopkg.in/mgo.v2"
)

var version string = "1.0.0"

var service micro.Service

func main() {
	service = config.NewService(version, "srv", "sde", initialize)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func initialize(config *config.Configuration) error {
	mongoConfig := config.Extensions["mongodb"].(map[interface{}]interface{})["dial"].(string)

	session, err := mgo.Dial(mongoConfig)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	proto.RegisterTypeQueryHandler(service.Server(), handler.NewTypeQueryHandler(session))

	c := session.DB("sde").C("types")
	index := mgo.Index{
		Key:        []string{"name", "typeID"},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     false,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	return nil
}
