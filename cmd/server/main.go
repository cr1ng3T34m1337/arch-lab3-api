package main

import (
	"flag"
	"fmt"

	"github.com/cr1ng3T34m1337/arch-lab3-api/db"
	"github.com/cr1ng3T34m1337/arch-lab3-api/tools"
)

var (
	portNum = flag.Int("p", 8080, "http port to be run on")
)

func main() {
	flag.Parse()
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
	apiServer, err := ComposeApiServer(conn, HttpPortNumber(*portNum))
	if err != nil {
		panic(err)
	}
	fmt.Println("starting on port", apiServer.Port)
	if err := apiServer.Start(); err != nil {
		panic(err)
	}
}
