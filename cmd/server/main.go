package main

import (
	"fmt"

	"github.com/cr1ng3T34m1337/arch-lab3-api/db"
	"github.com/cr1ng3T34m1337/arch-lab3-api/tools"
)

func main() {
	config, err := tools.ParseConfig("config.json")
	if err != nil {
		panic(err)
	}

	conn := &db.Connection{
		DbName:   config.Db.DbName,
		User:     config.Db.User,
		Password: config.Db.Password,
		Host:     config.Db.Host,
	}
	apiServer, err := ComposeApiServer(conn, config.Api.Addr, HttpPortNumber(config.Api.Port))
	if err != nil {
		panic(err)
	}
	fmt.Println("starting on port", apiServer.Port)
	if err := apiServer.Start(); err != nil {
		panic(err)
	}
}
