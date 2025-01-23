package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shyfeng955/go-clean/api/handlers"
)

type Router struct {
	userHandel handlers.UserHandel
}

func (r Router) Routers(engine *gin.Engine) {
	engine.POST("/getUserInfo", r.userHandel.GetUserInfo)
	engine.POST("/addUser", r.userHandel.AddUser)
}

func NewRouter(userHandel handlers.UserHandel) *Router {
	return &Router{
		userHandel: userHandel,
	}
}
