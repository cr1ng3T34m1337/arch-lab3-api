// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/cr1ng3T34m1337/arch-lab3-api/api/dataProviders"
	"github.com/cr1ng3T34m1337/arch-lab3-api/api/httpHandlers"
	"github.com/cr1ng3T34m1337/arch-lab3-api/db"
)

// Injectors from wire.go:

func ComposeApiServer(connectionConfig *db.Connection, addr string, port HttpPortNumber) (*ApiServer, error) {
	sqlDB, err := db.DbProvider(connectionConfig)
	if err != nil {
		return nil, err
	}
	store := dataProviders.StoreProvider(sqlDB)
	v := httpHandlers.HandlersProvider(store)
	apiServer := &ApiServer{
		Addr:         addr,
		Port:         port,
		HttpHandlers: v,
	}
	return apiServer, nil
}
