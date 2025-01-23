package main

import (
	"github.com/shyfeng955/go-clean/api"
	"github.com/shyfeng955/go-clean/api/handlers"
	"github.com/shyfeng955/go-clean/app"
	"github.com/shyfeng955/go-clean/repo"
	"github.com/shyfeng955/go-clean/service"
)

func main() {
	db := app.InitDB()
	// 初始化repo
	userRepo := repo.NewUserRepo(db)
	// 初始化service
	userService := service.NewUserService(userRepo)
	// 初始化api
	handler := handlers.NewUserHandel(userService)
	// 初始化router
	router := api.NewRouter(handler)
	engine := app.NewGinEngine()

	server := app.NewServer(engine, router)
	server.Start()
}
