//go:build wireinject
// +build wireinject

package main

import (
	"github.com/cr1ng3T34m1337/arch-lab3-api/api/dataProviders"
	"github.com/cr1ng3T34m1337/arch-lab3-api/api/httpHandlers"
	"github.com/cr1ng3T34m1337/arch-lab3-api/db"
	"github.com/google/wire"
)

func ComposeApiServer(connectionConfig *db.Connection, port HttpPortNumber) (*ApiServer, error) {
	wire.Build(
		db.DbProvider,
		dataProviders.StoreProvider,
		httpHandlers.HandlersProvider,
		wire.Struct(new(ApiServer), "Port", "HttpHandlers"),
	)
	return nil, nil
}
