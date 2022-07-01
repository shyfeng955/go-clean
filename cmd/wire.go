//+build wireinject

package main

import (
	"fhz/api"
	"fhz/api/handlers"
	"fhz/app"
	"fhz/repo"
	"fhz/service"
	"github.com/google/wire"
)

func InitServer() *app.Server {
	wire.Build(
		app.InitDB,
		repo.NewArticleRepository,
		service.NewArticleService,
		handlers.NewArticleHandel,
		api.NewRouter,
		app.NewGinEngine,
		app.NewServer,
	)
	return &app.Server{}
}
