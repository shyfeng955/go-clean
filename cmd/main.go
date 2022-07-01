package main

/**
 *	start here!!
 */

func main() {
	//// 初始化db
	//db := app.InitDB()
	//// 初始化repo
	//repository := repo.NewArticleRepository(db)
	//// 初始化service
	//articleService := service.NewArticleService(repository)
	//// 初始化api
	//handler := handlers.NewArticleHandel(articleService)
	//// 初始化router
	//router := api.NewRouter(handler)
	//
	//engine := app.NewGinEngine()
	//
	//server := app.NewServer(engine, router)
	//
	//server.Start()
	server := InitServer()

	server.Start()
}
